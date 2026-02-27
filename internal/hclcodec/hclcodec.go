// Package hclcodec provides reflection-based generation of HCL resource blocks
// from Terraform provider model structs.
//
// It reads `tfsdk` struct tags for attribute names and `json` struct tags for
// the required/optional/computed classification. Only attributes tagged as
// required, optional (when non-null), or computed_optional (when non-null) are
// emitted. Computed-only attributes are excluded from HCL since they exist
// only in state.
package hclcodec

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/zclconf/go-cty/cty"
)

// NestedMode specifies how a nested field should be rendered in HCL.
type NestedMode int

const (
	// NestedModeBlock renders as: name { ... } (default, backward compatible)
	NestedModeBlock NestedMode = iota
	// NestedModeAttr renders as: name = { ... }
	NestedModeAttr
)

// SchemaInfo provides schema-level hints for HCL serialization of nested fields.
// Keys are tfsdk attribute names. Only nested/object fields need entries;
// scalar fields are serialized identically regardless of schema info.
// A nil SchemaInfo is safe and causes all nested fields to use block syntax.
type SchemaInfo map[string]FieldSchema

// FieldSchema describes HCL serialization hints for a single field.
type FieldSchema struct {
	NestedMode NestedMode
	Children   SchemaInfo

	// AllowedValues lists the canonical enum values for a string field, extracted
	// from schema validators (e.g., OneOfCaseInsensitive). When non-nil, the codec
	// maps the model's value to the matching canonical value before emitting HCL.
	AllowedValues []string

	// ComputedOnly is true for attributes that are Computed but not Optional.
	// These fields should appear in state but NOT in HCL configuration.
	ComputedOnly bool
}

// lookup returns the FieldSchema for a given tfsdk attribute name.
// If the key is absent or the receiver is nil it returns a zero FieldSchema
// (NestedModeBlock, nil Children), which preserves backward-compatible behavior.
func (s SchemaInfo) lookup(name string) FieldSchema {
	if s == nil {
		return FieldSchema{}
	}
	return s[name]
}

// GenerateResourceBlock converts a provider model struct to a complete HCL resource block.
// Only required, optional (non-null), and computed_optional (non-null) fields are included.
// The optional schema parameter provides hints for nested field serialization
// (attribute syntax vs block syntax). Pass nil for backward-compatible block syntax.
func GenerateResourceBlock(resourceType, resourceName string, model any, schema SchemaInfo) (string, error) {
	body, err := generateBody(model, 2, schema)
	if err != nil {
		return "", err
	}

	var sb strings.Builder
	fmt.Fprintf(&sb, "resource %q %q {\n", resourceType, resourceName)
	sb.WriteString(body)
	sb.WriteString("}\n")

	return sb.String(), nil
}

// GenerateAttributes returns only the inner attributes of a resource block (no resource wrapper).
func GenerateAttributes(model any, schema SchemaInfo) (string, error) {
	return generateBody(model, 2, schema)
}

// fieldInfo holds parsed tag information for a struct field.
type fieldInfo struct {
	tfsdkName   string
	jsonOption  string // "required", "optional", "computed", "computed_optional"
	isPathField bool   // has path tag instead of/in addition to json
	fieldValue  reflect.Value
	fieldType   reflect.StructField
}

// generateBody produces the indented HCL body for a struct.
func generateBody(model any, indent int, schema SchemaInfo) (string, error) {
	v := reflect.ValueOf(model)
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		if v.IsNil() {
			return "", nil
		}
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return "", fmt.Errorf("hclcodec: expected struct, got %s", v.Kind())
	}

	fields := collectFields(v)

	var sb strings.Builder
	for _, fi := range fields {
		if !shouldIncludeInHCL(fi) {
			continue
		}

		hcl, err := serializeFieldToHCL(fi, indent, schema)
		if err != nil {
			return "", fmt.Errorf("hclcodec: field %s: %w", fi.fieldType.Name, err)
		}
		if hcl != "" {
			sb.WriteString(hcl)
		}
	}

	return sb.String(), nil
}

// collectFields extracts field information from a struct.
func collectFields(v reflect.Value) []fieldInfo {
	t := v.Type()
	var fields []fieldInfo

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if !field.IsExported() {
			continue
		}

		tfsdkTag := field.Tag.Get("tfsdk")
		if tfsdkTag == "" || tfsdkTag == "-" {
			continue
		}

		jsonTag := field.Tag.Get("json")
		pathTag := field.Tag.Get("path")
		option := parseTagOption(jsonTag)
		isPath := false

		// Some fields use path tag for the option (e.g., access_policy_id)
		if option == "" && pathTag != "" {
			option = parseTagOption(pathTag)
			isPath = true
		}

		fields = append(fields, fieldInfo{
			tfsdkName:   tfsdkTag,
			jsonOption:  option,
			isPathField: isPath,
			fieldValue:  v.Field(i),
			fieldType:   field,
		})
	}

	return fields
}

// parseTagOption extracts the option (required/optional/computed/computed_optional)
// from a struct tag value like "fieldName,required" or "fieldName,computed_optional".
func parseTagOption(tag string) string {
	parts := strings.Split(tag, ",")
	for _, part := range parts[1:] {
		switch part {
		case "required", "optional", "computed", "computed_optional":
			return part
		}
	}
	// For path tags, check if "required" appears
	if len(parts) > 1 {
		for _, part := range parts[1:] {
			if part == "required" {
				return "required"
			}
		}
	}
	return ""
}

// shouldIncludeInHCL determines if a field should be included in HCL config.
func shouldIncludeInHCL(fi fieldInfo) bool {
	switch fi.jsonOption {
	case "required":
		return true
	case "optional", "computed_optional":
		return !isNullOrUnknown(fi.fieldValue)
	case "computed":
		return false
	default:
		// Unknown option or no option: skip
		return false
	}
}

// isNullOrUnknown checks if a value is null, unknown, or nil.
func isNullOrUnknown(v reflect.Value) bool {
	if v.Kind() == reflect.Ptr {
		return v.IsNil()
	}

	if v.Kind() == reflect.Interface {
		if v.IsNil() {
			return true
		}
		v = v.Elem()
	}

	iface := v.Interface()

	if attrVal, ok := iface.(attr.Value); ok {
		return attrVal.IsNull() || attrVal.IsUnknown()
	}

	// For custom list types, check embedded ListValue
	if lv, ok := extractListValue(v); ok {
		return lv.IsNull() || lv.IsUnknown()
	}

	return false
}

// serializeFieldToHCL serializes a single field to HCL text.
func serializeFieldToHCL(fi fieldInfo, indent int, schema SchemaInfo) (string, error) {
	v := fi.fieldValue

	// Handle pointer types
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return "", nil
		}
		v = v.Elem()
	}

	iface := v.Interface()
	prefix := strings.Repeat(" ", indent)

	// Look up schema info for this field.
	fs := schema.lookup(fi.tfsdkName)

	// Handle attr.Value types
	if attrVal, ok := iface.(attr.Value); ok {
		if attrVal.IsNull() || attrVal.IsUnknown() {
			return "", nil
		}
		return serializeAttrValueToHCL(fi.tfsdkName, attrVal, v, prefix, indent, fs)
	}

	// Handle pointer-to-slice types like *[]types.String
	if v.Kind() == reflect.Slice {
		return serializeSliceToHCL(fi.tfsdkName, v, prefix, indent)
	}

	// Handle nested struct (without attr.Value)
	if v.Kind() == reflect.Struct {
		return serializeNestedStructToHCL(fi.tfsdkName, v, indent, fs.NestedMode == NestedModeAttr, fs.Children)
	}

	return "", nil
}

// mapEnumValue maps an API value to a canonical schema value using the allowed
// values list extracted from schema validators. It tries, in order:
//  1. Case-insensitive exact match
//  2. Case-insensitive suffix match after "_" (e.g., "admin" matches "USER_ROLE_ORG_ADMIN")
//
// If no match is found, the original value is returned unchanged.
func mapEnumValue(apiValue string, allowed []string) string {
	if len(allowed) == 0 {
		return apiValue
	}

	upper := strings.ToUpper(apiValue)

	// 1. Case-insensitive exact match.
	for _, v := range allowed {
		if strings.EqualFold(v, apiValue) {
			return v
		}
	}

	// 2. Suffix match: "admin" → "USER_ROLE_ORG_ADMIN".
	suffix := "_" + upper
	var match string
	for _, v := range allowed {
		if strings.HasSuffix(strings.ToUpper(v), suffix) {
			if match != "" {
				// Ambiguous — multiple candidates. Return the original.
				return apiValue
			}
			match = v
		}
	}
	if match != "" {
		return match
	}

	return apiValue
}

// hclQuote produces an HCL-safe double-quoted string literal by delegating to
// hclwrite.TokensForValue, which uses the HCL library's own escapeQuotedStringLit.
// This correctly handles all HCL special sequences (${, %{), control characters,
// non-printable Unicode, and any future syntax additions.
func hclQuote(s string) string {
	return string(hclwrite.TokensForValue(cty.StringVal(s)).Bytes())
}

// serializeAttrValueToHCL converts an attr.Value to HCL syntax.
func serializeAttrValueToHCL(name string, attrVal attr.Value, v reflect.Value, prefix string, indent int, fs FieldSchema) (string, error) {
	switch val := attrVal.(type) {
	case basetypes.StringValue:
		s := mapEnumValue(val.ValueString(), fs.AllowedValues)
		return fmt.Sprintf("%s%s = %s\n", prefix, name, hclQuote(s)), nil

	case types.Int64:
		return fmt.Sprintf("%s%s = %d\n", prefix, name, val.ValueInt64()), nil

	case basetypes.BoolValue:
		return fmt.Sprintf("%s%s = %v\n", prefix, name, val.ValueBool()), nil

	case basetypes.Float64Value:
		return fmt.Sprintf("%s%s = %g\n", prefix, name, val.ValueFloat64()), nil

	case basetypes.ListValue:
		return serializeListValueToHCL(name, val, prefix, indent, fs.NestedMode == NestedModeAttr, fs.Children)

	case basetypes.ObjectValue:
		return serializeObjectValueToHCL(name, val, prefix, indent, fs.NestedMode == NestedModeAttr, fs.Children)
	}

	// Custom types wrapping string (timetypes.RFC3339, jsontypes.Normalized)
	if sv, ok := attrVal.(stringValuer); ok {
		return fmt.Sprintf("%s%s = %s\n", prefix, name, hclQuote(sv.ValueString())), nil
	}

	// Custom list types (customfield.List, customfield.NestedObjectList)
	if lv, ok := extractListValue(v); ok {
		if lv.IsNull() || lv.IsUnknown() {
			return "", nil
		}
		return serializeListValueToHCL(name, lv, prefix, indent, fs.NestedMode == NestedModeAttr, fs.Children)
	}

	return "", fmt.Errorf("unsupported attr.Value type %T for HCL", attrVal)
}

// stringValuer is an interface for types with ValueString().
type stringValuer interface {
	ValueString() string
}

// serializeListValueToHCL serializes a list to HCL.
// When asAttr is true and elements are objects, it uses attribute syntax: name = [{ ... }]
// When asAttr is false and elements are objects, it uses repeated block syntax: name { ... }
func serializeListValueToHCL(name string, lv basetypes.ListValue, prefix string, indent int, asAttr bool, childSchema SchemaInfo) (string, error) {
	elements := lv.Elements()

	// Check if elements are objects (nested blocks or attribute list-of-objects)
	if len(elements) > 0 {
		if _, isObj := elements[0].(basetypes.ObjectValue); isObj {
			if asAttr {
				return serializeListOfObjectsAsAttrToHCL(name, elements, prefix, indent, childSchema)
			}
			return serializeListOfObjectsToHCL(name, elements, indent, childSchema)
		}
	}

	// Simple list (strings, numbers, etc.)
	var sb strings.Builder
	fmt.Fprintf(&sb, "%s%s = [", prefix, name)

	for i, elem := range elements {
		if i > 0 {
			sb.WriteString(", ")
		}
		hcl, err := elementToHCL(elem)
		if err != nil {
			return "", err
		}
		sb.WriteString(hcl)
	}

	sb.WriteString("]\n")
	return sb.String(), nil
}

// serializeListOfObjectsAsAttrToHCL serializes a list of objects using attribute syntax:
//
//	name = [{
//	  key = value
//	}]
//
// This is required for ListNestedAttribute fields in the Terraform Plugin Framework.
func serializeListOfObjectsAsAttrToHCL(name string, elements []attr.Value, prefix string, indent int, childSchema SchemaInfo) (string, error) {
	var sb strings.Builder
	innerPrefix := strings.Repeat(" ", indent+2)

	fmt.Fprintf(&sb, "%s%s = [", prefix, name)

	written := 0
	for _, elem := range elements {
		ov, ok := elem.(basetypes.ObjectValue)
		if !ok {
			continue
		}
		if ov.IsNull() || ov.IsUnknown() {
			continue
		}

		if written > 0 {
			sb.WriteString(", ")
		}
		written++
		sb.WriteString("{\n")
		attrs := ov.Attributes()

		keys := make([]string, 0, len(attrs))
		for k := range attrs {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, key := range keys {
			val := attrs[key]
			if val.IsNull() || val.IsUnknown() {
				continue
			}

			cfs := childSchema.lookup(key)
			hcl, err := serializeObjectAttrToHCL(key, val, innerPrefix, indent+2, cfs)
			if err != nil {
				return "", err
			}
			sb.WriteString(hcl)
		}
		fmt.Fprintf(&sb, "%s}", prefix)
	}

	sb.WriteString("]\n")
	return sb.String(), nil
}

// serializeListOfObjectsToHCL serializes a list of objects as repeated HCL blocks.
func serializeListOfObjectsToHCL(name string, elements []attr.Value, indent int, childSchema SchemaInfo) (string, error) {
	var sb strings.Builder
	prefix := strings.Repeat(" ", indent)
	innerPrefix := strings.Repeat(" ", indent+2)

	for _, elem := range elements {
		ov, ok := elem.(basetypes.ObjectValue)
		if !ok {
			continue
		}
		if ov.IsNull() || ov.IsUnknown() {
			continue
		}

		fmt.Fprintf(&sb, "%s%s {\n", prefix, name)
		attrs := ov.Attributes()

		// Sort keys for deterministic output
		keys := make([]string, 0, len(attrs))
		for k := range attrs {
			keys = append(keys, k)
		}
		sort.Strings(keys)

	for _, key := range keys {
		val := attrs[key]
		if val.IsNull() || val.IsUnknown() {
			continue
		}

		// Check if this child field has nested schema info.
		cfs := childSchema.lookup(key)

		hcl, err := serializeObjectAttrToHCL(key, val, innerPrefix, indent+2, cfs)
		if err != nil {
			return "", err
		}
		sb.WriteString(hcl)
	}

	fmt.Fprintf(&sb, "%s}\n", prefix)
}

	return sb.String(), nil
}

// serializeObjectValueToHCL serializes an object value to HCL.
// When asAttr is true it uses attribute syntax: name = { ... }
// When asAttr is false it uses block syntax: name { ... }
func serializeObjectValueToHCL(name string, ov basetypes.ObjectValue, prefix string, indent int, asAttr bool, childSchema SchemaInfo) (string, error) {
	var sb strings.Builder
	innerPrefix := strings.Repeat(" ", indent+2)

	if asAttr {
		fmt.Fprintf(&sb, "%s%s = {\n", prefix, name)
	} else {
		fmt.Fprintf(&sb, "%s%s {\n", prefix, name)
	}
	attrs := ov.Attributes()

	keys := make([]string, 0, len(attrs))
	for k := range attrs {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		val := attrs[key]
		if val.IsNull() || val.IsUnknown() {
			continue
		}

		// Check if this child field has nested schema info.
		cfs := childSchema.lookup(key)

		hcl, err := serializeObjectAttrToHCL(key, val, innerPrefix, indent+2, cfs)
		if err != nil {
			return "", err
		}
		sb.WriteString(hcl)
	}

	fmt.Fprintf(&sb, "%s}\n", prefix)
	return sb.String(), nil
}

// serializeSliceToHCL serializes a Go slice (like []types.String) to HCL.
func serializeSliceToHCL(name string, v reflect.Value, prefix string, indent int) (string, error) {
	if v.IsNil() {
		return "", nil
	}

	var sb strings.Builder
	fmt.Fprintf(&sb, "%s%s = [", prefix, name)

	for i := 0; i < v.Len(); i++ {
		if i > 0 {
			sb.WriteString(", ")
		}
		elem := v.Index(i)
		iface := elem.Interface()

		if attrVal, ok := iface.(attr.Value); ok {
			hcl, err := elementToHCL(attrVal)
			if err != nil {
				return "", err
			}
			sb.WriteString(hcl)
		} else if elem.Kind() == reflect.Ptr && !elem.IsNil() {
			// Handle *NestedStruct elements in arrays - these need block syntax
			hcl, err := serializeNestedStructToHCL(name, elem.Elem(), indent, false, nil)
			if err != nil {
				return "", err
			}
			sb.WriteString(hcl)
		}
	}

	sb.WriteString("]\n")
	return sb.String(), nil
}

// serializeNestedStructToHCL serializes a nested struct to HCL.
// When asAttr is true it uses attribute syntax: name = { ... }
// When asAttr is false it uses block syntax: name { ... }
func serializeNestedStructToHCL(name string, v reflect.Value, indent int, asAttr bool, childSchema SchemaInfo) (string, error) {
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		if v.IsNil() {
			return "", nil
		}
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return "", nil
	}

	body, err := generateBody(v.Interface(), indent+2, childSchema)
	if err != nil {
		return "", err
	}

	if body == "" {
		return "", nil
	}

	prefix := strings.Repeat(" ", indent)
	var sb strings.Builder
	if asAttr {
		fmt.Fprintf(&sb, "%s%s = {\n", prefix, name)
	} else {
		fmt.Fprintf(&sb, "%s%s {\n", prefix, name)
	}
	sb.WriteString(body)
	fmt.Fprintf(&sb, "%s}\n", prefix)
	return sb.String(), nil
}

// serializeObjectAttrToHCL serializes a single attribute within an object to HCL.
// Unlike elementToHCL, this handles complex types (lists, nested objects) as well as scalars.
// Fields marked ComputedOnly in the schema are skipped (they belong in state, not HCL config).
func serializeObjectAttrToHCL(key string, val attr.Value, prefix string, indent int, cfs FieldSchema) (string, error) {
	if cfs.ComputedOnly {
		return "", nil
	}
	switch v := val.(type) {
	case basetypes.ObjectValue:
		return serializeObjectValueToHCL(key, v, prefix, indent, cfs.NestedMode == NestedModeAttr, cfs.Children)
	case basetypes.ListValue:
		return serializeListValueToHCL(key, v, prefix, indent, cfs.NestedMode == NestedModeAttr, cfs.Children)
	default:
		hcl, err := elementToHCL(val)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s%s = %s\n", prefix, key, hcl), nil
	}
}

// elementToHCL converts a single attr.Value to its HCL representation (without key).
func elementToHCL(val attr.Value) (string, error) {
	if val.IsNull() || val.IsUnknown() {
		return "null", nil
	}

	switch v := val.(type) {
	case basetypes.StringValue:
		return hclQuote(v.ValueString()), nil
	case basetypes.Int64Value:
		return fmt.Sprintf("%d", v.ValueInt64()), nil
	case basetypes.BoolValue:
		return fmt.Sprintf("%v", v.ValueBool()), nil
	case basetypes.Float64Value:
		return fmt.Sprintf("%g", v.ValueFloat64()), nil
	}

	if sv, ok := val.(stringValuer); ok {
		return hclQuote(sv.ValueString()), nil
	}

	return "", fmt.Errorf("unsupported element type %T for HCL", val)
}

// extractListValue tries to find a basetypes.ListValue embedded in a struct.
func extractListValue(v reflect.Value) (basetypes.ListValue, bool) {
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		if v.IsNil() {
			return basetypes.ListValue{}, false
		}
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return basetypes.ListValue{}, false
	}

	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if !field.IsExported() {
			continue
		}
		fieldVal := v.Field(i)

		if lv, ok := fieldVal.Interface().(basetypes.ListValue); ok {
			return lv, true
		}

		if field.Anonymous && fieldVal.Kind() == reflect.Struct {
			if lv, found := extractListValue(fieldVal); found {
				return lv, true
			}
		}
	}

	return basetypes.ListValue{}, false
}

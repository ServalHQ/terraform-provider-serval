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

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// GenerateResourceBlock converts a provider model struct to a complete HCL resource block.
// Only required, optional (non-null), and computed_optional (non-null) fields are included.
func GenerateResourceBlock(resourceType, resourceName string, model any) (string, error) {
	body, err := generateBody(model, 2)
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
func GenerateAttributes(model any) (string, error) {
	return generateBody(model, 2)
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
func generateBody(model any, indent int) (string, error) {
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

		hcl, err := serializeFieldToHCL(fi, indent)
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
func serializeFieldToHCL(fi fieldInfo, indent int) (string, error) {
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

	// Handle attr.Value types
	if attrVal, ok := iface.(attr.Value); ok {
		if attrVal.IsNull() || attrVal.IsUnknown() {
			return "", nil
		}
		return serializeAttrValueToHCL(fi.tfsdkName, attrVal, v, prefix, indent)
	}

	// Handle pointer-to-slice types like *[]types.String
	if v.Kind() == reflect.Slice {
		return serializeSliceToHCL(fi.tfsdkName, v, prefix, indent)
	}

	// Handle nested struct (without attr.Value)
	if v.Kind() == reflect.Struct {
		return serializeNestedStructToHCL(fi.tfsdkName, v, indent)
	}

	return "", nil
}

// serializeAttrValueToHCL converts an attr.Value to HCL syntax.
func serializeAttrValueToHCL(name string, attrVal attr.Value, v reflect.Value, prefix string, indent int) (string, error) {
	switch val := attrVal.(type) {
	case basetypes.StringValue:
		return fmt.Sprintf("%s%s = %q\n", prefix, name, val.ValueString()), nil

	case types.Int64:
		return fmt.Sprintf("%s%s = %d\n", prefix, name, val.ValueInt64()), nil

	case basetypes.BoolValue:
		return fmt.Sprintf("%s%s = %v\n", prefix, name, val.ValueBool()), nil

	case basetypes.Float64Value:
		return fmt.Sprintf("%s%s = %g\n", prefix, name, val.ValueFloat64()), nil

	case basetypes.ListValue:
		return serializeListValueToHCL(name, val, prefix, indent)

	case basetypes.ObjectValue:
		return serializeObjectValueToHCL(name, val, prefix, indent)
	}

	// Custom types wrapping string (timetypes.RFC3339, jsontypes.Normalized)
	if sv, ok := attrVal.(stringValuer); ok {
		return fmt.Sprintf("%s%s = %q\n", prefix, name, sv.ValueString()), nil
	}

	// Custom list types (customfield.List, customfield.NestedObjectList)
	if lv, ok := extractListValue(v); ok {
		if lv.IsNull() || lv.IsUnknown() {
			return "", nil
		}
		return serializeListValueToHCL(name, lv, prefix, indent)
	}

	return "", fmt.Errorf("unsupported attr.Value type %T for HCL", attrVal)
}

// stringValuer is an interface for types with ValueString().
type stringValuer interface {
	ValueString() string
}

// serializeListValueToHCL serializes a list to HCL.
func serializeListValueToHCL(name string, lv basetypes.ListValue, prefix string, indent int) (string, error) {
	elements := lv.Elements()

	// Check if elements are objects (nested blocks)
	if len(elements) > 0 {
		if _, isObj := elements[0].(basetypes.ObjectValue); isObj {
			return serializeListOfObjectsToHCL(name, elements, indent)
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

// serializeListOfObjectsToHCL serializes a list of objects as repeated HCL blocks.
func serializeListOfObjectsToHCL(name string, elements []attr.Value, indent int) (string, error) {
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
			hcl, err := elementToHCL(val)
			if err != nil {
				return "", err
			}
			fmt.Fprintf(&sb, "%s%s = %s\n", innerPrefix, key, hcl)
		}

		fmt.Fprintf(&sb, "%s}\n", prefix)
	}

	return sb.String(), nil
}

// serializeObjectValueToHCL serializes an object as an HCL block.
func serializeObjectValueToHCL(name string, ov basetypes.ObjectValue, prefix string, indent int) (string, error) {
	var sb strings.Builder
	innerPrefix := strings.Repeat(" ", indent+2)

	fmt.Fprintf(&sb, "%s%s {\n", prefix, name)
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
		hcl, err := elementToHCL(val)
		if err != nil {
			return "", err
		}
		fmt.Fprintf(&sb, "%s%s = %s\n", innerPrefix, key, hcl)
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
			hcl, err := serializeNestedStructToHCL(name, elem.Elem(), indent)
			if err != nil {
				return "", err
			}
			sb.WriteString(hcl)
		}
	}

	sb.WriteString("]\n")
	return sb.String(), nil
}

// serializeNestedStructToHCL serializes a nested struct as an HCL block.
func serializeNestedStructToHCL(name string, v reflect.Value, indent int) (string, error) {
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		if v.IsNil() {
			return "", nil
		}
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return "", nil
	}

	body, err := generateBody(v.Interface(), indent+2)
	if err != nil {
		return "", err
	}

	if body == "" {
		return "", nil
	}

	prefix := strings.Repeat(" ", indent)
	var sb strings.Builder
	fmt.Fprintf(&sb, "%s%s {\n", prefix, name)
	sb.WriteString(body)
	fmt.Fprintf(&sb, "%s}\n", prefix)
	return sb.String(), nil
}

// elementToHCL converts a single attr.Value to its HCL representation (without key).
func elementToHCL(val attr.Value) (string, error) {
	if val.IsNull() || val.IsUnknown() {
		return "null", nil
	}

	switch v := val.(type) {
	case basetypes.StringValue:
		return fmt.Sprintf("%q", v.ValueString()), nil
	case basetypes.Int64Value:
		return fmt.Sprintf("%d", v.ValueInt64()), nil
	case basetypes.BoolValue:
		return fmt.Sprintf("%v", v.ValueBool()), nil
	case basetypes.Float64Value:
		return fmt.Sprintf("%g", v.ValueFloat64()), nil
	}

	if sv, ok := val.(stringValuer); ok {
		return fmt.Sprintf("%q", sv.ValueString()), nil
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

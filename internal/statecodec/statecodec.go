// Package statecodec provides reflection-based serialization of Terraform provider
// model structs to the JSON attributes format used in terraform.tfstate files.
//
// It reads `tfsdk` struct tags for attribute names and handles all standard
// Terraform plugin framework value types (types.String, types.Int64, types.Bool,
// timetypes.RFC3339, jsontypes.Normalized, customfield.List, customfield.NestedObjectList,
// pointer-to-slice types, and nested structs).
package statecodec

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// SerializeToStateAttributes converts a provider model struct to the
// JSON attributes map used in terraform.tfstate, using tfsdk tags as keys.
// All fields with tfsdk tags are included (required, optional, computed, computed_optional).
func SerializeToStateAttributes(model any) (json.RawMessage, error) {
	v := reflect.ValueOf(model)
	attrs, err := serializeStruct(v)
	if err != nil {
		return nil, err
	}
	return json.Marshal(attrs)
}

// serializeStruct converts a struct with tfsdk tags to a map.
func serializeStruct(v reflect.Value) (map[string]interface{}, error) {
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		if v.IsNil() {
			return nil, nil
		}
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("statecodec: expected struct, got %s", v.Kind())
	}

	attrs := make(map[string]interface{})
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if !field.IsExported() {
			continue
		}

		tfsdkTag := field.Tag.Get("tfsdk")
		if tfsdkTag == "" || tfsdkTag == "-" {
			// Check for embedded struct (anonymous field) without tfsdk tag
			if field.Anonymous {
				// Skip embedded types like basetypes.ListValue that don't have tfsdk tags
				continue
			}
			continue
		}

		fieldVal := v.Field(i)
		jsonVal, err := serializeValue(fieldVal)
		if err != nil {
			return nil, fmt.Errorf("statecodec: field %s (tfsdk:%q): %w", field.Name, tfsdkTag, err)
		}
		attrs[tfsdkTag] = jsonVal
	}

	return attrs, nil
}

// serializeValue converts a single field value to its JSON representation for state.
func serializeValue(v reflect.Value) (interface{}, error) {
	// Handle interface values
	if v.Kind() == reflect.Interface {
		if v.IsNil() {
			return nil, nil
		}
		v = v.Elem()
	}

	// Handle nil pointers
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil, nil
		}
		return serializeValue(v.Elem())
	}

	// Try to get the underlying value as an interface
	iface := v.Interface()

	// Check for attr.Value interface (all terraform types implement this)
	if attrVal, ok := iface.(attr.Value); ok {
		return serializeAttrValue(attrVal, v)
	}

	// Handle pointer-to-slice types like *[]types.String
	if v.Kind() == reflect.Slice {
		return serializeSlice(v)
	}

	// Handle nested struct (without attr.Value interface)
	if v.Kind() == reflect.Struct {
		return serializeStruct(v)
	}

	return nil, fmt.Errorf("statecodec: unsupported type %s", v.Type())
}

// serializeAttrValue converts an attr.Value to its JSON representation.
func serializeAttrValue(attrVal attr.Value, v reflect.Value) (interface{}, error) {
	if attrVal.IsNull() || attrVal.IsUnknown() {
		return nil, nil
	}

	switch val := attrVal.(type) {
	// types.String and any type that embeds/wraps basetypes.StringValue
	case basetypes.StringValue:
		return val.ValueString(), nil

	case types.Int64:
		return val.ValueInt64(), nil

	case basetypes.BoolValue:
		return val.ValueBool(), nil

	case basetypes.Float64Value:
		return val.ValueFloat64(), nil

	case basetypes.NumberValue:
		f, _ := val.ValueBigFloat().Float64()
		return f, nil

	case basetypes.ListValue:
		return serializeListValue(val)

	case basetypes.SetValue:
		return serializeSetValue(val)

	case basetypes.MapValue:
		return serializeMapValue(val)

	case basetypes.ObjectValue:
		return serializeObjectValue(val)
	}

	// For custom types that wrap standard types (e.g., timetypes.RFC3339 wraps StringValue,
	// jsontypes.Normalized wraps StringValue), try the stringValuer interface
	if sv, ok := attrVal.(stringValuer); ok {
		return sv.ValueString(), nil
	}

	// For custom list types (customfield.List, customfield.NestedObjectList),
	// check if the underlying struct contains a ListValue
	if lv, ok := extractListValue(v); ok {
		if lv.IsNull() || lv.IsUnknown() {
			return nil, nil
		}
		return serializeListValue(lv)
	}

	return nil, fmt.Errorf("statecodec: unsupported attr.Value type %T", attrVal)
}

// stringValuer is implemented by types that can return a string value.
type stringValuer interface {
	ValueString() string
}

// serializeListValue serializes a Terraform list value.
func serializeListValue(lv basetypes.ListValue) (interface{}, error) {
	elements := lv.Elements()
	result := make([]interface{}, 0, len(elements))
	for _, elem := range elements {
		val, err := serializeAttrValueSimple(elem)
		if err != nil {
			return nil, err
		}
		result = append(result, val)
	}
	return result, nil
}

// serializeSetValue serializes a Terraform set value.
func serializeSetValue(sv basetypes.SetValue) (interface{}, error) {
	elements := sv.Elements()
	result := make([]interface{}, 0, len(elements))
	for _, elem := range elements {
		val, err := serializeAttrValueSimple(elem)
		if err != nil {
			return nil, err
		}
		result = append(result, val)
	}
	return result, nil
}

// serializeMapValue serializes a Terraform map value.
func serializeMapValue(mv basetypes.MapValue) (interface{}, error) {
	elements := mv.Elements()
	result := make(map[string]interface{}, len(elements))
	for key, elem := range elements {
		val, err := serializeAttrValueSimple(elem)
		if err != nil {
			return nil, err
		}
		result[key] = val
	}
	return result, nil
}

// serializeObjectValue serializes a Terraform object value.
func serializeObjectValue(ov basetypes.ObjectValue) (interface{}, error) {
	attributes := ov.Attributes()
	result := make(map[string]interface{}, len(attributes))
	for key, elem := range attributes {
		val, err := serializeAttrValueSimple(elem)
		if err != nil {
			return nil, err
		}
		result[key] = val
	}
	return result, nil
}

// serializeAttrValueSimple serializes an attr.Value element (for use inside lists/maps/objects).
func serializeAttrValueSimple(val attr.Value) (interface{}, error) {
	if val.IsNull() || val.IsUnknown() {
		return nil, nil
	}

	switch v := val.(type) {
	case basetypes.StringValue:
		return v.ValueString(), nil
	case basetypes.Int64Value:
		return v.ValueInt64(), nil
	case basetypes.BoolValue:
		return v.ValueBool(), nil
	case basetypes.Float64Value:
		return v.ValueFloat64(), nil
	case basetypes.NumberValue:
		f, _ := v.ValueBigFloat().Float64()
		return f, nil
	case basetypes.ListValue:
		return serializeListValue(v)
	case basetypes.SetValue:
		return serializeSetValue(v)
	case basetypes.MapValue:
		return serializeMapValue(v)
	case basetypes.ObjectValue:
		return serializeObjectValue(v)
	}

	if sv, ok := val.(stringValuer); ok {
		return sv.ValueString(), nil
	}

	return nil, fmt.Errorf("statecodec: unsupported element type %T", val)
}

// serializeSlice handles Go slices (e.g., []types.String from *[]types.String fields).
func serializeSlice(v reflect.Value) (interface{}, error) {
	if v.IsNil() {
		return nil, nil
	}

	result := make([]interface{}, 0, v.Len())
	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i)
		val, err := serializeValue(elem)
		if err != nil {
			return nil, err
		}
		result = append(result, val)
	}
	return result, nil
}

// extractListValue tries to find a basetypes.ListValue embedded in a struct.
// This handles customfield.List and customfield.NestedObjectList which embed ListValue.
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

		// Direct match
		if lv, ok := fieldVal.Interface().(basetypes.ListValue); ok {
			return lv, true
		}

		// Recurse into embedded structs
		if field.Anonymous && fieldVal.Kind() == reflect.Struct {
			if lv, found := extractListValue(fieldVal); found {
				return lv, true
			}
		}
	}

	return basetypes.ListValue{}, false
}

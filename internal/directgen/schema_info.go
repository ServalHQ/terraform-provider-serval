package directgen

import (
	"context"
	"strings"

	"github.com/ServalHQ/terraform-provider-serval/internal/hclcodec"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// extractSchemaInfo walks a resource schema and builds an hclcodec.SchemaInfo
// that tells the HCL codec:
//   - which nested fields use attribute syntax (name = { ... }) vs block syntax (name { ... })
//   - which string fields have enum validators, and what their allowed values are
func extractSchemaInfo(s schema.Schema) hclcodec.SchemaInfo {
	info := make(hclcodec.SchemaInfo)
	extractFromAttributes(s.Attributes, info)
	return info
}

// extractFromAttributes recursively walks schema attributes looking for nested
// types and string enum validators.
func extractFromAttributes(attrs map[string]schema.Attribute, info hclcodec.SchemaInfo) {
	for name, attr := range attrs {
		switch a := attr.(type) {
		case schema.SingleNestedAttribute:
			entry := hclcodec.FieldSchema{NestedMode: hclcodec.NestedModeAttr}
			if len(a.Attributes) > 0 {
				entry.Children = make(hclcodec.SchemaInfo)
				extractFromAttributes(a.Attributes, entry.Children)
			}
			info[name] = entry

		case schema.ListNestedAttribute:
			entry := hclcodec.FieldSchema{NestedMode: hclcodec.NestedModeAttr}
			if len(a.NestedObject.Attributes) > 0 {
				entry.Children = make(hclcodec.SchemaInfo)
				extractFromAttributes(a.NestedObject.Attributes, entry.Children)
			}
			info[name] = entry

		case schema.StringAttribute:
			if vals := extractEnumValues(a.Validators); len(vals) > 0 {
				info[name] = hclcodec.FieldSchema{AllowedValues: vals}
			}
		}
	}
}

// extractEnumValues extracts the allowed values from OneOf / OneOfCaseInsensitive
// string validators by parsing their Description output.
//
// Both validators produce a description like:
//
//	value must be one of: ["VAL1" "VAL2" "VAL3"]
//
// We extract the bracket-delimited list and strip quotes from each element.
func extractEnumValues(validators []validator.String) []string {
	ctx := context.Background()
	for _, v := range validators {
		desc := v.Description(ctx)

		start := strings.Index(desc, "[")
		end := strings.LastIndex(desc, "]")
		if start == -1 || end == -1 || end <= start {
			continue
		}

		inner := desc[start+1 : end]
		var vals []string
		for _, field := range strings.Fields(inner) {
			val := strings.Trim(field, `"`)
			if val != "" {
				vals = append(vals, val)
			}
		}
		if len(vals) > 0 {
			return vals
		}
	}
	return nil
}

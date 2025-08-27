// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package team

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/serval-terraform/internal/customfield"
)

var _ resource.ResourceWithConfigValidators = (*TeamResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"description": schema.StringAttribute{
				Optional: true,
			},
			"name": schema.StringAttribute{
				Optional: true,
			},
			"prefix": schema.StringAttribute{
				Optional: true,
			},
			"user_ids": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},
			"created_at": schema.StringAttribute{
				Description: "A Timestamp represents a point in time independent of any time zone or local\n calendar, encoded as a count of seconds and fractions of seconds at\n nanosecond resolution. The count is relative to an epoch at UTC midnight on\n January 1, 1970, in the proleptic Gregorian calendar which extends the\n Gregorian calendar backwards to year one.\n\n All minutes are 60 seconds long. Leap seconds are \"smeared\" so that no leap\n second table is needed for interpretation, using a [24-hour linear\n smear](https://developers.google.com/time/smear).\n\n The range is from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59.999999999Z. By\n restricting to that range, we ensure that we can convert to and from [RFC\n 3339](https://www.ietf.org/rfc/rfc3339.txt) date strings.\n\n # Examples\n\n Example 1: Compute Timestamp from POSIX `time()`.\n\n     Timestamp timestamp;\n     timestamp.set_seconds(time(NULL));\n     timestamp.set_nanos(0);\n\n Example 2: Compute Timestamp from POSIX `gettimeofday()`.\n\n     struct timeval tv;\n     gettimeofday(&tv, NULL);\n\n     Timestamp timestamp;\n     timestamp.set_seconds(tv.tv_sec);\n     timestamp.set_nanos(tv.tv_usec * 1000);\n\n Example 3: Compute Timestamp from Win32 `GetSystemTimeAsFileTime()`.\n\n     FILETIME ft;\n     GetSystemTimeAsFileTime(&ft);\n     UINT64 ticks = (((UINT64)ft.dwHighDateTime) << 32) | ft.dwLowDateTime;\n\n     // A Windows tick is 100 nanoseconds. Windows epoch 1601-01-01T00:00:00Z\n     // is 11644473600 seconds before Unix epoch 1970-01-01T00:00:00Z.\n     Timestamp timestamp;\n     timestamp.set_seconds((INT64) ((ticks / 10000000) - 11644473600LL));\n     timestamp.set_nanos((INT32) ((ticks % 10000000) * 100));\n\n Example 4: Compute Timestamp from Java `System.currentTimeMillis()`.\n\n     long millis = System.currentTimeMillis();\n\n     Timestamp timestamp = Timestamp.newBuilder().setSeconds(millis / 1000)\n         .setNanos((int) ((millis % 1000) * 1000000)).build();\n\n Example 5: Compute Timestamp from Java `Instant.now()`.\n\n     Instant now = Instant.now();\n\n     Timestamp timestamp =\n         Timestamp.newBuilder().setSeconds(now.getEpochSecond())\n             .setNanos(now.getNano()).build();\n\n Example 6: Compute Timestamp from current time in Python.\n\n     timestamp = Timestamp()\n     timestamp.GetCurrentTime()\n\n # JSON Mapping\n\n In JSON format, the Timestamp type is encoded as a string in the\n [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format. That is, the\n format is \"{year}-{month}-{day}T{hour}:{min}:{sec}[.{frac_sec}]Z\"\n where {year} is always expressed using four digits while {month}, {day},\n {hour}, {min}, and {sec} are zero-padded to two digits each. The fractional\n seconds, which can go up to 9 digits (i.e. up to 1 nanosecond resolution),\n are optional. The \"Z\" suffix indicates the timezone (\"UTC\"); the timezone\n is required. A proto3 JSON serializer should always use UTC (as indicated by\n \"Z\") when printing the Timestamp type and a proto3 JSON parser should be\n able to accept both UTC and other timezones (as indicated by an offset).\n\n For example, \"2017-01-15T01:30:15.01Z\" encodes 15.01 seconds past\n 01:30 UTC on January 15, 2017.\n\n In JavaScript, one can convert a Date object to this format using the\n standard\n [toISOString()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString)\n method. In Python, a standard `datetime.datetime` object can be converted\n to this format using\n [`strftime`](https://docs.python.org/2/library/time.html#time.strftime) with\n the time format spec '%Y-%m-%dT%H:%M:%S.%fZ'. Likewise, in Java, one can use\n the Joda Time's [`ISODateTimeFormat.dateTime()`](\n http://joda-time.sourceforge.net/apidocs/org/joda/time/format/ISODateTimeFormat.html#dateTime()\n ) to obtain a formatter capable of generating timestamps in this format.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"organization_id": schema.StringAttribute{
				Computed: true,
			},
			"users": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[TeamUsersModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"avatar_url": schema.StringAttribute{
							Computed: true,
						},
						"created_at": schema.StringAttribute{
							Description: "A Timestamp represents a point in time independent of any time zone or local\n calendar, encoded as a count of seconds and fractions of seconds at\n nanosecond resolution. The count is relative to an epoch at UTC midnight on\n January 1, 1970, in the proleptic Gregorian calendar which extends the\n Gregorian calendar backwards to year one.\n\n All minutes are 60 seconds long. Leap seconds are \"smeared\" so that no leap\n second table is needed for interpretation, using a [24-hour linear\n smear](https://developers.google.com/time/smear).\n\n The range is from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59.999999999Z. By\n restricting to that range, we ensure that we can convert to and from [RFC\n 3339](https://www.ietf.org/rfc/rfc3339.txt) date strings.\n\n # Examples\n\n Example 1: Compute Timestamp from POSIX `time()`.\n\n     Timestamp timestamp;\n     timestamp.set_seconds(time(NULL));\n     timestamp.set_nanos(0);\n\n Example 2: Compute Timestamp from POSIX `gettimeofday()`.\n\n     struct timeval tv;\n     gettimeofday(&tv, NULL);\n\n     Timestamp timestamp;\n     timestamp.set_seconds(tv.tv_sec);\n     timestamp.set_nanos(tv.tv_usec * 1000);\n\n Example 3: Compute Timestamp from Win32 `GetSystemTimeAsFileTime()`.\n\n     FILETIME ft;\n     GetSystemTimeAsFileTime(&ft);\n     UINT64 ticks = (((UINT64)ft.dwHighDateTime) << 32) | ft.dwLowDateTime;\n\n     // A Windows tick is 100 nanoseconds. Windows epoch 1601-01-01T00:00:00Z\n     // is 11644473600 seconds before Unix epoch 1970-01-01T00:00:00Z.\n     Timestamp timestamp;\n     timestamp.set_seconds((INT64) ((ticks / 10000000) - 11644473600LL));\n     timestamp.set_nanos((INT32) ((ticks % 10000000) * 100));\n\n Example 4: Compute Timestamp from Java `System.currentTimeMillis()`.\n\n     long millis = System.currentTimeMillis();\n\n     Timestamp timestamp = Timestamp.newBuilder().setSeconds(millis / 1000)\n         .setNanos((int) ((millis % 1000) * 1000000)).build();\n\n Example 5: Compute Timestamp from Java `Instant.now()`.\n\n     Instant now = Instant.now();\n\n     Timestamp timestamp =\n         Timestamp.newBuilder().setSeconds(now.getEpochSecond())\n             .setNanos(now.getNano()).build();\n\n Example 6: Compute Timestamp from current time in Python.\n\n     timestamp = Timestamp()\n     timestamp.GetCurrentTime()\n\n # JSON Mapping\n\n In JSON format, the Timestamp type is encoded as a string in the\n [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format. That is, the\n format is \"{year}-{month}-{day}T{hour}:{min}:{sec}[.{frac_sec}]Z\"\n where {year} is always expressed using four digits while {month}, {day},\n {hour}, {min}, and {sec} are zero-padded to two digits each. The fractional\n seconds, which can go up to 9 digits (i.e. up to 1 nanosecond resolution),\n are optional. The \"Z\" suffix indicates the timezone (\"UTC\"); the timezone\n is required. A proto3 JSON serializer should always use UTC (as indicated by\n \"Z\") when printing the Timestamp type and a proto3 JSON parser should be\n able to accept both UTC and other timezones (as indicated by an offset).\n\n For example, \"2017-01-15T01:30:15.01Z\" encodes 15.01 seconds past\n 01:30 UTC on January 15, 2017.\n\n In JavaScript, one can convert a Date object to this format using the\n standard\n [toISOString()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString)\n method. In Python, a standard `datetime.datetime` object can be converted\n to this format using\n [`strftime`](https://docs.python.org/2/library/time.html#time.strftime) with\n the time format spec '%Y-%m-%dT%H:%M:%S.%fZ'. Likewise, in Java, one can use\n the Joda Time's [`ISODateTimeFormat.dateTime()`](\n http://joda-time.sourceforge.net/apidocs/org/joda/time/format/ISODateTimeFormat.html#dateTime()\n ) to obtain a formatter capable of generating timestamps in this format.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"deactivated_at": schema.StringAttribute{
							Description: "A Timestamp represents a point in time independent of any time zone or local\n calendar, encoded as a count of seconds and fractions of seconds at\n nanosecond resolution. The count is relative to an epoch at UTC midnight on\n January 1, 1970, in the proleptic Gregorian calendar which extends the\n Gregorian calendar backwards to year one.\n\n All minutes are 60 seconds long. Leap seconds are \"smeared\" so that no leap\n second table is needed for interpretation, using a [24-hour linear\n smear](https://developers.google.com/time/smear).\n\n The range is from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59.999999999Z. By\n restricting to that range, we ensure that we can convert to and from [RFC\n 3339](https://www.ietf.org/rfc/rfc3339.txt) date strings.\n\n # Examples\n\n Example 1: Compute Timestamp from POSIX `time()`.\n\n     Timestamp timestamp;\n     timestamp.set_seconds(time(NULL));\n     timestamp.set_nanos(0);\n\n Example 2: Compute Timestamp from POSIX `gettimeofday()`.\n\n     struct timeval tv;\n     gettimeofday(&tv, NULL);\n\n     Timestamp timestamp;\n     timestamp.set_seconds(tv.tv_sec);\n     timestamp.set_nanos(tv.tv_usec * 1000);\n\n Example 3: Compute Timestamp from Win32 `GetSystemTimeAsFileTime()`.\n\n     FILETIME ft;\n     GetSystemTimeAsFileTime(&ft);\n     UINT64 ticks = (((UINT64)ft.dwHighDateTime) << 32) | ft.dwLowDateTime;\n\n     // A Windows tick is 100 nanoseconds. Windows epoch 1601-01-01T00:00:00Z\n     // is 11644473600 seconds before Unix epoch 1970-01-01T00:00:00Z.\n     Timestamp timestamp;\n     timestamp.set_seconds((INT64) ((ticks / 10000000) - 11644473600LL));\n     timestamp.set_nanos((INT32) ((ticks % 10000000) * 100));\n\n Example 4: Compute Timestamp from Java `System.currentTimeMillis()`.\n\n     long millis = System.currentTimeMillis();\n\n     Timestamp timestamp = Timestamp.newBuilder().setSeconds(millis / 1000)\n         .setNanos((int) ((millis % 1000) * 1000000)).build();\n\n Example 5: Compute Timestamp from Java `Instant.now()`.\n\n     Instant now = Instant.now();\n\n     Timestamp timestamp =\n         Timestamp.newBuilder().setSeconds(now.getEpochSecond())\n             .setNanos(now.getNano()).build();\n\n Example 6: Compute Timestamp from current time in Python.\n\n     timestamp = Timestamp()\n     timestamp.GetCurrentTime()\n\n # JSON Mapping\n\n In JSON format, the Timestamp type is encoded as a string in the\n [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format. That is, the\n format is \"{year}-{month}-{day}T{hour}:{min}:{sec}[.{frac_sec}]Z\"\n where {year} is always expressed using four digits while {month}, {day},\n {hour}, {min}, and {sec} are zero-padded to two digits each. The fractional\n seconds, which can go up to 9 digits (i.e. up to 1 nanosecond resolution),\n are optional. The \"Z\" suffix indicates the timezone (\"UTC\"); the timezone\n is required. A proto3 JSON serializer should always use UTC (as indicated by\n \"Z\") when printing the Timestamp type and a proto3 JSON parser should be\n able to accept both UTC and other timezones (as indicated by an offset).\n\n For example, \"2017-01-15T01:30:15.01Z\" encodes 15.01 seconds past\n 01:30 UTC on January 15, 2017.\n\n In JavaScript, one can convert a Date object to this format using the\n standard\n [toISOString()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString)\n method. In Python, a standard `datetime.datetime` object can be converted\n to this format using\n [`strftime`](https://docs.python.org/2/library/time.html#time.strftime) with\n the time format spec '%Y-%m-%dT%H:%M:%S.%fZ'. Likewise, in Java, one can use\n the Joda Time's [`ISODateTimeFormat.dateTime()`](\n http://joda-time.sourceforge.net/apidocs/org/joda/time/format/ISODateTimeFormat.html#dateTime()\n ) to obtain a formatter capable of generating timestamps in this format.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"email": schema.StringAttribute{
							Computed: true,
						},
						"first_name": schema.StringAttribute{
							Computed: true,
						},
						"last_name": schema.StringAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"role": schema.StringAttribute{
							Description: `Available values: "USER_ROLE_UNSPECIFIED", "USER_ROLE_ORG_MEMBER", "USER_ROLE_ORG_ADMIN".`,
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"USER_ROLE_UNSPECIFIED",
									"USER_ROLE_ORG_MEMBER",
									"USER_ROLE_ORG_ADMIN",
								),
							},
						},
					},
				},
			},
		},
	}
}

func (r *TeamResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *TeamResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}

// Code generated by generators/resource/main.go; DO NOT EDIT.

package networkfirewall

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_networkfirewall_logging_configuration", loggingConfigurationResourceType)
}

// loggingConfigurationResourceType returns the Terraform awscc_networkfirewall_logging_configuration resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::NetworkFirewall::LoggingConfiguration resource type.
func loggingConfigurationResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"firewall_arn": {
			// Property: FirewallArn
			// CloudFormation resource type schema:
			// {
			//   "description": "A resource ARN.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "^arn:aws.*$",
			//   "type": "string"
			// }
			Description: "A resource ARN.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 256),
				validate.StringMatch(regexp.MustCompile("^arn:aws.*$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"firewall_name": {
			// Property: FirewallName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "pattern": "^[a-zA-Z0-9-]+$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 128),
				validate.StringMatch(regexp.MustCompile("^[a-zA-Z0-9-]+$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
				tfsdk.RequiresReplace(),
			},
		},
		"logging_configuration": {
			// Property: LoggingConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "LogDestinationConfigs": {
			//       "insertionOrder": false,
			//       "items": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "LogDestination": {
			//             "additionalProperties": false,
			//             "description": "A key-value pair to configure the logDestinations.",
			//             "patternProperties": {
			//               "": {
			//                 "maxLength": 1024,
			//                 "minLength": 1,
			//                 "type": "string"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "LogDestinationType": {
			//             "enum": [
			//               "S3",
			//               "CloudWatchLogs",
			//               "KinesisDataFirehose"
			//             ],
			//             "type": "string"
			//           },
			//           "LogType": {
			//             "enum": [
			//               "ALERT",
			//               "FLOW"
			//             ],
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "LogType",
			//           "LogDestinationType",
			//           "LogDestination"
			//         ],
			//         "type": "object"
			//       },
			//       "minItems": 1,
			//       "type": "array"
			//     }
			//   },
			//   "required": [
			//     "LogDestinationConfigs"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"log_destination_configs": {
						// Property: LogDestinationConfigs
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"log_destination": {
									// Property: LogDestination
									Description: "A key-value pair to configure the logDestinations.",
									// Pattern: ""
									Type:     types.MapType{ElemType: types.StringType},
									Required: true,
								},
								"log_destination_type": {
									// Property: LogDestinationType
									Type:     types.StringType,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringInSlice([]string{
											"S3",
											"CloudWatchLogs",
											"KinesisDataFirehose",
										}),
									},
								},
								"log_type": {
									// Property: LogType
									Type:     types.StringType,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringInSlice([]string{
											"ALERT",
											"FLOW",
										}),
									},
								},
							},
							tfsdk.ListNestedAttributesOptions{},
						),
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayLenAtLeast(1),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							Multiset(),
						},
					},
				},
			),
			Required: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			tfsdk.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource type definition for AWS::NetworkFirewall::LoggingConfiguration",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::NetworkFirewall::LoggingConfiguration").WithTerraformTypeName("awscc_networkfirewall_logging_configuration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"firewall_arn":            "FirewallArn",
		"firewall_name":           "FirewallName",
		"log_destination":         "LogDestination",
		"log_destination_configs": "LogDestinationConfigs",
		"log_destination_type":    "LogDestinationType",
		"log_type":                "LogType",
		"logging_configuration":   "LoggingConfiguration",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}

// Code generated by generators/resource/main.go; DO NOT EDIT.

package iotcoredeviceadvisor

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_iotcoredeviceadvisor_suite_definition", suiteDefinitionResourceType)
}

// suiteDefinitionResourceType returns the Terraform awscc_iotcoredeviceadvisor_suite_definition resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::IoTCoreDeviceAdvisor::SuiteDefinition resource type.
func suiteDefinitionResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"suite_definition_arn": {
			// Property: SuiteDefinitionArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource name for the suite definition.",
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "type": "string"
			// }
			Description: "The Amazon Resource name for the suite definition.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"suite_definition_configuration": {
			// Property: SuiteDefinitionConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "DevicePermissionRoleArn": {
			//       "description": "The device permission role arn of the test suite.",
			//       "maxLength": 2048,
			//       "minLength": 20,
			//       "type": "string"
			//     },
			//     "Devices": {
			//       "description": "The devices being tested in the test suite",
			//       "items": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "CertificateArn": {
			//             "maxLength": 2048,
			//             "minLength": 20,
			//             "type": "string"
			//           },
			//           "ThingArn": {
			//             "maxLength": 2048,
			//             "minLength": 20,
			//             "type": "string"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "maxItems": 2,
			//       "minItems": 0,
			//       "type": "array"
			//     },
			//     "IntendedForQualification": {
			//       "description": "Whether the tests are intended for qualification in a suite.",
			//       "type": "boolean"
			//     },
			//     "RootGroup": {
			//       "description": "The root group of the test suite.",
			//       "maxLength": 2048,
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "SuiteDefinitionName": {
			//       "description": "The Name of the suite definition.",
			//       "maxLength": 256,
			//       "minLength": 1,
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "DevicePermissionRoleArn",
			//     "RootGroup"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"device_permission_role_arn": {
						// Property: DevicePermissionRoleArn
						Description: "The device permission role arn of the test suite.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(20, 2048),
						},
					},
					"devices": {
						// Property: Devices
						Description: "The devices being tested in the test suite",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"certificate_arn": {
									// Property: CertificateArn
									Type:     types.StringType,
									Optional: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(20, 2048),
									},
								},
								"thing_arn": {
									// Property: ThingArn
									Type:     types.StringType,
									Optional: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(20, 2048),
									},
								},
							},
							tfsdk.ListNestedAttributesOptions{},
						),
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayLenBetween(0, 2),
						},
					},
					"intended_for_qualification": {
						// Property: IntendedForQualification
						Description: "Whether the tests are intended for qualification in a suite.",
						Type:        types.BoolType,
						Optional:    true,
					},
					"root_group": {
						// Property: RootGroup
						Description: "The root group of the test suite.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 2048),
						},
					},
					"suite_definition_name": {
						// Property: SuiteDefinitionName
						Description: "The Name of the suite definition.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 256),
						},
					},
				},
			),
			Required: true,
		},
		"suite_definition_id": {
			// Property: SuiteDefinitionId
			// CloudFormation resource type schema:
			// {
			//   "description": "The unique identifier for the suite definition.",
			//   "maxLength": 36,
			//   "minLength": 12,
			//   "type": "string"
			// }
			Description: "The unique identifier for the suite definition.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"suite_definition_version": {
			// Property: SuiteDefinitionVersion
			// CloudFormation resource type schema:
			// {
			//   "description": "The suite definition version of a test suite.",
			//   "maxLength": 255,
			//   "minLength": 2,
			//   "type": "string"
			// }
			Description: "The suite definition version of a test suite.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
						},
					},
				},
				tfsdk.SetNestedAttributesOptions{},
			),
			Optional: true,
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
		Description: "An example resource schema demonstrating some basic constructs and validation rules.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTCoreDeviceAdvisor::SuiteDefinition").WithTerraformTypeName("awscc_iotcoredeviceadvisor_suite_definition")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"certificate_arn":                "CertificateArn",
		"device_permission_role_arn":     "DevicePermissionRoleArn",
		"devices":                        "Devices",
		"intended_for_qualification":     "IntendedForQualification",
		"key":                            "Key",
		"root_group":                     "RootGroup",
		"suite_definition_arn":           "SuiteDefinitionArn",
		"suite_definition_configuration": "SuiteDefinitionConfiguration",
		"suite_definition_id":            "SuiteDefinitionId",
		"suite_definition_name":          "SuiteDefinitionName",
		"suite_definition_version":       "SuiteDefinitionVersion",
		"tags":                           "Tags",
		"thing_arn":                      "ThingArn",
		"value":                          "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}

// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package refactorspaces

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_refactorspaces_route", routeDataSource)
}

// routeDataSource returns the Terraform awscc_refactorspaces_route data source.
// This Terraform data source corresponds to the CloudFormation AWS::RefactorSpaces::Route resource.
func routeDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ApplicationIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 14,
		//	  "minLength": 14,
		//	  "pattern": "^app-([0-9A-Za-z]{10}$)",
		//	  "type": "string"
		//	}
		"application_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "pattern": "^arn:(aws[a-zA-Z-]*)?:refactor-spaces:[a-zA-Z0-9\\-]+:\\w{12}:[a-zA-Z_0-9+=,.@\\-_/]+$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: DefaultRoute
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "ActivationState": {
		//	      "enum": [
		//	        "INACTIVE",
		//	        "ACTIVE"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "ActivationState"
		//	  ],
		//	  "type": "object"
		//	}
		"default_route": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ActivationState
				"activation_state": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: EnvironmentIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 14,
		//	  "minLength": 14,
		//	  "pattern": "^env-([0-9A-Za-z]{10}$)",
		//	  "type": "string"
		//	}
		"environment_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: PathResourceToId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"path_resource_to_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: RouteIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 14,
		//	  "minLength": 14,
		//	  "pattern": "^rte-([0-9A-Za-z]{10}$)",
		//	  "type": "string"
		//	}
		"route_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: RouteType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "DEFAULT",
		//	    "URI_PATH"
		//	  ],
		//	  "type": "string"
		//	}
		"route_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ServiceIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 14,
		//	  "minLength": 14,
		//	  "pattern": "^svc-([0-9A-Za-z]{10}$)",
		//	  "type": "string"
		//	}
		"service_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A label for tagging Environment resource",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "A string used to identify this tag",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "A string containing the value for the tag",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "A string used to identify this tag",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "A string containing the value for the tag",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: UriPathRoute
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "ActivationState": {
		//	      "enum": [
		//	        "INACTIVE",
		//	        "ACTIVE"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "IncludeChildPaths": {
		//	      "type": "boolean"
		//	    },
		//	    "Methods": {
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "enum": [
		//	          "DELETE",
		//	          "GET",
		//	          "HEAD",
		//	          "OPTIONS",
		//	          "PATCH",
		//	          "POST",
		//	          "PUT"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "type": "array"
		//	    },
		//	    "SourcePath": {
		//	      "maxLength": 2048,
		//	      "minLength": 1,
		//	      "pattern": "^(/[a-zA-Z0-9._-]+)+$",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "ActivationState"
		//	  ],
		//	  "type": "object"
		//	}
		"uri_path_route": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ActivationState
				"activation_state": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: IncludeChildPaths
				"include_child_paths": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Methods
				"methods": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: SourcePath
				"source_path": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::RefactorSpaces::Route",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::RefactorSpaces::Route").WithTerraformTypeName("awscc_refactorspaces_route")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"activation_state":       "ActivationState",
		"application_identifier": "ApplicationIdentifier",
		"arn":                    "Arn",
		"default_route":          "DefaultRoute",
		"environment_identifier": "EnvironmentIdentifier",
		"include_child_paths":    "IncludeChildPaths",
		"key":                    "Key",
		"methods":                "Methods",
		"path_resource_to_id":    "PathResourceToId",
		"route_identifier":       "RouteIdentifier",
		"route_type":             "RouteType",
		"service_identifier":     "ServiceIdentifier",
		"source_path":            "SourcePath",
		"tags":                   "Tags",
		"uri_path_route":         "UriPathRoute",
		"value":                  "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}

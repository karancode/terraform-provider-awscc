---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_iotwireless_wireless_device_import_task Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::IoTWireless::WirelessDeviceImportTask
---

# awscc_iotwireless_wireless_device_import_task (Data Source)

Data Source schema for AWS::IoTWireless::WirelessDeviceImportTask



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `arn` (String) Arn for Wireless Device Import Task, Returned upon successful start.
- `creation_date` (String) CreationDate for import task
- `destination_name` (String) Destination Name for import task
- `failed_imported_devices_count` (Number) Failed Imported Devices Count
- `initialized_imported_devices_count` (Number) Initialized Imported Devices Count
- `onboarded_imported_devices_count` (Number) Onboarded Imported Devices Count
- `pending_imported_devices_count` (Number) Pending Imported Devices Count
- `sidewalk` (Attributes) sidewalk contain file for created device and role (see [below for nested schema](#nestedatt--sidewalk))
- `status` (String) Status for import task
- `status_reason` (String) StatusReason for import task
- `tags` (Attributes Set) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--sidewalk"></a>
### Nested Schema for `sidewalk`

Read-Only:

- `device_creation_file` (String)
- `device_creation_file_list` (List of String) sidewalk create device's file path
- `role` (String) sidewalk role
- `sidewalk_manufacturing_sn` (String)


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.



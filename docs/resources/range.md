# b1ddi_range (Resource)

## Example Usage

```terraform
terraform {
  required_providers {
    b1ddi = {
      version = "0.1"
      source  = "infobloxopen/b1ddi"
    }
  }
}

resource "b1ddi_ip_space" "example_tf_space" {
  name = "example_tf_space"
  comment = "Example IP space created by the terraform provider"
}

resource "b1ddi_subnet" "example_tf_subnet" {
  name = "example_tf_subnet"
  space = b1ddi_ip_space.example_tf_space.id
  address = "192.168.1.0"
  cidr = 24
  comment = "Example Subnet created by the terraform provider"
}

resource "b1ddi_range" "tf_acc_test_range" {
  start = "192.168.1.15"
  end = "192.168.1.30"
  name = "example_tf_range"
  space = b1ddi_ip_space.example_tf_space.id
  comment = "Example Range created by the terraform provider"
  depends_on = [b1ddi_subnet.example_tf_subnet]
}
```

## Schema

### Required

- **end** (String) The end IP address of the range.
- **space** (String) The resource identifier.
- **start** (String) The start IP address of the range.

### Optional

- **comment** (String) The description for the range. May contain 0 to 1024 characters. Can include UTF-8.
- **dhcp_host** (String) The resource identifier.
- **dhcp_options** (Block List) The list of DHCP options. May be either a specific option or a group of options. (see [below for nested schema](#nestedblock--dhcp_options))
- **exclusion_ranges** (Block List) The list of all exclusion ranges in the scope of the range. (see [below for nested schema](#nestedblock--exclusion_ranges))
- **id** (String) The ID of this resource.
- **inheritance_parent** (String) The resource identifier.
- **inheritance_sources** (Block List, Max: 1) The DHCP inheritance configuration for the range. (see [below for nested schema](#nestedblock--inheritance_sources))
- **name** (String) The name of the range. May contain 1 to 256 characters. Can include UTF-8.
- **parent** (String) The resource identifier.
- **tags** (Map of String) The tags for the range in JSON format.
- **threshold** (Block List, Max: 1) The utilization threshold settings for the range. (see [below for nested schema](#nestedblock--threshold))

### Read-Only

- **created_at** (String) Time when the object has been created.
- **inheritance_assigned_hosts** (List of Object) The list of the inheritance assigned hosts of the object. (see [below for nested schema](#nestedatt--inheritance_assigned_hosts))
- **protocol** (String) The type of protocol (_ipv4_ or _ipv6_).
- **updated_at** (String) Time when the object has been updated. Equals to _created_at_ if not updated after creation.
- **utilization** (List of Object) The utilization statistics for the range. (see [below for nested schema](#nestedatt--utilization))

<a id="nestedblock--dhcp_options"></a>
### Nested Schema for `dhcp_options`

Optional:

- **group** (String) The resource identifier.
- **option_code** (String) The resource identifier.
- **option_value** (String) The option value.
- **type** (String) The type of item.

  Valid values are:
  * _group_
  * _option_


<a id="nestedblock--exclusion_ranges"></a>
### Nested Schema for `exclusion_ranges`

Required:

- **end** (String) The end address of the exclusion range.
- **start** (String) The start address of the exclusion range.

Optional:

- **comment** (String) The description for the exclusion range. May contain 0 to 1024 characters. Can include UTF-8.


<a id="nestedblock--inheritance_sources"></a>
### Nested Schema for `inheritance_sources`

Optional:

- **dhcp_options** (Block List, Max: 1) The inheritance configuration for the _dhcp_options_ field. (see [below for nested schema](#nestedblock--inheritance_sources--dhcp_options))

<a id="nestedblock--inheritance_sources--dhcp_options"></a>
### Nested Schema for `inheritance_sources.dhcp_options`

Optional:

- **action** (String) The inheritance setting.

  Valid values are:
  * _inherit_: Use the inherited value.
  * _block_: Don't use the inherited value.

  Defaults to _inherit_.
- **value** (Block List) The inherited DHCP option values. (see [below for nested schema](#nestedblock--inheritance_sources--dhcp_options--value))

<a id="nestedblock--inheritance_sources--dhcp_options--value"></a>
### Nested Schema for `inheritance_sources.dhcp_options.value`

Optional:

- **action** (String) The inheritance setting.

  Valid values are:
  * _inherit_: Use the inherited value.
  * _block_: Don't use the inherited value.

  Defaults to _inherit_.
- **source** (String) The resource identifier.

Read-Only:

- **display_name** (String) The human-readable display name for the object referred to by _source_.
- **value** (List of Object) The inherited value for the DHCP option. (see [below for nested schema](#nestedatt--inheritance_sources--dhcp_options--value--value))

<a id="nestedatt--inheritance_sources--dhcp_options--value--value"></a>
### Nested Schema for `inheritance_sources.dhcp_options.value.value`

Read-Only:

- **group** (String) The resource identifier.
- **option_code** (String) The resource identifier.
- **option_value** (String) The option value.
- **type** (String) The type of item.
  Valid values are:
  * _group_
  * _option_

<a id="nestedblock--threshold"></a>
### Nested Schema for `threshold`

Required:

- **enabled** (Boolean) Indicates whether the utilization threshold for IP addresses is enabled or not.
- **high** (Number) The high threshold value for the percentage of used IP addresses relative to the total IP addresses available in the scope of the object. Thresholds are inclusive in the comparison test.
- **low** (Number) The low threshold value for the percentage of used IP addresses relative to the total IP addresses available in the scope of the object. Thresholds are inclusive in the comparison test.


<a id="nestedatt--inheritance_assigned_hosts"></a>
### Nested Schema for `inheritance_assigned_hosts`

Read-Only:

- **display_name** (String) The human-readable display name for the host referred to by _ophid_.
- **host** (String) The resource identifier.
- **ophid** (String) The on-prem host ID.


<a id="nestedatt--utilization"></a>
### Nested Schema for `utilization`

Read-Only:

- **abandon_utilization** (Number) The percentage of abandoned IP addresses relative to the total IP addresses available in the scope of the object.
- **abandoned** (String) The number of IP addresses in the scope of the object which are in the abandoned state (issued by a DHCP server and then declined by the client).
- **dynamic** (String) The number of IP addresses handed out by DHCP in the scope of the object. This includes all leased addresses, fixed addresses that are defined but not currently leased and abandoned leases.
- **free** (String) The number of IP addresses available in the scope of the object.
- **static** (String) The number of defined IP addresses such as reservations or DNS records. It can be computed as _static_ = _used_ - _dynamic_.
- **total** (String) The total number of IP addresses available in the scope of the object.
- **used** (String) The number of IP addresses used in the scope of the object.
- **utilization** (Number) The percentage of used IP addresses relative to the total IP addresses available in the scope of the object.

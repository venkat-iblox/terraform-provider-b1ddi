# b1ddi_fixed_address (Resource)

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

resource "b1ddi_ip_space" "tf_example_space" {
  name    = "tf_example_space"
  comment = "This is the example IP Space created by the B1DDI terraform provider"
}

resource "b1ddi_subnet" "tf_example_subnet" {
  name    = "tf_example_subnet"
  address = "192.168.1.0"
  cidr    = 24
  space   = b1ddi_ip_space.tf_example_space.id
}

resource "b1ddi_fixed_address" "tf_example_fixed_address" {
  name        = "tf_example_fixed_address"
  address     = "192.168.1.15"
  ip_space    = b1ddi_ip_space.tf_example_space.id
  match_type  = "mac"
  match_value = "00:00:00:00:00:00"
  comment     = "This is the example Fixed Address created by the B1DDI terraform provider"
  depends_on  = [b1ddi_subnet.tf_example_subnet]
}
```

## Schema

### Required

- **address** (String) The reserved address.
- **match_type** (String) Indicates how to match the client:
  * _mac_: match the client MAC address,
  * _client_text_ or _client_hex_: match the client identifier,
  * _relay_text_ or _relay_hex_: match the circuit ID or remote ID in the DHCP relay agent option (82).
- **match_value** (String) The value to match.

### Optional

- **comment** (String) The description for the fixed address. May contain 0 to 1024 characters. Can include UTF-8.
- **dhcp_options** (Block List) The list of DHCP options. May be either a specific option or a group of options. (see [below for nested schema](#nestedblock--dhcp_options))
- **header_option_filename** (String) The configuration for header option filename field.
- **header_option_server_address** (String) The configuration for header option server address field.
- **header_option_server_name** (String) The configuration for header option server name field.
- **hostname** (String) The DHCP host name associated with this fixed address. It is of FQDN type and it defaults to empty.
- **id** (String) The ID of this resource.
- **inheritance_parent** (String) The resource identifier.
- **inheritance_sources** (Block List, Max: 1) The inheritance configuration. (see [below for nested schema](#nestedblock--inheritance_sources))
- **ip_space** (String) The resource identifier.
- **name** (String) The name of the fixed address. May contain 1 to 256 characters. Can include UTF-8.
- **parent** (String) The resource identifier.
- **tags** (Map of String) The tags for the fixed address in JSON format.

### Read-Only

- **created_at** (String) Time when the object has been created.
- **inheritance_assigned_hosts** (List of Object) The list of the inheritance assigned hosts of the object. (see [below for nested schema](#nestedatt--inheritance_assigned_hosts))
- **updated_at** (String) Time when the object has been updated. Equals to _created_at_ if not updated after creation.

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

<a id="nestedblock--inheritance_sources"></a>
### Nested Schema for `inheritance_sources`

Optional:

- **dhcp_options** (Block List, Max: 1) The inheritance configuration for _dhcp_options_ field. (see [below for nested schema](#nestedblock--inheritance_sources--dhcp_options))
- **header_option_filename** (Block List, Max: 1) The inheritance configuration for _header_option_filename_ field. (see [below for nested schema](#nestedblock--inheritance_sources--header_option_filename))
- **header_option_server_address** (Block List, Max: 1) The inheritance configuration for _header_option_server_address_ field. (see [below for nested schema](#nestedblock--inheritance_sources--header_option_server_address))
- **header_option_server_name** (Block List, Max: 1) The inheritance configuration for _header_option_server_name_ field. (see [below for nested schema](#nestedblock--inheritance_sources--header_option_server_name))

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

<a id="nestedblock--inheritance_sources--header_option_filename"></a>
### Nested Schema for `inheritance_sources.header_option_filename`

Optional:

- **action** (String) The inheritance setting for a field.

  Valid values are:
  * _inherit_: Use the inherited value.
  * _override_: Use the value set in the object.

  Defaults to _inherit_.
- **source** (String) The resource identifier.

Read-Only:

- **display_name** (String) The human-readable display name for the object referred to by _source_.
- **value** (String) The inherited value.


<a id="nestedblock--inheritance_sources--header_option_server_address"></a>
### Nested Schema for `inheritance_sources.header_option_server_address`

Optional:

- **action** (String) The inheritance setting for a field.

  Valid values are:
  * _inherit_: Use the inherited value.
  * _override_: Use the value set in the object.

  Defaults to _inherit_.
- **source** (String) The resource identifier.

Read-Only:

- **display_name** (String) The human-readable display name for the object referred to by _source_.
- **value** (String) The inherited value.


<a id="nestedblock--inheritance_sources--header_option_server_name"></a>
### Nested Schema for `inheritance_sources.header_option_server_name`

Optional:

- **action** (String) The inheritance setting for a field.

  Valid values are:
  * _inherit_: Use the inherited value.
  * _override_: Use the value set in the object.

  Defaults to _inherit_.
- **source** (String) The resource identifier.

Read-Only:

- **display_name** (String) The human-readable display name for the object referred to by _source_.
- **value** (String) The inherited value.

<a id="nestedatt--inheritance_assigned_hosts"></a>
### Nested Schema for `inheritance_assigned_hosts`

Read-Only:

- **display_name** (String) The human-readable display name for the host referred to by _ophid_.
- **host** (String) The resource identifier.
- **ophid** (String) The on-prem host ID.

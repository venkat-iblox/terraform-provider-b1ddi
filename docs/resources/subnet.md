# b1ddi_subnet (Resource)

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
  comment = "Example IP space for the terraform provider"
}

resource "b1ddi_subnet" "example_tf_subnet" {
  name = "example_tf_subnet"
  space = b1ddi_ip_space.example_tf_space.id
  address = "192.168.1.0"
  cidr = 24
  comment = "Example Subnet created by the terraform provider"
}
```

## Schema

### Required

- **address** (String) The address of the subnet in the form “a.b.c.d/n” where the “/n” may be omitted. In this case, the CIDR value must be defined in the _cidr_ field. When reading, the _address_ field is always in the form “a.b.c.d”.
- **space** (String) The resource identifier.

### Optional

- **asm_config** (Block List, Max: 1) The Automated Scope Management configuration for the subnet. (see [below for nested schema](#nestedblock--asm_config))
- **cidr** (Number) The CIDR of the subnet. This is required if _address_ does not include CIDR.
- **comment** (String) The description for the subnet. May contain 0 to 1024 characters. Can include UTF-8.
- **ddns_client_update** (String) Controls who does the DDNS updates.

  Valid values are:
  * _client_: DHCP server updates DNS if requested by client.
  * _server_: DHCP server always updates DNS, overriding an update request from the client, unless the client requests no updates.
  * _ignore_: DHCP server always updates DNS, even if the client says not to.
  * _over_client_update_: Same as _server_. DHCP server always updates DNS, overriding an update request from the client, unless the client requests no updates.
  * _over_no_update_: DHCP server updates DNS even if the client requests that no updates be done. If the client requests to do the update, DHCP server allows it.

  Defaults to _client_.
- **ddns_domain** (String) The domain suffix for DDNS updates. FQDN, may be empty.
  Defaults to empty.
- **ddns_generate_name** (Boolean) Indicates if DDNS needs to generate a hostname when not supplied by the client.
  Defaults to _false_.
- **ddns_generated_prefix** (String) The prefix used in the generation of an FQDN.

  When generating a name, DHCP server will construct the name in the format: [ddns-generated-prefix]-[address-text].[ddns-qualifying-suffix].
  where address-text is simply the lease IP address converted to a hyphenated string.

  Defaults to "myhost".
- **ddns_send_updates** (Boolean) Determines if DDNS updates are enabled at the subnet level.
  Defaults to _true_.
- **ddns_update_on_renew** (Boolean) Instructs the DHCP server to always update the DNS information when a lease is renewed even if its DNS information has not changed.
  Defaults to _false_.
- **ddns_use_conflict_resolution** (Boolean) When true, DHCP server will apply conflict resolution, as described in RFC 4703, when attempting to fulfill the update request.

  When false, DHCP server will simply attempt to update the DNS entries per the request, regardless of whether or not they conflict with existing entries owned by other DHCP4 clients.

  Defaults to _true_.
- **dhcp_config** (Block List, Max: 1) The DHCP configuration of the subnet that controls how leases are issued. (see [below for nested schema](#nestedblock--dhcp_config))
- **dhcp_host** (String) The resource identifier.
- **dhcp_options** (Block List) The DHCP options of the subnet. This can either be a specific option or a group of options. (see [below for nested schema](#nestedblock--dhcp_options))
- **header_option_filename** (String) The configuration for header option filename field.
- **header_option_server_address** (String) The configuration for header option server address field.
- **header_option_server_name** (String) The configuration for header option server name field.
- **hostname_rewrite_char** (String) The character to replace non-matching characters with, when hostname rewrite is enabled.

  Any single ASCII character.

  Defaults to "_".
- **hostname_rewrite_enabled** (Boolean) Indicates if client supplied hostnames will be rewritten prior to DDNS update by replacing every character that does not match _hostname_rewrite_regex_ by _hostname_rewrite_char_.
  Defaults to _false_.
- **hostname_rewrite_regex** (String) The regex bracket expression to match valid characters.

  Must begin with "[" and end with "]" and be a compilable POSIX regex.

  Defaults to "[^a-zA-Z0-9_.]".
- **id** (String) The ID of this resource.
- **inheritance_sources** (Block List, Max: 1) The DHCP inheritance configuration for the subnet. (see [below for nested schema](#nestedblock--inheritance_sources))
- **name** (String) The name of the subnet. May contain 1 to 256 characters. Can include UTF-8.
- **tags** (Map of String) The tags for the subnet in JSON format.
- **threshold** (Block List, Max: 1) The IP address utilization threshold settings for the subnet. (see [below for nested schema](#nestedblock--threshold))

### Read-Only

- **asm_scope_flag** (Number) Set to 1 to indicate that the subnet may run out of addresses.
- **created_at** (String) Time when the object has been created.
- **dhcp_utilization** (List of Object) The utilization of IP addresses within the DHCP ranges of the subnet. (see [below for nested schema](#nestedatt--dhcp_utilization))
- **inheritance_assigned_hosts** (List of Object) The list of the inheritance assigned hosts of the object. (see [below for nested schema](#nestedatt--inheritance_assigned_hosts))
- **inheritance_parent** (String) The resource identifier.
- **parent** (String) The resource identifier.
- **protocol** (String) The type of protocol of the subnet (_ipv4_ or _ipv6_).
- **updated_at** (String) Time when the object has been updated. Equals to _created_at_ if not updated after creation.
- **utilization** (List of Object) The IP address utilization statistics of the subnet. (see [below for nested schema](#nestedatt--utilization))

<a id="nestedblock--asm_config"></a>
### Nested Schema for `asm_config`

Optional:

- **asm_threshold** (Number) ASM shows the number of addresses forecast to be used _forecast_period_ days in the future, if it is greater than _asm_threshold_ percent * _dhcp_total_ (see _dhcp_utilization_) then the subnet is flagged.
- **enable** (Boolean) Indicates if Automated Scope Management is enabled.
- **enable_notification** (Boolean) Indicates if ASM should send notifications to the user.
- **forecast_period** (Number) The forecast period in days.
- **growth_factor** (Number) Indicates the growth in the number or percentage of IP addresses.
- **growth_type** (String) The type of factor to use: _percent_ or _count_.
- **history** (Number) The minimum amount of history needed before ASM can run on this subnet.
- **min_total** (Number) The minimum size of range needed for ASM to run on this subnet.
- **min_unused** (Number) The minimum percentage of addresses that must be available outside of the DHCP ranges and fixed addresses
  when making a suggested change..
- **reenable_date** (String)


<a id="nestedblock--dhcp_config"></a>
### Nested Schema for `dhcp_config`

Optional:

- **allow_unknown** (Boolean) Disable to allow leases only for known clients, those for which a fixed address is configured.
- **filters** (List of String) The resource identifier.
- **ignore_list** (Block List) The list of clients to ignore requests from. (see [below for nested schema](#nestedblock--dhcp_config--ignore_list))
- **lease_time** (Number) The lease duration in seconds.

<a id="nestedblock--dhcp_config--ignore_list"></a>
### Nested Schema for `dhcp_config.ignore_list`

Required:

- **type** (String) Type of ignore matching: client to ignore by client identifier (client hex or client text) or hardware to ignore by hardware identifier (MAC address). 
  
  It can have one of the following values:
  * _client_hex_,
  * _client_text_,
  * _hardware_.
- **value** (String) Value to match.

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

- **asm_config** (Block List, Max: 1) The inheritance configuration for _asm_config_ field. (see [below for nested schema](#nestedblock--inheritance_sources--asm_config))
- **ddns_client_update** (Block List, Max: 1) The inheritance configuration for _ddns_client_update_ field. (see [below for nested schema](#nestedblock--inheritance_sources--ddns_client_update))
- **ddns_enabled** (Block List, Max: 1) The inheritance configuration for _ddns_enabled_ field. Only action allowed is 'inherit'. (see [below for nested schema](#nestedblock--inheritance_sources--ddns_enabled))
- **ddns_hostname_block** (Block List, Max: 1) The inheritance configuration for _ddns_generate_name_ and _ddns_generated_prefix_ fields. (see [below for nested schema](#nestedblock--inheritance_sources--ddns_hostname_block))
- **ddns_update_block** (Block List, Max: 1) The inheritance configuration for _ddns_send_updates_ and _ddns_domain_ fields. (see [below for nested schema](#nestedblock--inheritance_sources--ddns_update_block))
- **ddns_update_on_renew** (Block List, Max: 1) The inheritance configuration for _ddns_update_on_renew_ field. (see [below for nested schema](#nestedblock--inheritance_sources--ddns_update_on_renew))
- **ddns_use_conflict_resolution** (Block List, Max: 1) The inheritance configuration for _ddns_use_conflict_resolution_ field. (see [below for nested schema](#nestedblock--inheritance_sources--ddns_use_conflict_resolution))
- **dhcp_config** (Block List, Max: 1) The inheritance configuration for _dhcp_config_ field. (see [below for nested schema](#nestedblock--inheritance_sources--dhcp_config))
- **dhcp_options** (Block List, Max: 1) The inheritance configuration for _dhcp_options_ field. (see [below for nested schema](#nestedblock--inheritance_sources--dhcp_options))
- **header_option_filename** (Block List, Max: 1) The inheritance configuration for _header_option_filename_ field. (see [below for nested schema](#nestedblock--inheritance_sources--header_option_filename))
- **header_option_server_address** (Block List, Max: 1) The inheritance configuration for _header_option_server_address_ field. (see [below for nested schema](#nestedblock--inheritance_sources--header_option_server_address))
- **header_option_server_name** (Block List, Max: 1) The inheritance configuration for _header_option_server_name_ field. (see [below for nested schema](#nestedblock--inheritance_sources--header_option_server_name))
- **hostname_rewrite_block** (Block List, Max: 1) The inheritance configuration for _hostname_rewrite_enabled_, _hostname_rewrite_regex_, and _hostname_rewrite_char_ fields. (see [below for nested schema](#nestedblock--inheritance_sources--hostname_rewrite_block))

<a id="nestedblock--inheritance_sources--asm_config"></a>
### Nested Schema for `inheritance_sources.asm_config`

Optional:

- **asm_enable_block** (Block List, Max: 1) The block of ASM fields: _enable_, _enable_notification_, _reenable_date_. (see [below for nested schema](#nestedblock--inheritance_sources--asm_config--asm_enable_block))
- **asm_growth_block** (Block List, Max: 1) The block of ASM fields: _growth_factor_, _growth_type_. (see [below for nested schema](#nestedblock--inheritance_sources--asm_config--asm_growth_block))
- **asm_threshold** (Block List, Max: 1) ASM shows the number of addresses forecast to be used _forecast_period_ days in the future, if it is greater than _asm_threshold_percent_ * _dhcp_total_ (see _dhcp_utilization_) then the subnet is flagged. (see [below for nested schema](#nestedblock--inheritance_sources--asm_config--asm_threshold))
- **forecast_period** (Block List, Max: 1) The forecast period in days. (see [below for nested schema](#nestedblock--inheritance_sources--asm_config--forecast_period))
- **history** (Block List, Max: 1) The minimum amount of history needed before ASM can run on this subnet. (see [below for nested schema](#nestedblock--inheritance_sources--asm_config--history))
- **min_total** (Block List, Max: 1) The minimum size of range needed for ASM to run on this subnet. (see [below for nested schema](#nestedblock--inheritance_sources--asm_config--min_total))
- **min_unused** (Block List, Max: 1) The minimum percentage of addresses that must be available outside of the DHCP ranges and fixed addresses when making a suggested change. (see [below for nested schema](#nestedblock--inheritance_sources--asm_config--min_unused))

<a id="nestedblock--inheritance_sources--asm_config--asm_enable_block"></a>
### Nested Schema for `inheritance_sources.asm_config.asm_enable_block`

Optional:

- **action** (String) The inheritance setting.

  Valid values are:
  * _inherit_: Use the inherited value.
  * _override_: Use the value set in the object.

  Defaults to _inherit_.
- **source** (String) The resource identifier.

Read-Only:

- **display_name** (String) The human-readable display name for the object referred to by _source_.
- **value** (List of Object) The inherited value. (see [below for nested schema](#nestedatt--inheritance_sources--asm_config--asm_enable_block--value))

<a id="nestedatt--inheritance_sources--asm_config--asm_enable_block--value"></a>
### Nested Schema for `inheritance_sources.asm_config.asm_enable_block.value`

Read-Only:

- **enable** (Boolean) Indicates whether Automated Scope Management is enabled or not.
- **enable_notification** (Boolean) Indicates whether sending notifications to the users is enabled or not.
- **reenable_date** (String) The date at which notifications will be re-enabled automatically.



<a id="nestedblock--inheritance_sources--asm_config--asm_growth_block"></a>
### Nested Schema for `inheritance_sources.asm_config.asm_growth_block`

Optional:

- **action** (String) The inheritance setting.

  Valid values are:
  * _inherit_: Use the inherited value.
  * _override_: Use the value set in the object.

  Defaults to _inherit_.
- **source** (String) The resource identifier.

Read-Only:

- **display_name** (String) The human-readable display name for the object referred to by _source_.
- **value** (List of Object) The inherited value. (see [below for nested schema](#nestedatt--inheritance_sources--asm_config--asm_growth_block--value))

<a id="nestedatt--inheritance_sources--asm_config--asm_growth_block--value"></a>
### Nested Schema for `inheritance_sources.asm_config.asm_growth_block.value`

Read-Only:

- **growth_factor** (Number) Either the number or percentage of addresses to grow by.
- **growth_type** (String) The type of factor to use: _percent_ or _count_.

<a id="nestedblock--inheritance_sources--asm_config--asm_threshold"></a>
### Nested Schema for `inheritance_sources.asm_config.asm_threshold`

Optional:

- **action** (String) The inheritance setting for a field.

  Valid values are:
  * _inherit_: Use the inherited value.
  * _override_: Use the value set in the object.

  Defaults to _inherit_.
- **source** (String) The resource identifier.

Read-Only:

- **display_name** (String) The human-readable display name for the object referred to by _source_.
- **value** (Number) The inherited value.


<a id="nestedblock--inheritance_sources--asm_config--forecast_period"></a>
### Nested Schema for `inheritance_sources.asm_config.forecast_period`

Optional:

- **action** (String) The inheritance setting for a field.

  Valid values are:
  * _inherit_: Use the inherited value.
  * _override_: Use the value set in the object.

  Defaults to _inherit_.
- **source** (String) The resource identifier.

Read-Only:

- **display_name** (String) The human-readable display name for the object referred to by _source_.
- **value** (Number) The inherited value.


<a id="nestedblock--inheritance_sources--asm_config--history"></a>
### Nested Schema for `inheritance_sources.asm_config.history`

Optional:

- **action** (String) The inheritance setting for a field.

  Valid values are:
  * _inherit_: Use the inherited value.
  * _override_: Use the value set in the object.

  Defaults to _inherit_.
- **source** (String) The resource identifier.

Read-Only:

- **display_name** (String) The human-readable display name for the object referred to by _source_.
- **value** (Number) The inherited value.


<a id="nestedblock--inheritance_sources--asm_config--min_total"></a>
### Nested Schema for `inheritance_sources.asm_config.min_total`

Optional:

- **action** (String) The inheritance setting for a field.

  Valid values are:
  * _inherit_: Use the inherited value.
  * _override_: Use the value set in the object.

  Defaults to _inherit_.
- **source** (String) The resource identifier.

Read-Only:

- **display_name** (String) The human-readable display name for the object referred to by _source_.
- **value** (Number) The inherited value.


<a id="nestedblock--inheritance_sources--asm_config--min_unused"></a>
### Nested Schema for `inheritance_sources.asm_config.min_unused`

Optional:

- **action** (String) The inheritance setting for a field.

  Valid values are:
  * _inherit_: Use the inherited value.
  * _override_: Use the value set in the object.

  Defaults to _inherit_.
- **source** (String) The resource identifier.

Read-Only:

- **display_name** (String) The human-readable display name for the object referred to by _source_.
- **value** (Number) The inherited value.



<a id="nestedblock--inheritance_sources--ddns_client_update"></a>
### Nested Schema for `inheritance_sources.ddns_client_update`

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


<a id="nestedblock--inheritance_sources--ddns_enabled"></a>
### Nested Schema for `inheritance_sources.ddns_enabled`

Optional:

- **action** (String) The inheritance setting for a field.

  Valid values are:
  * _inherit_: Use the inherited value.
  * _override_: Use the value set in the object.

  Defaults to _inherit_.
- **source** (String) The resource identifier.

Read-Only:

- **display_name** (String) The human-readable display name for the object referred to by _source_.
- **value** (Boolean) The inherited value.

<a id="nestedblock--inheritance_sources--ddns_hostname_block"></a>
### Nested Schema for `inheritance_sources.ddns_hostname_block`

Optional:

- **action** (String) The inheritance setting.

  Valid values are:
  * _inherit_: Use the inherited value.
  * _override_: Use the value set in the object.

  Defaults to _inherit_.
- **source** (String) The resource identifier.

Read-Only:

- **display_name** (String) The human-readable display name for the object referred to by _source_.
- **value** (List of Object) The inherited value. (see [below for nested schema](#nestedatt--inheritance_sources--ddns_hostname_block--value))

<a id="nestedatt--inheritance_sources--ddns_hostname_block--value"></a>
### Nested Schema for `inheritance_sources.ddns_hostname_block.value`

Read-Only:

- **ddns_generate_name** (Boolean) Indicates if DDNS should generate a hostname when not supplied by the client.
- **ddns_generated_prefix** (String) The prefix used in the generation of an FQDN.

<a id="nestedblock--inheritance_sources--ddns_update_block"></a>
### Nested Schema for `inheritance_sources.ddns_update_block`

Optional:

- **action** (String) The inheritance setting.

  Valid values are:
  * _inherit_: Use the inherited value.
  * _override_: Use the value set in the object.

  Defaults to _inherit_.
- **source** (String) The resource identifier.

Read-Only:

- **display_name** (String) The human-readable display name for the object referred to by _source_.
- **value** (List of Object) The inherited value. (see [below for nested schema](#nestedatt--inheritance_sources--ddns_update_block--value))

<a id="nestedatt--inheritance_sources--ddns_update_block--value"></a>
### Nested Schema for `inheritance_sources.ddns_update_block.value`

Read-Only:

- **ddns_domain** (String) The domain name for DDNS.
- **ddns_send_updates** (Boolean) Determines if DDNS updates are enabled at this level.

<a id="nestedblock--inheritance_sources--ddns_update_on_renew"></a>
### Nested Schema for `inheritance_sources.ddns_update_on_renew`

Optional:

- **action** (String) The inheritance setting for a field.

  Valid values are:
  * _inherit_: Use the inherited value.
  * _override_: Use the value set in the object.

  Defaults to _inherit_.
- **source** (String) The resource identifier.

Read-Only:

- **display_name** (String) The human-readable display name for the object referred to by _source_.
- **value** (Boolean) The inherited value.


<a id="nestedblock--inheritance_sources--ddns_use_conflict_resolution"></a>
### Nested Schema for `inheritance_sources.ddns_use_conflict_resolution`

Optional:

- **action** (String) The inheritance setting for a field.

  Valid values are:
  * _inherit_: Use the inherited value.
  * _override_: Use the value set in the object.

  Defaults to _inherit_.
- **source** (String) The resource identifier.

Read-Only:

- **display_name** (String) The human-readable display name for the object referred to by _source_.
- **value** (Boolean) The inherited value.


<a id="nestedblock--inheritance_sources--dhcp_config"></a>
### Nested Schema for `inheritance_sources.dhcp_config`

Optional:

- **allow_unknown** (Block List, Max: 1) The inheritance configuration for _allow_unknown_ field from _DHCPConfig_ object. (see [below for nested schema](#nestedblock--inheritance_sources--dhcp_config--allow_unknown))
- **filters** (Block List, Max: 1) The inheritance configuration for filters field from _DHCPConfig_ object. (see [below for nested schema](#nestedblock--inheritance_sources--dhcp_config--filters))
- **ignore_list** (Block List, Max: 1) The inheritance configuration for _ignore_list_ field from _DHCPConfig_ object. (see [below for nested schema](#nestedblock--inheritance_sources--dhcp_config--ignore_list))
- **lease_time** (Block List, Max: 1) The inheritance configuration for _lease_time_ field from _DHCPConfig_ object. (see [below for nested schema](#nestedblock--inheritance_sources--dhcp_config--lease_time))

<a id="nestedblock--inheritance_sources--dhcp_config--allow_unknown"></a>
### Nested Schema for `inheritance_sources.dhcp_config.allow_unknown`

Optional:

- **action** (String) The inheritance setting for a field.

  Valid values are:
  * _inherit_: Use the inherited value.
  * _override_: Use the value set in the object.

  Defaults to _inherit_.
- **source** (String) The resource identifier.

Read-Only:

- **display_name** (String) The human-readable display name for the object referred to by _source_.
- **value** (Boolean) The inherited value.

<a id="nestedblock--inheritance_sources--dhcp_config--filters"></a>
### Nested Schema for `inheritance_sources.dhcp_config.filters`

Optional:

- **action** (String) The inheritance setting.

  Valid values are:
  * _inherit_: Use the inherited value.
  * _override_: Use the value set in the object.

  Defaults to _inherit_.
- **source** (String) The resource identifier.
- **value** (List of String) The resource identifier.

Read-Only:

- **display_name** (String) The human-readable display name for the object referred to by _source_.


<a id="nestedblock--inheritance_sources--dhcp_config--ignore_list"></a>
### Nested Schema for `inheritance_sources.dhcp_config.ignore_list`

Optional:

- **action** (String) The inheritance setting.

  Valid values are:
  * _inherit_: Use the inherited value.
  * _override_: Use the value set in the object.

  Defaults to _inherit_.
- **source** (String) The resource identifier.

Read-Only:

- **display_name** (String) The human-readable display name for the object referred to by _source_.
- **value** (List of Object) The inherited value. (see [below for nested schema](#nestedatt--inheritance_sources--dhcp_config--ignore_list--value))

<a id="nestedatt--inheritance_sources--dhcp_config--ignore_list--value"></a>
### Nested Schema for `inheritance_sources.dhcp_config.ignore_list.value`

Read-Only:

- **type** (String) Type of ignore matching: client to ignore by client identifier (client hex or client text) or hardware to ignore by hardware identifier (MAC address). 

  It can have one of the following values:
  * _client_hex_,
  * _client_text_,
  * _hardware_.
- **value** (String) Value to match.

<a id="nestedblock--inheritance_sources--dhcp_config--lease_time"></a>
### Nested Schema for `inheritance_sources.dhcp_config.lease_time`

Optional:

- **action** (String) The inheritance setting for a field.

  Valid values are:
  * _inherit_: Use the inherited value.
  * _override_: Use the value set in the object.

  Defaults to _inherit_.
- **source** (String) The resource identifier.

Read-Only:

- **display_name** (String) The human-readable display name for the object referred to by _source_.
- **value** (Number) The inherited value.

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

<a id="nestedblock--inheritance_sources--hostname_rewrite_block"></a>
### Nested Schema for `inheritance_sources.hostname_rewrite_block`

Optional:

- **action** (String) The inheritance setting.

  Valid values are:
  * _inherit_: Use the inherited value.
  * _override_: Use the value set in the object.

  Defaults to _inherit_.
- **source** (String) The resource identifier.

Read-Only:

- **display_name** (String) The human-readable display name for the object referred to by _source_.
- **value** (List of Object) The inherited value. (see [below for nested schema](#nestedatt--inheritance_sources--hostname_rewrite_block--value))

<a id="nestedatt--inheritance_sources--hostname_rewrite_block--value"></a>
### Nested Schema for `inheritance_sources.hostname_rewrite_block.value`

Read-Only:

- **hostname_rewrite_char** (String) The inheritance configuration for _hostname_rewrite_char_ field.
- **hostname_rewrite_enabled** (Boolean) The inheritance configuration for _hostname_rewrite_enabled_ field.
- **hostname_rewrite_regex** (String) The inheritance configuration for _hostname_rewrite_regex_ field.


<a id="nestedblock--threshold"></a>
### Nested Schema for `threshold`

Required:

- **enabled** (Boolean) Indicates whether the utilization threshold for IP addresses is enabled or not.
- **high** (Number) The high threshold value for the percentage of used IP addresses relative to the total IP addresses available in the scope of the object. Thresholds are inclusive in the comparison test.
- **low** (Number) The low threshold value for the percentage of used IP addresses relative to the total IP addresses available in the scope of the object. Thresholds are inclusive in the comparison test.


<a id="nestedatt--dhcp_utilization"></a>
### Nested Schema for `dhcp_utilization`

Read-Only:

- **dhcp_free** (String) The total free IP addresses in the DHCP ranges in the scope of this object. It can be computed as _dhcp_total_ - _dhcp_used_.
- **dhcp_total** (String) The total IP addresses available in the DHCP ranges in the scope of this object.
- **dhcp_used** (String) The total IP addresses marked as used in the DHCP ranges in the scope of this object.
- **dhcp_utilization** (Number) The percentage of used IP addresses relative to the total IP addresses available in the DHCP ranges in the scope of this object.

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

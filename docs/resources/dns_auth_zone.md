# b1ddi_dns_auth_zone (Resource)

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

variable "internal_secondary_name" {
  type = string
  description = "Internal secondary DNS Host name for the DNS Auth Zone configuration"
}

# Get DNS Host by name
data "b1ddi_dns_hosts" "dns_host_by_name" {
  filters = {
    name = var.internal_secondary_name
  }
}

resource "b1ddi_dns_view" "tf_example_dns_view" {
  name = "example_tf_dns_view"
}

resource "b1ddi_dns_auth_zone" "tf_example_auth_zone" {
  internal_secondaries {
    host = data.b1ddi_dns_hosts.dns_host_by_name.results.0.id
  }
  fqdn = "tf-example.com."
  primary_type = "cloud"
  view = b1ddi_dns_view.tf_example_dns_view.id
}
```

## Schema

### Required

- **fqdn** (String) Zone FQDN.
  The FQDN supplied at creation will be converted to canonical form.
  Read-only after creation.
- **primary_type** (String) Primary type for an authoritative zone.
  Read only after creation.
  
  Allowed values:
   * _external_: zone data owned by an external nameserver,
   * _cloud_: zone data is owned by a BloxOne DDI host.

### Optional

- **comment** (String) Optional. Comment for zone configuration.
- **disabled** (Boolean) Optional. _true_ to disable object. A disabled object is effectively non-existent when generating configuration.
- **external_primaries** (Block List) Optional. DNS primaries external to BloxOne DDI. Order is not significant. (see [below for nested schema](#nestedblock--external_primaries))
- **external_secondaries** (Block List) DNS secondaries external to BloxOne DDI. Order is not significant. (see [below for nested schema](#nestedblock--external_secondaries))
- **gss_tsig_enabled** (Boolean) _gss_tsig_enabled_ enables/disables GSS-TSIG signed dynamic updates.
  Defaults to _false_.
- **id** (String) The ID of this resource.
- **inheritance_sources** (Block List, Max: 1) Optional. Inheritance configuration. (see [below for nested schema](#nestedblock--inheritance_sources))
- **initial_soa_serial** (Number) On-create-only. SOA serial is allowed to be set when the authoritative zone is created.
- **internal_secondaries** (Block List) Optional. BloxOne DDI hosts acting as internal secondaries. Order is not significant. (see [below for nested schema](#nestedblock--internal_secondaries))
- **notify** (Boolean) Also notify all external secondary DNS servers if enabled.
  Defaults to _false_.
- **nsgs** (List of String) The resource identifier.
- **parent** (String) The resource identifier.
- **query_acl** (Block List) Optional. Clients must match this ACL to make authoritative queries.
  Also used for recursive queries if that ACL is unset.
  Defaults to empty. (see [below for nested schema](#nestedblock--query_acl))
- **tags** (Map of String) Tagging specifics.
- **transfer_acl** (Block List) Optional. Clients must match this ACL to receive zone transfers. (see [below for nested schema](#nestedblock--transfer_acl))
- **update_acl** (Block List) Optional. Specifies which hosts are allowed to submit Dynamic DNS updates for authoritative zones of _primary_type_ _cloud_.
  Defaults to empty. (see [below for nested schema](#nestedblock--update_acl))
- **use_forwarders_for_subzones** (Boolean) Optional. Use default forwarders to resolve queries for subzones.
  Defaults to _true_.
- **view** (String) The resource identifier.
- **zone_authority** (Block List, Max: 1) Optional. ZoneAuthority. (see [below for nested schema](#nestedblock--zone_authority))

### Read-Only

- **created_at** (String) Time when the object has been created.
- **inheritance_assigned_hosts** (List of Object) The list of the inheritance assigned hosts of the object. (see [below for nested schema](#nestedatt--inheritance_assigned_hosts))
- **mapped_subnet** (String) Reverse zone network address in the following format: "ip-address/cidr".
  Defaults to empty.
- **mapping** (String) Zone mapping type.
  Allowed values:
   * _forward_,
   * _ipv4_reverse_.
   * _ipv6_reverse_.

  Defaults to forward.
- **protocol_fqdn** (String) Zone FQDN in punycode.
- **updated_at** (String) Time when the object has been updated. Equals to _created_at_ if not updated after creation.

<a id="nestedblock--external_primaries"></a>
### Nested Schema for `external_primaries`

Required:

- **type** (String) Allowed values:
  * _nsg_,
  * _primary_.

Optional:

- **address** (String) Optional. Required only if _type_ is _server_. IP Address of nameserver.
- **fqdn** (String) Optional. Required only if _type_ is _server_. FQDN of nameserver.
- **nsg** (String) The resource identifier.
- **tsig_enabled** (Boolean) Optional. If enabled, secondaries will use the configured TSIG key when requesting a zone transfer from this primary.
- **tsig_key** (Block List, Max: 1) Optional. TSIG key.

  Error if empty while _tsig_enabled_ is _true_. (see [below for nested schema](#nestedblock--external_primaries--tsig_key))

Read-Only:

- **protocol_fqdn** (String) FQDN of nameserver in punycode.

<a id="nestedblock--external_primaries--tsig_key"></a>
### Nested Schema for `external_primaries.tsig_key`

Required:

- **key** (String) The resource identifier.

Optional:

- **algorithm** (String) TSIG key algorithm.

  Possible values:
   * _hmac_sha256_,
   * _hmac_sha1_,
   * _hmac_sha224_,
   * _hmac_sha384_,
   * _hmac_sha512_.
- **comment** (String) Comment for TSIG key.
- **name** (String) TSIG key name, FQDN.
- **secret** (String) TSIG key secret, base64 string.

Read-Only:

- **protocol_name** (String) TSIG key name in punycode.


<a id="nestedblock--external_secondaries"></a>
### Nested Schema for `external_secondaries`

Required:

- **address** (String) IP Address of nameserver.
- **fqdn** (String) FQDN of nameserver.

Optional:

- **stealth** (Boolean) If enabled, the NS record and glue record will NOT be automatically generated
  according to secondaries nameserver assignment.
  Default: _false_
- **tsig_enabled** (Boolean) If enabled, secondaries will use the configured TSIG key when requesting a zone transfer.
  Default: _false_
- **tsig_key** (Block List, Max: 1) TSIG key.

  Error if empty while _tsig_enabled_ is _true_. (see [below for nested schema](#nestedblock--external_secondaries--tsig_key))

Read-Only:

- **protocol_fqdn** (String) FQDN of nameserver in punycode.

<a id="nestedblock--external_secondaries--tsig_key"></a>
### Nested Schema for `external_secondaries.tsig_key`

Required:

- **key** (String) The resource identifier.

Optional:

- **algorithm** (String) TSIG key algorithm.

  Possible values:
   * _hmac_sha256_,
   * _hmac_sha1_,
   * _hmac_sha224_,
   * _hmac_sha384_,
   * _hmac_sha512_.

- **comment** (String) Comment for TSIG key.
- **name** (String) TSIG key name, FQDN.
- **secret** (String) TSIG key secret, base64 string.

Read-Only:

- **protocol_name** (String) TSIG key name in punycode.



<a id="nestedblock--inheritance_sources"></a>
### Nested Schema for `inheritance_sources`

Optional:

- **gss_tsig_enabled** (Block List, Max: 1) Optional. Field config for _gss_tsig_enabled_ field from _AuthZone_ object. (see [below for nested schema](#nestedblock--inheritance_sources--gss_tsig_enabled))
- **notify** (Block List, Max: 1) Field config for _notify_ field from _AuthZone_ object. (see [below for nested schema](#nestedblock--inheritance_sources--notify))
- **query_acl** (Block List, Max: 1) Optional. Field config for _query_acl_ field from _AuthZone_ object. (see [below for nested schema](#nestedblock--inheritance_sources--query_acl))
- **transfer_acl** (Block List, Max: 1) Optional. Field config for _transfer_acl_ field from _AuthZone_ object. (see [below for nested schema](#nestedblock--inheritance_sources--transfer_acl))
- **update_acl** (Block List, Max: 1) Optional. Field config for _update_acl_ field from _AuthZone_ object. (see [below for nested schema](#nestedblock--inheritance_sources--update_acl))
- **use_forwarders_for_subzones** (Block List, Max: 1) Optional. Field config for _use_forwarders_for_subzones_ field from _AuthZone_ object. (see [below for nested schema](#nestedblock--inheritance_sources--use_forwarders_for_subzones))
- **zone_authority** (Block List, Max: 1) Optional. Field config for _zone_authority_ field from _AuthZone_ object. (see [below for nested schema](#nestedblock--inheritance_sources--zone_authority))

<a id="nestedblock--inheritance_sources--gss_tsig_enabled"></a>
### Nested Schema for `inheritance_sources.gss_tsig_enabled`

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


<a id="nestedblock--inheritance_sources--notify"></a>
### Nested Schema for `inheritance_sources.notify`

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


<a id="nestedblock--inheritance_sources--query_acl"></a>
### Nested Schema for `inheritance_sources.query_acl`

Optional:

- **action** (String) Optional. Inheritance setting for a field.
  Defaults to _inherit_.
- **source** (String) The resource identifier.

Read-Only:

- **display_name** (String) Human-readable display name for the object referred to by _source_.
- **value** (List of Object) Inherited value. (see [below for nested schema](#nestedatt--inheritance_sources--query_acl--value))

<a id="nestedatt--inheritance_sources--query_acl--value"></a>
### Nested Schema for `inheritance_sources.query_acl.value`

Read-Only:

- **access** (String) Access permission for _element_.
  
  Allowed values:
  * _allow_,
  * _deny_.
- **acl** (String) The resource identifier.
- **address** (String) Optional. Data for _ip_ _element_.
  Must be empty if _element_ is not _ip_.
- **element** (String) Type of element.
                       
  Allowed values:
  * _any_,
  * _ip_,
  * _acl_,
  * _tsig_key_.
- **tsig_key** (List of Object) Optional. TSIG key.
  Must be empty if _element_ is not _tsig_key_. (see [below for nested schema](#nestedobjatt--inheritance_sources--query_acl--value--tsig_key))

<a id="nestedobjatt--inheritance_sources--query_acl--value--tsig_key"></a>
### Nested Schema for `inheritance_sources.query_acl.value.tsig_key`

Read-Only:

- **algorithm** (String) TSIG key algorithm.

  Possible values:
  * _hmac_sha256_,
  * _hmac_sha1_,
  * _hmac_sha224_,
  * _hmac_sha384_,
  * _hmac_sha512_.

- **comment** (String) Comment for TSIG key.
- **key** (String) The resource identifier.
- **name** (String) TSIG key name, FQDN.
- **protocol_name** (String) TSIG key name in punycode.
- **secret** (String) TSIG key secret, base64 string.


<a id="nestedblock--inheritance_sources--transfer_acl"></a>
### Nested Schema for `inheritance_sources.transfer_acl`

Optional:

- **action** (String) Optional. Inheritance setting for a field.
  Defaults to _inherit_.
- **source** (String) The resource identifier.

Read-Only:

- **display_name** (String) Human-readable display name for the object referred to by _source_.
- **value** (List of Object) Inherited value. (see [below for nested schema](#nestedatt--inheritance_sources--transfer_acl--value))

<a id="nestedatt--inheritance_sources--transfer_acl--value"></a>
### Nested Schema for `inheritance_sources.transfer_acl.value`

Read-Only:

- **access** (String) Access permission for _element_.
  Allowed values:
  * _allow_,
  * _deny_.
- **acl** (String) The resource identifier.
- **address** (String) Optional. Data for _ip_ _element_.
  Must be empty if _element_ is not _ip_.
- **element** (String) Type of element.
  Allowed values:
  * _any_,
  * _ip_,
  * _acl_,
  * _tsig_key_.
- **tsig_key** (List of Object) Optional. TSIG key.
  Must be empty if _element_ is not _tsig_key_. (see [below for nested schema](#nestedobjatt--inheritance_sources--transfer_acl--value--tsig_key))

<a id="nestedobjatt--inheritance_sources--transfer_acl--value--tsig_key"></a>
### Nested Schema for `inheritance_sources.transfer_acl.value.tsig_key`

Read-Only:

- **algorithm** (String) TSIG key algorithm.
  Possible values:
  * _hmac_sha256_,
  * _hmac_sha1_,
  * _hmac_sha224_,
  * _hmac_sha384_,
  * _hmac_sha512_.
- **comment** (String) Comment for TSIG key.
- **key** (String) The resource identifier.
- **name** (String) TSIG key name, FQDN.
- **protocol_name** (String) TSIG key name in punycode.
- **secret** (String) TSIG key secret, base64 string.

<a id="nestedblock--inheritance_sources--update_acl"></a>
### Nested Schema for `inheritance_sources.update_acl`

Optional:

- **action** (String) Optional. Inheritance setting for a field.
  Defaults to _inherit_.
- **source** (String) The resource identifier.

Read-Only:

- **display_name** (String) Human-readable display name for the object referred to by _source_.
- **value** (List of Object) Inherited value. (see [below for nested schema](#nestedatt--inheritance_sources--update_acl--value))

<a id="nestedatt--inheritance_sources--update_acl--value"></a>
### Nested Schema for `inheritance_sources.update_acl.value`

Read-Only:

- **access** (String) Access permission for _element_.
  Allowed values:
  * _allow_,
  * _deny_.
- **acl** (String) The resource identifier.
- **address** (String) Optional. Data for _ip_ _element_.
  Must be empty if _element_ is not _ip_.
- **element** (String) Type of element.
  Allowed values:
  * _any_,
  * _ip_,
  * _acl_,
  * _tsig_key_.
- **tsig_key** (List of Object) Optional. TSIG key.
  Must be empty if _element_ is not _tsig_key_. (see [below for nested schema](#nestedobjatt--inheritance_sources--update_acl--value--tsig_key))

<a id="nestedobjatt--inheritance_sources--update_acl--value--tsig_key"></a>
### Nested Schema for `inheritance_sources.update_acl.value.tsig_key`

Read-Only:

- **algorithm** (String) TSIG key algorithm.
  Possible values:
  * _hmac_sha256_,
  * _hmac_sha1_,
  * _hmac_sha224_,
  * _hmac_sha384_,
  * _hmac_sha512_.
- **comment** (String) Comment for TSIG key.
- **key** (String) The resource identifier.
- **name** (String) TSIG key name, FQDN.
- **protocol_name** (String) TSIG key name in punycode.
- **secret** (String) TSIG key secret, base64 string.

<a id="nestedblock--inheritance_sources--use_forwarders_for_subzones"></a>
### Nested Schema for `inheritance_sources.use_forwarders_for_subzones`

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


<a id="nestedblock--inheritance_sources--zone_authority"></a>
### Nested Schema for `inheritance_sources.zone_authority`

Optional:

- **default_ttl** (Block List, Max: 1) Optional. Field config for _default_ttl_ field from _ZoneAuthority_ object. (see [below for nested schema](#nestedblock--inheritance_sources--zone_authority--default_ttl))
- **expire** (Block List, Max: 1) Optional. Field config for _expire_ field from _ZoneAuthority_ object. (see [below for nested schema](#nestedblock--inheritance_sources--zone_authority--expire))
- **mname_block** (Block List, Max: 1) Optional. Field config for _mname_ block from _ZoneAuthority_ object. (see [below for nested schema](#nestedblock--inheritance_sources--zone_authority--mname_block))
- **negative_ttl** (Block List, Max: 1) Optional. Field config for _negative_ttl_ field from _ZoneAuthority_ object. (see [below for nested schema](#nestedblock--inheritance_sources--zone_authority--negative_ttl))
- **protocol_rname** (Block List, Max: 1) Optional. Field config for _protocol_rname_ field from _ZoneAuthority_ object. (see [below for nested schema](#nestedblock--inheritance_sources--zone_authority--protocol_rname))
- **refresh** (Block List, Max: 1) Optional. Field config for _refresh_ field from _ZoneAuthority_ object. (see [below for nested schema](#nestedblock--inheritance_sources--zone_authority--refresh))
- **retry** (Block List, Max: 1) Optional. Field config for _retry_ field from _ZoneAuthority_ object. (see [below for nested schema](#nestedblock--inheritance_sources--zone_authority--retry))
- **rname** (Block List, Max: 1) Optional. Field config for _rname_ field from _ZoneAuthority_ object. (see [below for nested schema](#nestedblock--inheritance_sources--zone_authority--rname))

<a id="nestedblock--inheritance_sources--zone_authority--default_ttl"></a>
### Nested Schema for `inheritance_sources.zone_authority.default_ttl`

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


<a id="nestedblock--inheritance_sources--zone_authority--expire"></a>
### Nested Schema for `inheritance_sources.zone_authority.expire`

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


<a id="nestedblock--inheritance_sources--zone_authority--mname_block"></a>
### Nested Schema for `inheritance_sources.zone_authority.mname_block`

Optional:

- **action** (String) Defaults to _inherit_.
- **source** (String) The resource identifier.

Read-Only:

- **display_name** (String) Human-readable display name for the object referred to by _source_.
- **value** (List of Object) Inherited value. (see [below for nested schema](#nestedatt--inheritance_sources--zone_authority--mname_block--value))

<a id="nestedatt--inheritance_sources--zone_authority--mname_block--value"></a>
### Nested Schema for `inheritance_sources.zone_authority.mname_block.value`

Read-Only:

- **mname** (String) Defaults to empty.
- **protocol_mname** (String) Optional. Master name server in punycode.
  Defaults to empty.
- **use_default_mname** (Boolean) Optional. Use default value for master name server.
  Defaults to true.

<a id="nestedblock--inheritance_sources--zone_authority--negative_ttl"></a>
### Nested Schema for `inheritance_sources.zone_authority.negative_ttl`

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


<a id="nestedblock--inheritance_sources--zone_authority--protocol_rname"></a>
### Nested Schema for `inheritance_sources.zone_authority.protocol_rname`

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


<a id="nestedblock--inheritance_sources--zone_authority--refresh"></a>
### Nested Schema for `inheritance_sources.zone_authority.refresh`

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


<a id="nestedblock--inheritance_sources--zone_authority--retry"></a>
### Nested Schema for `inheritance_sources.zone_authority.retry`

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


<a id="nestedblock--inheritance_sources--zone_authority--rname"></a>
### Nested Schema for `inheritance_sources.zone_authority.rname`

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




<a id="nestedblock--internal_secondaries"></a>
### Nested Schema for `internal_secondaries`

Required:

- **host** (String) The resource identifier.


<a id="nestedblock--query_acl"></a>
### Nested Schema for `query_acl`

Required:

- **access** (String) Access permission for _element_.

  Allowed values:
   * _allow_,
   * _deny_.
- **element** (String) Type of element.

  Allowed values:
   * _any_,
   * _ip_,
   * _acl_,
   * _tsig_key_.

Optional:

- **acl** (String) The resource identifier.
- **address** (String) Optional. Data for _ip_ _element_.

  Must be empty if _element_ is not _ip_.
- **tsig_key** (Block List, Max: 1) Optional. TSIG key.

  Must be empty if _element_ is not _tsig_key_. (see [below for nested schema](#nestedblock--query_acl--tsig_key))

<a id="nestedblock--query_acl--tsig_key"></a>
### Nested Schema for `query_acl.tsig_key`

Required:

- **key** (String) The resource identifier.

Optional:

- **algorithm** (String) TSIG key algorithm.

  Possible values:
   * _hmac_sha256_,
   * _hmac_sha1_,
   * _hmac_sha224_,
   * _hmac_sha384_,
   * _hmac_sha512_.
- **comment** (String) Comment for TSIG key.
- **name** (String) TSIG key name, FQDN.
- **secret** (String) TSIG key secret, base64 string.

Read-Only:

- **protocol_name** (String) TSIG key name in punycode.

<a id="nestedblock--transfer_acl"></a>
### Nested Schema for `transfer_acl`

Required:

- **access** (String) Access permission for _element_.

  Allowed values:
   * _allow_,
   * _deny_.
- **element** (String) Type of element.

  Allowed values:
   * _any_,
   * _ip_,
   * _acl_,
   * _tsig_key_.

Optional:

- **acl** (String) The resource identifier.
- **address** (String) Optional. Data for _ip_ _element_.

  Must be empty if _element_ is not _ip_.
- **tsig_key** (Block List, Max: 1) Optional. TSIG key.

  Must be empty if _element_ is not _tsig_key_. (see [below for nested schema](#nestedblock--transfer_acl--tsig_key))

<a id="nestedblock--transfer_acl--tsig_key"></a>
### Nested Schema for `transfer_acl.tsig_key`

Required:

- **key** (String) The resource identifier.

Optional:

- **algorithm** (String) TSIG key algorithm.

  Possible values:
   * _hmac_sha256_,
   * _hmac_sha1_,
   * _hmac_sha224_,
   * _hmac_sha384_,
   * _hmac_sha512_.
- **comment** (String) Comment for TSIG key.
- **name** (String) TSIG key name, FQDN.
- **secret** (String) TSIG key secret, base64 string.

Read-Only:

- **protocol_name** (String) TSIG key name in punycode.

<a id="nestedblock--update_acl"></a>
### Nested Schema for `update_acl`

Required:

- **access** (String) Access permission for _element_.

  Allowed values:
   * _allow_,
   * _deny_.
- **element** (String) Type of element.

  Allowed values:
   * _any_,
   * _ip_,
   * _acl_,
   * _tsig_key_.

Optional:

- **acl** (String) The resource identifier.
- **address** (String) Optional. Data for _ip_ _element_.

  Must be empty if _element_ is not _ip_.
- **tsig_key** (Block List, Max: 1) Optional. TSIG key.

  Must be empty if _element_ is not _tsig_key_. (see [below for nested schema](#nestedblock--update_acl--tsig_key))

<a id="nestedblock--update_acl--tsig_key"></a>
### Nested Schema for `update_acl.tsig_key`

Required:

- **key** (String) The resource identifier.

Optional:

- **algorithm** (String) TSIG key algorithm.

  Possible values:
   * _hmac_sha256_,
   * _hmac_sha1_,
   * _hmac_sha224_,
   * _hmac_sha384_,
   * _hmac_sha512_.
- **comment** (String) Comment for TSIG key.
- **name** (String) TSIG key name, FQDN.
- **secret** (String) TSIG key secret, base64 string.

Read-Only:

- **protocol_name** (String) TSIG key name in punycode.

<a id="nestedblock--zone_authority"></a>
### Nested Schema for `zone_authority`

Optional:

- **default_ttl** (Number) Optional. ZoneAuthority default ttl for resource records in zone (value in seconds).

  Defaults to 28800.
- **expire** (Number) Optional. ZoneAuthority expire time in seconds.
  Defaults to 2419200.
- **mname** (String) Defaults to empty.
- **negative_ttl** (Number) Optional. ZoneAuthority negative caching (minimum) ttl in seconds.
  Defaults to 900.
- **refresh** (Number) Optional. ZoneAuthority refresh.
  Defaults to 10800.
- **retry** (Number) Optional. ZoneAuthority retry.
  Defaults to 3600.
- **rname** (String) Optional. ZoneAuthority rname.
  Defaults to empty.
- **use_default_mname** (Boolean) Optional. Use default value for master name server.
  Defaults to true.

Read-Only:

- **protocol_mname** (String) Optional. ZoneAuthority master name server in punycode.
  Defaults to empty.
- **protocol_rname** (String) Optional. A domain name which specifies the mailbox of the person responsible for this zone.
  Defaults to empty.


<a id="nestedatt--inheritance_assigned_hosts"></a>
### Nested Schema for `inheritance_assigned_hosts`

Read-Only:

- **display_name** (String) The human-readable display name for the host referred to by _ophid_.
- **host** (String) The resource identifier.
- **ophid** (String) The on-prem host ID.

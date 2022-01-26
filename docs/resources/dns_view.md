# b1ddi_dns_view (Resource)

## Example Usage

```terraform
terraform {
  required_providers {
    b1ddi = {
      version = "0.1"
      source  = "github.com/infobloxopen/b1ddi"
    }
  }
}

resource "b1ddi_dns_view" "tf_example_dns_view" {
  name = "example_tf_dns_view"
}
```

## Schema

### Required

- **name** (String) Name of view.

### Optional

- **comment** (String) Optional. Comment for view.
- **custom_root_ns** (Block List) Optional. List of custom root nameservers. The order does not matter.
  Error if empty while _custom_root_ns_enabled_ is _true_.
  Error if there are duplicate items in the list.
  Defaults to empty. (see [below for nested schema](#nestedblock--custom_root_ns))
- **custom_root_ns_enabled** (Boolean) Optional. _true_ to use custom root nameservers instead of the default ones.
  The _custom_root_ns_ is validated when enabled.
  Defaults to _false_.
- **disabled** (Boolean) Optional. _true_ to disable object. A disabled object is effectively non-existent when generating configuration.
- **dnssec_enable_validation** (Boolean) Optional. _true_ to perform DNSSEC validation.
  Ignored if _dnssec_enabled_ is _false_.
  Defaults to _true_.
- **dnssec_enabled** (Boolean) Optional. Master toggle for all DNSSEC processing.
  Other _dnssec*_ configuration is unused if this is disabled.

  Defaults to _true_.
- **dnssec_trust_anchors** (Block List) Optional. DNSSEC trust anchors.

  Error if there are list items with duplicate (_zone_, _sep_, _algorithm_) combinations.

  Defaults to empty. (see [below for nested schema](#nestedblock--dnssec_trust_anchors))
- **dnssec_validate_expiry** (Boolean) Optional. _true_ to reject expired DNSSEC keys.
  Ignored if either _dnssec_enabled_ or _dnssec_enable_validation_ is _false_.

  Defaults to _true_.
- **ecs_enabled** (Boolean) Optional. _true_ to enable EDNS client subnet for recursive queries.
  Other _ecs*_ fields are ignored if this field is not enabled.
  Defaults to _false_.
- **ecs_forwarding** (Boolean) Optional. _true_ to enable ECS options in outbound queries. This functionality has additional overhead so it is disabled by default.
  Defaults to _false_.
- **ecs_prefix_v4** (Number) Optional. Maximum scope length for v4 ECS.
  Unsigned integer, min 1 max 24
  Defaults to 24.
- **ecs_prefix_v6** (Number) Optional. Maximum scope length for v6 ECS.
  Unsigned integer, min 1 max 56
  Defaults to 56.
- **ecs_zones** (Block List) Optional. List of zones where ECS queries may be sent.
  Error if empty while _ecs_enabled_ is _true_.
  Error if there are duplicate FQDNs in the list.
  Defaults to empty. (see [below for nested schema](#nestedblock--ecs_zones))
- **edns_udp_size** (Number) Optional. _edns_udp_size_ represents the edns UDP size.
  The size a querying DNS server advertises to the DNS server it’s sending a query to.
  Defaults to 1232 bytes.
- **forwarders** (Block List) Optional. List of forwarders.
  Error if empty while _forwarders_only_ is _true_.
  Error if there are items in the list with duplicate addresses.
  Defaults to empty. (see [below for nested schema](#nestedblock--forwarders))
- **forwarders_only** (Boolean) Optional. _true_ to only forward.
  Defaults to _false_.
- **gss_tsig_enabled** (Boolean) _gss_tsig_enabled_ enables/disables GSS-TSIG signed dynamic updates.
  Defaults to _false_.
- **id** (String) The ID of this resource.
- **inheritance_sources** (Block List, Max: 1) Optional. Inheritance configuration. (see [below for nested schema](#nestedblock--inheritance_sources))
- **ip_spaces** (List of String) The resource identifier.
- **lame_ttl** (Number) Optional. Unused in the current on-prem DNS server implementation.
  Unsigned integer, min 0 max 3600 (1h).
  Defaults to 600.
- **match_clients_acl** (Block List) Optional. Specifies which clients have access to the view.
  Defaults to empty. (see [below for nested schema](#nestedblock--match_clients_acl))
- **match_destinations_acl** (Block List) Optional. Specifies which destination addresses have access to the view.
  Defaults to empty. (see [below for nested schema](#nestedblock--match_destinations_acl))
- **match_recursive_only** (Boolean) Optional. If _true_ only recursive queries from matching clients access the view.
  Defaults to _false_.
- **max_cache_ttl** (Number) Optional. Seconds to cache positive responses.
  Unsigned integer, min 1 max 604800 (7d).
  Defaults to 604800 (7d).
- **max_negative_ttl** (Number) Optional. Seconds to cache negative responses.

  Unsigned integer, min 1 max 604800 (7d).

  Defaults to 10800 (3h).
- **max_udp_size** (Number) Optional. _max_udp_size_ represents maximum UDP payload size.
  The maximum number of bytes a responding DNS server will send to a UDP datagram.
  Defaults to 1232 bytes.
- **minimal_responses** (Boolean) Optional. When enabled, the DNS server will only add records to the authority and additional data sections when they are required.
  Defaults to _false_.
- **notify** (Boolean) _notify_ all external secondary DNS servers.
  Defaults to _false_.
- **query_acl** (Block List) Optional. Clients must match this ACL to make authoritative queries.
  Also used for recursive queries if that ACL is unset.
  Defaults to empty. (see [below for nested schema](#nestedblock--query_acl))
- **recursion_acl** (Block List) Optional. Clients must match this ACL to make recursive queries. If this ACL is empty, then the _query_acl_ will be used instead.
  Defaults to empty. (see [below for nested schema](#nestedblock--recursion_acl))
- **recursion_enabled** (Boolean) Optional. _true_ to allow recursive DNS queries.
  Defaults to _true_.
- **tags** (Map of String) Tagging specifics.
- **transfer_acl** (Block List) Optional. Clients must match this ACL to receive zone transfers.
  Defaults to empty. (see [below for nested schema](#nestedblock--transfer_acl))
- **update_acl** (Block List) Optional. Specifies which hosts are allowed to issue Dynamic DNS updates for authoritative zones of _primary_type_ _cloud_.
  Defaults to empty. (see [below for nested schema](#nestedblock--update_acl))
- **use_forwarders_for_subzones** (Boolean) Optional. Use default forwarders to resolve queries for subzones.
  Defaults to _true_.
- **zone_authority** (Block List, Max: 1) Optional. ZoneAuthority. (see [below for nested schema](#nestedblock--zone_authority))

### Read-Only

- **created_at** (String) The timestamp when the object has been created.
- **dnssec_root_keys** (List of Object) DNSSEC root keys. The root keys are not configurable.
  A default list is provided by cloud management and included here for config generation. (see [below for nested schema](#nestedatt--dnssec_root_keys))
- **updated_at** (String) The timestamp when the object has been updated. Equals to _created_at_ if not updated after creation.

<a id="nestedblock--custom_root_ns"></a>
### Nested Schema for `custom_root_ns`

Required:

- **address** (String) IPv4 address.
- **fqdn** (String) FQDN.

Read-Only:

- **protocol_fqdn** (String) FQDN in punycode.


<a id="nestedblock--dnssec_trust_anchors"></a>
### Nested Schema for `dnssec_trust_anchors`

Required:

- **algorithm** (Number) Key algorithm. Algogorithm values are as per standards.
  
  The mapping is as follows:
  * _RSAMD5_ = 1,
  * _DH_ = 2,
  * _DSA_ = 3,
  * _RSASHA1_ = 5,
  * _DSANSEC3SHA1_ = 6,
  * _RSASHA1NSEC3SHA1_ = 7,
  * _RSASHA256_ = 8,
  * _RSASHA512_ = 10,
  * _ECDSAP256SHA256_ = 13,
  * _ECDSAP384SHA384_ = 14.
  
  Below algorithms are deprecated and not supported anymore
  * _RSAMD5_ = 1,
  * _DSA_ = 3,
  * _DSANSEC3SHA1_ = 6,
- **public_key** (String) DNSSEC key data. Non-empty, valid base64 string.
- **zone** (String) Zone FQDN.

Optional:

- **sep** (Boolean) Optional. Secure Entry Point flag.

  Defaults to _true_.

Read-Only:

- **protocol_zone** (String) Zone FQDN in punycode.


<a id="nestedblock--ecs_zones"></a>
### Nested Schema for `ecs_zones`

Required:

- **access** (String) Access control for zone.

  Allowed values:
  * _allow_,
  * _deny_.
- **fqdn** (String) Zone FQDN.

Read-Only:

- **protocol_fqdn** (String) Zone FQDN in punycode.


<a id="nestedblock--forwarders"></a>
### Nested Schema for `forwarders`

Required:

- **address** (String) Server IP address.
- **fqdn** (String) Server FQDN.

Read-Only:

- **protocol_fqdn** (String) Server FQDN in punycode.


<a id="nestedblock--inheritance_sources"></a>
### Nested Schema for `inheritance_sources`

Optional:

- **custom_root_ns_block** (Block List, Max: 1) Optional. Field config for _custom_root_ns_block_ field from _View_ object. (see [below for nested schema](#nestedblock--inheritance_sources--custom_root_ns_block))
- **dnssec_validation_block** (Block List, Max: 1) Optional. Field config for _dnssec_validation_block_ field from _View_ object. (see [below for nested schema](#nestedblock--inheritance_sources--dnssec_validation_block))
- **ecs_block** (Block List, Max: 1) Optional. Field config for _ecs_block_ field from _View_ object. (see [below for nested schema](#nestedblock--inheritance_sources--ecs_block))
- **edns_udp_size** (Block List, Max: 1) Optional. Field config for _edns_udp_size_ field from [View] object. (see [below for nested schema](#nestedblock--inheritance_sources--edns_udp_size))
- **forwarders_block** (Block List, Max: 1) Optional. Field config for _forwarders_block_ field from _View_ object. (see [below for nested schema](#nestedblock--inheritance_sources--forwarders_block))
- **gss_tsig_enabled** (Block List, Max: 1) Optional. Field config for _gss_tsig_enabled_ field from _View_ object. (see [below for nested schema](#nestedblock--inheritance_sources--gss_tsig_enabled))
- **lame_ttl** (Block List, Max: 1) Optional. Field config for _lame_ttl_ field from _View_ object. (see [below for nested schema](#nestedblock--inheritance_sources--lame_ttl))
- **match_recursive_only** (Block List, Max: 1) Optional. Field config for _match_recursive_only_ field from _View_ object. (see [below for nested schema](#nestedblock--inheritance_sources--match_recursive_only))
- **max_cache_ttl** (Block List, Max: 1) Optional. Field config for _max_cache_ttl_ field from _View_ object. (see [below for nested schema](#nestedblock--inheritance_sources--max_cache_ttl))
- **max_negative_ttl** (Block List, Max: 1) Optional. Field config for _max_negative_ttl_ field from _View_ object. (see [below for nested schema](#nestedblock--inheritance_sources--max_negative_ttl))
- **max_udp_size** (Block List, Max: 1) Optional. Field config for _max_udp_size_ field from [View] object. (see [below for nested schema](#nestedblock--inheritance_sources--max_udp_size))
- **minimal_responses** (Block List, Max: 1) Optional. Field config for _minimal_responses_ field from _View_ object. (see [below for nested schema](#nestedblock--inheritance_sources--minimal_responses))
- **notify** (Block List, Max: 1) Field config for _notify_ field from _View_ object. (see [below for nested schema](#nestedblock--inheritance_sources--notify))
- **query_acl** (Block List, Max: 1) Optional. Field config for _query_acl_ field from _View_ object. (see [below for nested schema](#nestedblock--inheritance_sources--query_acl))
- **recursion_acl** (Block List, Max: 1) Optional. Field config for _recursion_acl_ field from _View_ object. (see [below for nested schema](#nestedblock--inheritance_sources--recursion_acl))
- **recursion_enabled** (Block List, Max: 1) Optional. Field config for _recursion_enabled_ field from _View_ object. (see [below for nested schema](#nestedblock--inheritance_sources--recursion_enabled))
- **transfer_acl** (Block List, Max: 1) Optional. Field config for _transfer_acl_ field from _View_ object. (see [below for nested schema](#nestedblock--inheritance_sources--transfer_acl))
- **update_acl** (Block List, Max: 1) Optional. Field config for _update_acl_ field from _View_ object. (see [below for nested schema](#nestedblock--inheritance_sources--update_acl))
- **use_forwarders_for_subzones** (Block List, Max: 1) Optional. Field config for _use_forwarders_for_subzones_ field from _View_ object. (see [below for nested schema](#nestedblock--inheritance_sources--use_forwarders_for_subzones))
- **zone_authority** (Block List, Max: 1) Optional. Field config for _zone_authority_ field from _View_ object. (see [below for nested schema](#nestedblock--inheritance_sources--zone_authority))

<a id="nestedblock--inheritance_sources--custom_root_ns_block"></a>
### Nested Schema for `inheritance_sources.custom_root_ns_block`

Optional:

- **action** (String) Defaults to _inherit_.
- **source** (String) The resource identifier.

Read-Only:

- **display_name** (String) Human-readable display name for the object referred to by _source_.
- **value** (List of Object) Inherited value. (see [below for nested schema](#nestedatt--inheritance_sources--custom_root_ns_block--value))

<a id="nestedatt--inheritance_sources--custom_root_ns_block--value"></a>
### Nested Schema for `inheritance_sources.custom_root_ns_block.value`

Read-Only:

- **custom_root_ns** (List of Object) Optional. Field config for _custom_root_ns_ field. (see [below for nested schema](#nestedobjatt--inheritance_sources--custom_root_ns_block--value--custom_root_ns))
- **custom_root_ns_enabled** (Boolean) Optional. Field config for _custom_root_ns_enabled_ field.

<a id="nestedobjatt--inheritance_sources--custom_root_ns_block--value--custom_root_ns"></a>
### Nested Schema for `inheritance_sources.custom_root_ns_block.value.custom_root_ns`

Read-Only:

- **address** (String) IPv4 address.
- **fqdn** (String) FQDN.
- **protocol_fqdn** (String) FQDN in punycode.


<a id="nestedblock--inheritance_sources--dnssec_validation_block"></a>
### Nested Schema for `inheritance_sources.dnssec_validation_block`

Optional:

- **action** (String) Defaults to _inherit_.
- **source** (String) The resource identifier.

Read-Only:

- **display_name** (String) Human-readable display name for the object referred to by _source_.
- **value** (List of Object) Inherited value. (see [below for nested schema](#nestedatt--inheritance_sources--dnssec_validation_block--value))

<a id="nestedatt--inheritance_sources--dnssec_validation_block--value"></a>
### Nested Schema for `inheritance_sources.dnssec_validation_block.value`

Read-Only:

- **dnssec_enable_validation** (Boolean) Optional. Field config for _dnssec_enable_validation_ field.
- **dnssec_enabled** (Boolean) Optional. Field config for _dnssec_enabled_ field.
- **dnssec_trust_anchors** (List of Object) Optional. Field config for _dnssec_trust_anchors_ field. (see [below for nested schema](#nestedobjatt--inheritance_sources--dnssec_validation_block--value--dnssec_trust_anchors))
- **dnssec_validate_expiry** (Boolean) Optional. Field config for _dnssec_validate_expiry_ field.

<a id="nestedobjatt--inheritance_sources--dnssec_validation_block--value--dnssec_trust_anchors"></a>
### Nested Schema for `inheritance_sources.dnssec_validation_block.value.dnssec_trust_anchors`

Read-Only:

- **algorithm** (Number) Key algorithm. Algorithm values are as per standards.
  
  The mapping is as follows:
  * _RSAMD5_ = 1,
  * _DH_ = 2,
  * _DSA_ = 3,
  * _RSASHA1_ = 5,
  * _DSANSEC3SHA1_ = 6,
  * _RSASHA1NSEC3SHA1_ = 7,
  * _RSASHA256_ = 8,
  * _RSASHA512_ = 10,
  * _ECDSAP256SHA256_ = 13,
  * _ECDSAP384SHA384_ = 14.
  
  Below algorithms are deprecated and not supported anymore
  * _RSAMD5_ = 1,
  * _DSA_ = 3,
  * _DSANSEC3SHA1_ = 6,

- **protocol_zone** (String) Zone FQDN in punycode.
- **public_key** (String) DNSSEC key data. Non-empty, valid base64 string.
- **sep** (Boolean) Optional. Secure Entry Point flag.
  Defaults to _true_.
- **zone** (String) Zone FQDN.




<a id="nestedblock--inheritance_sources--ecs_block"></a>
### Nested Schema for `inheritance_sources.ecs_block`

Optional:

- **action** (String) Defaults to _inherit_.
- **source** (String) The resource identifier.

Read-Only:

- **display_name** (String) Human-readable display name for the object referred to by _source_.
- **value** (List of Object) Inherited value. (see [below for nested schema](#nestedatt--inheritance_sources--ecs_block--value))

<a id="nestedatt--inheritance_sources--ecs_block--value"></a>
### Nested Schema for `inheritance_sources.ecs_block.value`

Read-Only:

- **ecs_enabled** (Boolean) Optional. Field config for _ecs_enabled_ field.
- **ecs_forwarding** (Boolean) Optional. Field config for _ecs_forwarding_ field.
- **ecs_prefix_v4** (Number) Optional. Field config for _ecs_prefix_v4_ field.
- **ecs_prefix_v6** (Number) Optional. Field config for _ecs_prefix_v6_ field.
- **ecs_zones** (List of Object) Optional. Field config for _ecs_zones_ field. (see [below for nested schema](#nestedobjatt--inheritance_sources--ecs_block--value--ecs_zones))

<a id="nestedobjatt--inheritance_sources--ecs_block--value--ecs_zones"></a>
### Nested Schema for `inheritance_sources.ecs_block.value.ecs_zones`

Read-Only:

- **access** (String) Access control for zone.
  Allowed values:
  * _allow_,
  * _deny_. 
- **fqdn** (String) Zone FQDN.
- **protocol_fqdn** (String) Zone FQDN in punycode.


<a id="nestedblock--inheritance_sources--edns_udp_size"></a>
### Nested Schema for `inheritance_sources.edns_udp_size`

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


<a id="nestedblock--inheritance_sources--forwarders_block"></a>
### Nested Schema for `inheritance_sources.forwarders_block`

Optional:

- **action** (String) Defaults to _inherit_.
- **source** (String) The resource identifier.

Read-Only:

- **display_name** (String) Human-readable display name for the object referred to by _source_.
- **value** (List of Object) Inherited value. (see [below for nested schema](#nestedatt--inheritance_sources--forwarders_block--value))

<a id="nestedatt--inheritance_sources--forwarders_block--value"></a>
### Nested Schema for `inheritance_sources.forwarders_block.value`

Read-Only:

- **forwarders** (List of Object) (see [below for nested schema](#nestedobjatt--inheritance_sources--forwarders_block--value--forwarders))
- **forwarders_only** (Boolean)

<a id="nestedobjatt--inheritance_sources--forwarders_block--value--forwarders"></a>
### Nested Schema for `inheritance_sources.forwarders_block.value.forwarders_only`

Read-Only:

- **address** (String) Server IP address.
- **fqdn** (String) Server FQDN.
- **protocol_fqdn** (String) Server FQDN in punycode.


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


<a id="nestedblock--inheritance_sources--lame_ttl"></a>
### Nested Schema for `inheritance_sources.lame_ttl`

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


<a id="nestedblock--inheritance_sources--match_recursive_only"></a>
### Nested Schema for `inheritance_sources.match_recursive_only`

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


<a id="nestedblock--inheritance_sources--max_cache_ttl"></a>
### Nested Schema for `inheritance_sources.max_cache_ttl`

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


<a id="nestedblock--inheritance_sources--max_negative_ttl"></a>
### Nested Schema for `inheritance_sources.max_negative_ttl`

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


<a id="nestedblock--inheritance_sources--max_udp_size"></a>
### Nested Schema for `inheritance_sources.max_udp_size`

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


<a id="nestedblock--inheritance_sources--minimal_responses"></a>
### Nested Schema for `inheritance_sources.minimal_responses`

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
- **address** (String) Optional. Data for _ip_ _element_. Must be empty if _element_ is not _ip_.
- **element** (String) Type of element.
  Allowed values:
  * _any_,
  * _ip_,
  * _acl_,
  * _tsig_key_.
- **tsig_key** (List of Object) Optional. TSIG key. Must be empty if _element_ is not _tsig_key_. (see [below for nested schema](#nestedobjatt--inheritance_sources--query_acl--value--tsig_key))

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

<a id="nestedblock--inheritance_sources--recursion_acl"></a>
### Nested Schema for `inheritance_sources.recursion_acl`

Optional:

- **action** (String) Optional. Inheritance setting for a field.
  Defaults to _inherit_.
- **source** (String) The resource identifier.

Read-Only:

- **display_name** (String) Human-readable display name for the object referred to by _source_.
- **value** (List of Object) Inherited value. (see [below for nested schema](#nestedatt--inheritance_sources--recursion_acl--value))

<a id="nestedatt--inheritance_sources--recursion_acl--value"></a>
### Nested Schema for `inheritance_sources.recursion_acl.value`

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

- **tsig_key** (List of Object) Optional. TSIG key. Must be empty if _element_ is not _tsig_key_. (see [below for nested schema](#nestedobjatt--inheritance_sources--recursion_acl--value--tsig_key))

<a id="nestedobjatt--inheritance_sources--recursion_acl--value--tsig_key"></a>
### Nested Schema for `inheritance_sources.recursion_acl.value.tsig_key`

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

<a id="nestedblock--inheritance_sources--recursion_enabled"></a>
### Nested Schema for `inheritance_sources.recursion_enabled`

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

- **tsig_key** (List of Object) Optional. TSIG key. Must be empty if _element_ is not _tsig_key_. (see [below for nested schema](#nestedobjatt--inheritance_sources--transfer_acl--value--tsig_key))

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

- **tsig_key** (List of Object) Optional. TSIG key. Must be empty if _element_ is not _tsig_key_. (see [below for nested schema](#nestedobjatt--inheritance_sources--update_acl--value--tsig_key))

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

<a id="nestedblock--match_clients_acl"></a>
### Nested Schema for `match_clients_acl`

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

  Must be empty if _element_ is not _tsig_key_. (see [below for nested schema](#nestedblock--match_clients_acl--tsig_key))

<a id="nestedblock--match_clients_acl--tsig_key"></a>
### Nested Schema for `match_clients_acl.tsig_key`

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

<a id="nestedblock--match_destinations_acl"></a>
### Nested Schema for `match_destinations_acl`

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
  Must be empty if _element_ is not _tsig_key_. (see [below for nested schema](#nestedblock--match_destinations_acl--tsig_key))

<a id="nestedblock--match_destinations_acl--tsig_key"></a>
### Nested Schema for `match_destinations_acl.tsig_key`

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

<a id="nestedblock--recursion_acl"></a>
### Nested Schema for `recursion_acl`

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
  Must be empty if _element_ is not _tsig_key_. (see [below for nested schema](#nestedblock--recursion_acl--tsig_key))

<a id="nestedblock--recursion_acl--tsig_key"></a>
### Nested Schema for `recursion_acl.tsig_key`

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


<a id="nestedatt--dnssec_root_keys"></a>
### Nested Schema for `dnssec_root_keys`

Read-Only:

- **algorithm** (Number) Key algorithm. Algorithm values are as per standards.
  
  The mapping is as follows:
  * _RSAMD5_ = 1,
  * _DH_ = 2,
  * _DSA_ = 3,
  * _RSASHA1_ = 5,
  * _DSANSEC3SHA1_ = 6,
  * _RSASHA1NSEC3SHA1_ = 7,
  * _RSASHA256_ = 8,
  * _RSASHA512_ = 10,
  * _ECDSAP256SHA256_ = 13,
  * _ECDSAP384SHA384_ = 14.
  
  Below algorithms are deprecated and not supported anymore
  * _RSAMD5_ = 1,
  * _DSA_ = 3,
  * _DSANSEC3SHA1_ = 6,

- **protocol_zone** (String) Zone FQDN in punycode.
- **public_key** (String) DNSSEC key data. Non-empty, valid base64 string.
- **sep** (Boolean) Optional. Secure Entry Point flag.
  Defaults to _true_.
- **zone** (String) Zone FQDN.

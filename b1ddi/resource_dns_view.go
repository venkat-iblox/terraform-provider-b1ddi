package b1ddi

import (
	"context"
	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/dns_config/view"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// ConfigView View
//
// Named collection of DNS View settings.
//
// swagger:model configView
func resourceConfigView() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceConfigViewCreate,
		ReadContext:   resourceConfigViewRead,
		UpdateContext: resourceConfigViewUpdate,
		DeleteContext: resourceConfigViewDelete,
		Schema: map[string]*schema.Schema{

			// Optional. Comment for view.
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Optional. Comment for view.",
			},

			// The timestamp when the object has been created.
			// Read Only: true
			// Format: date-time
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The timestamp when the object has been created.",
			},

			// Optional. List of custom root nameservers. The order does not matter.
			//
			// Error if empty while _custom_root_ns_enabled_ is _true_.
			// Error if there are duplicate items in the list.
			//
			// Defaults to empty.
			"custom_root_ns": {
				Type:        schema.TypeList,
				Elem:        schemaConfigRootNS(),
				Optional:    true,
				Description: "Optional. List of custom root nameservers. The order does not matter.\n\nError if empty while _custom_root_ns_enabled_ is _true_.\nError if there are duplicate items in the list.\n\nDefaults to empty.",
			},

			// Optional. _true_ to use custom root nameservers instead of the default ones.
			//
			// The _custom_root_ns_ is validated when enabled.
			//
			// Defaults to _false_.
			"custom_root_ns_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Optional. _true_ to use custom root nameservers instead of the default ones.\n\nThe _custom_root_ns_ is validated when enabled.\n\nDefaults to _false_.",
			},

			// Optional. _true_ to disable object. A disabled object is effectively non-existent when generating configuration.
			"disabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Optional. _true_ to disable object. A disabled object is effectively non-existent when generating configuration.",
			},

			// Optional. _true_ to perform DNSSEC validation.
			// Ignored if _dnssec_enabled_ is _false_.
			//
			// Defaults to _true_.
			"dnssec_enable_validation": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Optional. _true_ to perform DNSSEC validation.\nIgnored if _dnssec_enabled_ is _false_.\n\nDefaults to _true_.",
			},

			// Optional. Master toggle for all DNSSEC processing.
			// Other _dnssec_*_ configuration is unused if this is disabled.
			//
			// Defaults to _true_.
			"dnssec_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Optional. Master toggle for all DNSSEC processing.\nOther _dnssec_*_ configuration is unused if this is disabled.\n\nDefaults to _true_.",
			},

			// DNSSEC root keys. The root keys are not configurable.
			//
			// A default list is provided by cloud management and included here for config generation.
			// Read Only: true
			"dnssec_root_keys": {
				Type:        schema.TypeList,
				Elem:        schemaConfigTrustAnchor(),
				Computed:    true,
				Description: "DNSSEC root keys. The root keys are not configurable.\n\nA default list is provided by cloud management and included here for config generation.",
			},

			// Optional. DNSSEC trust anchors.
			//
			// Error if there are list items with duplicate (_zone_, _sep_, _algorithm_) combinations.
			//
			// Defaults to empty.
			"dnssec_trust_anchors": {
				Type:        schema.TypeList,
				Elem:        schemaConfigTrustAnchor(),
				Optional:    true,
				Description: "Optional. DNSSEC trust anchors.\n\nError if there are list items with duplicate (_zone_, _sep_, _algorithm_) combinations.\n\nDefaults to empty.",
			},

			// Optional. _true_ to reject expired DNSSEC keys.
			// Ignored if either _dnssec_enabled_ or _dnssec_enable_validation_ is _false_.
			//
			// Defaults to _true_.
			"dnssec_validate_expiry": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Optional. _true_ to reject expired DNSSEC keys.\nIgnored if either _dnssec_enabled_ or _dnssec_enable_validation_ is _false_.\n\nDefaults to _true_.",
			},

			// Optional. _true_ to enable EDNS client subnet for recursive queries.
			// Other _ecs_*_ fields are ignored if this field is not enabled.
			//
			// Defaults to _false-.
			"ecs_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Optional. _true_ to enable EDNS client subnet for recursive queries.\nOther _ecs_*_ fields are ignored if this field is not enabled.\n\nDefaults to _false-.",
			},

			// Optional. _true_ to enable ECS options in outbound queries. This functionality has additional overhead so it is disabled by default.
			//
			// Defaults to _false_.
			"ecs_forwarding": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Optional. _true_ to enable ECS options in outbound queries. This functionality has additional overhead so it is disabled by default.\n\nDefaults to _false_.",
			},

			// Optional. Maximum scope length for v4 ECS.
			//
			// Unsigned integer, min 1 max 24
			//
			// Defaults to 24.
			"ecs_prefix_v4": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Optional. Maximum scope length for v4 ECS.\n\nUnsigned integer, min 1 max 24\n\nDefaults to 24.",
			},

			// Optional. Maximum scope length for v6 ECS.
			//
			// Unsigned integer, min 1 max 56
			//
			// Defaults to 56.
			"ecs_prefix_v6": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Optional. Maximum scope length for v6 ECS.\n\nUnsigned integer, min 1 max 56\n\nDefaults to 56.",
			},

			// Optional. List of zones where ECS queries may be sent.
			//
			// Error if empty while _ecs_enabled_ is _true_.
			// Error if there are duplicate FQDNs in the list.
			//
			// Defaults to empty.
			"ecs_zones": {
				Type:        schema.TypeList,
				Elem:        schemaConfigECSZone(),
				Optional:    true,
				Description: "Optional. List of zones where ECS queries may be sent.\n\nError if empty while _ecs_enabled_ is _true_.\nError if there are duplicate FQDNs in the list.\n\nDefaults to empty.",
			},

			// Optional. _edns_udp_size_ represents the edns UDP size.
			// The size a querying DNS server advertises to the DNS server it’s sending a query to.
			//
			// Defaults to 1232 bytes.
			"edns_udp_size": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Optional. _edns_udp_size_ represents the edns UDP size.\nThe size a querying DNS server advertises to the DNS server it’s sending a query to.\n\nDefaults to 1232 bytes.",
			},

			// Optional. List of forwarders.
			//
			// Error if empty while _forwarders_only_ is _true_.
			// Error if there are items in the list with duplicate addresses.
			//
			// Defaults to empty.
			"forwarders": {
				Type:        schema.TypeList,
				Elem:        schemaConfigForwarder(),
				Optional:    true,
				Description: "Optional. List of forwarders.\n\nError if empty while _forwarders_only_ is _true_.\nError if there are items in the list with duplicate addresses.\n\nDefaults to empty.",
			},

			// Optional. _true_ to only forward.
			//
			// Defaults to _false_.
			"forwarders_only": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Optional. _true_ to only forward.\n\nDefaults to _false_.",
			},

			// _gss_tsig_enabled_ enables/disables GSS-TSIG signed dynamic updates.
			//
			// Defaults to _false_.
			"gss_tsig_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "_gss_tsig_enabled_ enables/disables GSS-TSIG signed dynamic updates.\n\nDefaults to _false_.",
			},

			// Optional. Inheritance configuration.
			"inheritance_sources": {
				Type:        schema.TypeList,
				Elem:        schemaConfigViewInheritance(),
				MaxItems:    1,
				Optional:    true,
				Description: "Optional. Inheritance configuration.",
			},

			// The resource identifier.
			"ip_spaces": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "The resource identifier.",
			},

			// Optional. Unused in the current on-prem DNS server implementation.
			//
			// Unsigned integer, min 0 max 3600 (1h).
			//
			// Defaults to 600.
			"lame_ttl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Optional. Unused in the current on-prem DNS server implementation.\n\nUnsigned integer, min 0 max 3600 (1h).\n\nDefaults to 600.",
			},

			// Optional. Specifies which clients have access to the view.
			//
			// Defaults to empty.
			"match_clients_acl": {
				Type:        schema.TypeList,
				Elem:        schemaConfigACLItem(),
				Optional:    true,
				Description: "Optional. Specifies which clients have access to the view.\n\nDefaults to empty.",
			},

			// Optional. Specifies which destination addresses have access to the view.
			//
			// Defaults to empty.
			"match_destinations_acl": {
				Type:        schema.TypeList,
				Elem:        schemaConfigACLItem(),
				Optional:    true,
				Description: "Optional. Specifies which destination addresses have access to the view.\n\nDefaults to empty.",
			},

			// Optional. If _true_ only recursive queries from matching clients access the view.
			//
			// Defaults to _false_.
			"match_recursive_only": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Optional. If _true_ only recursive queries from matching clients access the view.\n\nDefaults to _false_.",
			},

			// Optional. Seconds to cache positive responses.
			//
			// Unsigned integer, min 1 max 604800 (7d).
			//
			// Defaults to 604800 (7d).
			"max_cache_ttl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Optional. Seconds to cache positive responses.\n\nUnsigned integer, min 1 max 604800 (7d).\n\nDefaults to 604800 (7d).",
			},

			// Optional. Seconds to cache negative responses.
			//
			// Unsigned integer, min 1 max 604800 (7d).
			//
			// Defaults to 10800 (3h).
			"max_negative_ttl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Optional. Seconds to cache negative responses.\n\nUnsigned integer, min 1 max 604800 (7d).\n\nDefaults to 10800 (3h).",
			},

			// Optional. _max_udp_size_ represents maximum UDP payload size.
			// The maximum number of bytes a responding DNS server will send to a UDP datagram.
			//
			// Defaults to 1232 bytes.
			"max_udp_size": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Optional. _max_udp_size_ represents maximum UDP payload size.\nThe maximum number of bytes a responding DNS server will send to a UDP datagram.\n\nDefaults to 1232 bytes.",
			},

			// Optional. When enabled, the DNS server will only add records to the authority and additional data sections when they are required.
			//
			// Defaults to _false_.
			"minimal_responses": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Optional. When enabled, the DNS server will only add records to the authority and additional data sections when they are required.\n\nDefaults to _false_.",
			},

			// Name of view.
			// Required: true
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of view.",
			},

			// _notify_ all external secondary DNS servers.
			//
			// Defaults to _false_.
			"notify": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "_notify_ all external secondary DNS servers.\n\nDefaults to _false_.",
			},

			// Optional. Clients must match this ACL to make authoritative queries.
			// Also used for recursive queries if that ACL is unset.
			//
			// Defaults to empty.
			"query_acl": {
				Type:        schema.TypeList,
				Elem:        schemaConfigACLItem(),
				Optional:    true,
				Description: "Optional. Clients must match this ACL to make authoritative queries.\nAlso used for recursive queries if that ACL is unset.\n\nDefaults to empty.",
			},

			// Optional. Clients must match this ACL to make recursive queries. If this ACL is empty, then the _query_acl_ will be used instead.
			//
			// Defaults to empty.
			"recursion_acl": {
				Type:        schema.TypeList,
				Elem:        schemaConfigACLItem(),
				Optional:    true,
				Description: "Optional. Clients must match this ACL to make recursive queries. If this ACL is empty, then the _query_acl_ will be used instead.\n\nDefaults to empty.",
			},

			// Optional. _true_ to allow recursive DNS queries.
			//
			// Defaults to _true_.
			"recursion_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Optional. _true_ to allow recursive DNS queries.\n\nDefaults to _true_.",
			},

			// Tagging specifics.
			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Tagging specifics.",
			},

			// Optional. Clients must match this ACL to receive zone transfers.
			//
			// Defaults to empty.
			"transfer_acl": {
				Type:        schema.TypeList,
				Elem:        schemaConfigACLItem(),
				Optional:    true,
				Description: "Optional. Clients must match this ACL to receive zone transfers.\n\nDefaults to empty.",
			},

			// Optional. Specifies which hosts are allowed to issue Dynamic DNS updates for authoritative zones of _primary_type_ _cloud_.
			//
			// Defaults to empty.
			"update_acl": {
				Type:        schema.TypeList,
				Elem:        schemaConfigACLItem(),
				Optional:    true,
				Description: "Optional. Specifies which hosts are allowed to issue Dynamic DNS updates for authoritative zones of _primary_type_ _cloud_.\n\nDefaults to empty.",
			},

			// The timestamp when the object has been updated. Equals to _created_at_ if not updated after creation.
			// Read Only: true
			// Format: date-time
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The timestamp when the object has been updated. Equals to _created_at_ if not updated after creation.",
			},

			// Optional. Use default forwarders to resolve queries for subzones.
			//
			// Defaults to _true_.
			"use_forwarders_for_subzones": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Optional. Use default forwarders to resolve queries for subzones.\n\nDefaults to _true_.",
			},

			// Optional. ZoneAuthority.
			"zone_authority": {
				Type:        schema.TypeList,
				Elem:        schemaConfigZoneAuthority(),
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Description: "Optional. ZoneAuthority.",
			},
		},
	}
}

func resourceConfigViewCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	customRootNs := make([]*models.ConfigRootNS, 0)
	for _, rns := range d.Get("custom_root_ns").([]interface{}) {
		if rns != nil {
			customRootNs = append(customRootNs, expandConfigRootNS(rns.(map[string]interface{})))
		}
	}

	dnssecTrustAnchors := make([]*models.ConfigTrustAnchor, 0)
	for _, ta := range d.Get("dnssec_trust_anchors").([]interface{}) {
		if ta != nil {
			dnssecTrustAnchors = append(dnssecTrustAnchors, expandConfigTrustAnchor(ta.(map[string]interface{})))
		}
	}

	ecsZones := make([]*models.ConfigECSZone, 0)
	for _, ecsZone := range d.Get("ecs_zones").([]interface{}) {
		if ecsZone != nil {
			ecsZones = append(ecsZones, expandConfigECSZone(ecsZone.(map[string]interface{})))
		}
	}

	forwarders := make([]*models.ConfigForwarder, 0)
	for _, fwd := range d.Get("forwarders").([]interface{}) {
		if fwd != nil {
			forwarders = append(forwarders, expandConfigForwarder(fwd.(map[string]interface{})))
		}
	}

	ipSpaces := make([]string, 0)
	for _, is := range d.Get("ip_spaces").([]interface{}) {
		if is != nil {
			ipSpaces = append(ipSpaces, is.(string))
		}
	}

	matchClientsACL := make([]*models.ConfigACLItem, 0)
	for _, aclItem := range d.Get("match_clients_acl").([]interface{}) {
		if aclItem != nil {
			matchClientsACL = append(matchClientsACL, expandConfigACLItem(aclItem.(map[string]interface{})))
		}
	}

	matchDestinationsACL := make([]*models.ConfigACLItem, 0)
	for _, aclItem := range d.Get("match_destinations_acl").([]interface{}) {
		if aclItem != nil {
			matchDestinationsACL = append(matchDestinationsACL, expandConfigACLItem(aclItem.(map[string]interface{})))
		}
	}

	queryACL := make([]*models.ConfigACLItem, 0)
	for _, ai := range d.Get("query_acl").([]interface{}) {
		if ai != nil {
			queryACL = append(queryACL, expandConfigACLItem(ai.(map[string]interface{})))
		}
	}

	recursionACL := make([]*models.ConfigACLItem, 0)
	for _, ai := range d.Get("recursion_acl").([]interface{}) {
		if ai != nil {
			recursionACL = append(recursionACL, expandConfigACLItem(ai.(map[string]interface{})))
		}
	}

	transferACL := make([]*models.ConfigACLItem, 0)
	for _, ai := range d.Get("transfer_acl").([]interface{}) {
		if ai != nil {
			transferACL = append(transferACL, expandConfigACLItem(ai.(map[string]interface{})))
		}
	}

	updateACL := make([]*models.ConfigACLItem, 0)
	for _, ai := range d.Get("update_acl").([]interface{}) {
		if ai != nil {
			updateACL = append(updateACL, expandConfigACLItem(ai.(map[string]interface{})))
		}
	}

	v := &models.ConfigView{
		Comment:                  d.Get("comment").(string),
		CustomRootNs:             customRootNs,
		CustomRootNsEnabled:      d.Get("custom_root_ns_enabled").(bool),
		Disabled:                 d.Get("disabled").(bool),
		DnssecEnableValidation:   d.Get("dnssec_enable_validation").(bool),
		DnssecEnabled:            d.Get("dnssec_enabled").(bool),
		DnssecTrustAnchors:       dnssecTrustAnchors,
		DnssecValidateExpiry:     d.Get("dnssec_validate_expiry").(bool),
		EcsEnabled:               d.Get("ecs_enabled").(bool),
		EcsForwarding:            d.Get("ecs_forwarding").(bool),
		EcsPrefixV4:              int64(d.Get("ecs_prefix_v4").(int)),
		EcsPrefixV6:              int64(d.Get("ecs_prefix_v6").(int)),
		EcsZones:                 ecsZones,
		EdnsUDPSize:              int64(d.Get("edns_udp_size").(int)),
		Forwarders:               forwarders,
		ForwardersOnly:           d.Get("forwarders_only").(bool),
		GssTsigEnabled:           d.Get("gss_tsig_enabled").(bool),
		InheritanceSources:       expandConfigViewInheritance(d.Get("inheritance_sources").([]interface{})),
		IPSpaces:                 ipSpaces,
		LameTTL:                  int64(d.Get("lame_ttl").(int)),
		MatchClientsACL:          matchClientsACL,
		MatchDestinationsACL:     matchDestinationsACL,
		MatchRecursiveOnly:       d.Get("match_recursive_only").(bool),
		MaxCacheTTL:              int64(d.Get("max_cache_ttl").(int)),
		MaxNegativeTTL:           int64(d.Get("max_negative_ttl").(int)),
		MaxUDPSize:               int64(d.Get("max_udp_size").(int)),
		MinimalResponses:         d.Get("minimal_responses").(bool),
		Name:                     swag.String(d.Get("name").(string)),
		Notify:                   d.Get("notify").(bool),
		QueryACL:                 queryACL,
		RecursionACL:             recursionACL,
		RecursionEnabled:         d.Get("recursion_enabled").(bool),
		Tags:                     d.Get("tags"),
		TransferACL:              transferACL,
		UpdateACL:                updateACL,
		UseForwardersForSubzones: d.Get("use_forwarders_for_subzones").(bool),
		ZoneAuthority:            expandConfigZoneAuthority(d.Get("zone_authority").([]interface{})),
	}

	resp, err := c.DNSConfigurationAPI.View.ViewCreate(
		&view.ViewCreateParams{Body: v, Context: ctx},
		nil,
	)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.Payload.Result.ID)

	return resourceConfigViewRead(ctx, d, m)
}

func resourceConfigViewRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	var diags diag.Diagnostics

	resp, err := c.DNSConfigurationAPI.View.ViewRead(
		&view.ViewReadParams{ID: d.Id(), Context: ctx},
		nil,
	)
	if err != nil {
		return diag.FromErr(err)
	}

	err = d.Set("comment", resp.Payload.Result.Comment)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("created_at", resp.Payload.Result.CreatedAt.String())
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	customRootNs := make([]interface{}, 0, len(resp.Payload.Result.CustomRootNs))
	for _, ns := range resp.Payload.Result.CustomRootNs {
		customRootNs = append(customRootNs, flattenConfigRootNS(ns))
	}
	err = d.Set("custom_root_ns", customRootNs)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("custom_root_ns_enabled", resp.Payload.Result.CustomRootNsEnabled)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("disabled", resp.Payload.Result.Disabled)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("dnssec_enable_validation", resp.Payload.Result.DnssecEnableValidation)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("dnssec_enabled", resp.Payload.Result.DnssecEnabled)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	dnssecRootKeys := make([]interface{}, 0, len(resp.Payload.Result.DnssecRootKeys))
	for _, drk := range resp.Payload.Result.DnssecRootKeys {
		dnssecRootKeys = append(dnssecRootKeys, flattenConfigTrustAnchor(drk))
	}
	err = d.Set("dnssec_root_keys", dnssecRootKeys)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	dnssecTrustAnchors := make([]interface{}, 0, len(resp.Payload.Result.DnssecTrustAnchors))
	for _, dta := range resp.Payload.Result.DnssecTrustAnchors {
		dnssecTrustAnchors = append(dnssecTrustAnchors, flattenConfigTrustAnchor(dta))
	}
	err = d.Set("dnssec_trust_anchors", dnssecTrustAnchors)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("dnssec_validate_expiry", resp.Payload.Result.DnssecValidateExpiry)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("ecs_enabled", resp.Payload.Result.EcsEnabled)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("ecs_forwarding", resp.Payload.Result.EcsForwarding)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("ecs_prefix_v4", resp.Payload.Result.EcsPrefixV4)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("ecs_prefix_v6", resp.Payload.Result.EcsPrefixV6)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	ecsZones := make([]interface{}, 0, len(resp.Payload.Result.EcsZones))
	for _, ecsZone := range resp.Payload.Result.EcsZones {
		ecsZones = append(ecsZones, flattenConfigECSZone(ecsZone))
	}
	err = d.Set("ecs_zones", ecsZones)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("edns_udp_size", resp.Payload.Result.EdnsUDPSize)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	forwarders := make([]interface{}, 0, len(resp.Payload.Result.Forwarders))
	for _, f := range resp.Payload.Result.Forwarders {
		forwarders = append(forwarders, flattenConfigForwarder(f))
	}
	err = d.Set("forwarders", forwarders)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("forwarders_only", resp.Payload.Result.ForwardersOnly)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("gss_tsig_enabled", resp.Payload.Result.GssTsigEnabled)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("inheritance_sources", flattenConfigViewInheritance(resp.Payload.Result.InheritanceSources))
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("ip_spaces", resp.Payload.Result.IPSpaces)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("lame_ttl", resp.Payload.Result.LameTTL)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	matchClientsACL := make([]interface{}, 0, len(resp.Payload.Result.MatchClientsACL))
	for _, aclItem := range resp.Payload.Result.MatchClientsACL {
		matchClientsACL = append(matchClientsACL, flattenConfigACLItem(aclItem))
	}
	err = d.Set("match_clients_acl", matchClientsACL)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	matchDestinationsACL := make([]interface{}, 0, len(resp.Payload.Result.MatchDestinationsACL))
	for _, aclItem := range resp.Payload.Result.MatchDestinationsACL {
		matchDestinationsACL = append(matchDestinationsACL, flattenConfigACLItem(aclItem))
	}
	err = d.Set("match_destinations_acl", matchDestinationsACL)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("match_recursive_only", resp.Payload.Result.MatchRecursiveOnly)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("max_cache_ttl", resp.Payload.Result.MaxCacheTTL)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("max_negative_ttl", resp.Payload.Result.MaxNegativeTTL)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("max_udp_size", resp.Payload.Result.MaxUDPSize)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("minimal_responses", resp.Payload.Result.MinimalResponses)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("name", resp.Payload.Result.Name)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("notify", resp.Payload.Result.Notify)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	queryACL := make([]interface{}, 0, len(resp.Payload.Result.QueryACL))
	for _, aclItem := range resp.Payload.Result.QueryACL {
		queryACL = append(queryACL, flattenConfigACLItem(aclItem))
	}
	err = d.Set("query_acl", queryACL)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	recursionACL := make([]interface{}, 0, len(resp.Payload.Result.RecursionACL))
	for _, aclItem := range resp.Payload.Result.RecursionACL {
		recursionACL = append(recursionACL, flattenConfigACLItem(aclItem))
	}
	err = d.Set("recursion_acl", recursionACL)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("recursion_enabled", resp.Payload.Result.RecursionEnabled)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("tags", resp.Payload.Result.Tags)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	transferACL := make([]interface{}, 0, len(resp.Payload.Result.TransferACL))
	for _, aclItem := range resp.Payload.Result.TransferACL {
		transferACL = append(transferACL, flattenConfigACLItem(aclItem))
	}
	err = d.Set("transfer_acl", transferACL)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	updateACL := make([]interface{}, 0, len(resp.Payload.Result.UpdateACL))
	for _, aclItem := range resp.Payload.Result.UpdateACL {
		updateACL = append(updateACL, flattenConfigACLItem(aclItem))
	}
	err = d.Set("update_acl", updateACL)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("updated_at", resp.Payload.Result.UpdatedAt.String())
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("use_forwarders_for_subzones", resp.Payload.Result.UseForwardersForSubzones)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("zone_authority", flattenConfigZoneAuthority(resp.Payload.Result.ZoneAuthority))
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	return diags
}

func flattenConfigView(r *models.ConfigView) []interface{} {
	if r == nil {
		return nil
	}

	customRootNs := make([]interface{}, 0, len(r.CustomRootNs))
	for _, ns := range r.CustomRootNs {
		customRootNs = append(customRootNs, flattenConfigRootNS(ns))
	}

	dnssecRootKeys := make([]interface{}, 0, len(r.DnssecRootKeys))
	for _, drk := range r.DnssecRootKeys {
		dnssecRootKeys = append(dnssecRootKeys, flattenConfigTrustAnchor(drk))
	}

	dnssecTrustAnchors := make([]interface{}, 0, len(r.DnssecTrustAnchors))
	for _, dta := range r.DnssecTrustAnchors {
		dnssecTrustAnchors = append(dnssecTrustAnchors, flattenConfigTrustAnchor(dta))
	}

	ecsZones := make([]interface{}, 0, len(r.EcsZones))
	for _, ecsZone := range r.EcsZones {
		ecsZones = append(ecsZones, flattenConfigECSZone(ecsZone))
	}

	forwarders := make([]interface{}, 0, len(r.Forwarders))
	for _, f := range r.Forwarders {
		forwarders = append(forwarders, flattenConfigForwarder(f))
	}

	matchClientsACL := make([]interface{}, 0, len(r.MatchClientsACL))
	for _, aclItem := range r.MatchClientsACL {
		matchClientsACL = append(matchClientsACL, flattenConfigACLItem(aclItem))
	}

	matchDestinationsACL := make([]interface{}, 0, len(r.MatchDestinationsACL))
	for _, aclItem := range r.MatchDestinationsACL {
		matchDestinationsACL = append(matchDestinationsACL, flattenConfigACLItem(aclItem))
	}

	queryACL := make([]interface{}, 0, len(r.QueryACL))
	for _, aclItem := range r.QueryACL {
		queryACL = append(queryACL, flattenConfigACLItem(aclItem))
	}

	recursionACL := make([]interface{}, 0, len(r.RecursionACL))
	for _, aclItem := range r.RecursionACL {
		recursionACL = append(recursionACL, flattenConfigACLItem(aclItem))
	}

	transferACL := make([]interface{}, 0, len(r.TransferACL))
	for _, aclItem := range r.TransferACL {
		transferACL = append(transferACL, flattenConfigACLItem(aclItem))
	}

	updateACL := make([]interface{}, 0, len(r.UpdateACL))
	for _, aclItem := range r.UpdateACL {
		updateACL = append(updateACL, flattenConfigACLItem(aclItem))
	}

	return []interface{}{
		map[string]interface{}{
			"comment":                     r.Comment,
			"created_at":                  r.CreatedAt.String(),
			"custom_root_ns":              customRootNs,
			"custom_root_ns_enabled":      r.CustomRootNsEnabled,
			"disabled":                    r.Disabled,
			"dnssec_enable_validation":    r.DnssecEnableValidation,
			"dnssec_enabled":              r.DnssecEnabled,
			"dnssec_root_keys":            dnssecRootKeys,
			"dnssec_trust_anchors":        dnssecTrustAnchors,
			"dnssec_validate_expiry":      r.DnssecValidateExpiry,
			"ecs_enabled":                 r.EcsEnabled,
			"ecs_forwarding":              r.EcsForwarding,
			"ecs_prefix_v4":               r.EcsPrefixV4,
			"ecs_prefix_v6":               r.EcsPrefixV6,
			"ecs_zones":                   ecsZones,
			"edns_udp_size":               r.EdnsUDPSize,
			"forwarders":                  forwarders,
			"forwarders_only":             r.ForwardersOnly,
			"gss_tsig_enabled":            r.GssTsigEnabled,
			"inheritance_sources":         flattenConfigViewInheritance(r.InheritanceSources),
			"ip_spaces":                   r.IPSpaces,
			"lame_ttl":                    r.LameTTL,
			"match_clients_acl":           matchClientsACL,
			"match_destinations_acl":      matchDestinationsACL,
			"match_recursive_only":        r.MatchRecursiveOnly,
			"max_cache_ttl":               r.MaxCacheTTL,
			"max_negative_ttl":            r.MaxNegativeTTL,
			"max_udp_size":                r.MaxUDPSize,
			"minimal_responses":           r.MinimalResponses,
			"name":                        r.Name,
			"notify":                      r.Notify,
			"query_acl":                   queryACL,
			"recursion_acl":               recursionACL,
			"recursion_enabled":           r.RecursionEnabled,
			"tags":                        r.Tags,
			"transfer_acl":                transferACL,
			"update_acl":                  updateACL,
			"updated_at":                  r.UpdatedAt.String(),
			"use_forwarders_for_subzones": r.UseForwardersForSubzones,
			"zone_authority":              flattenConfigZoneAuthority(r.ZoneAuthority),
		},
	}
}

func resourceConfigViewUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// ToDo Implement resourceConfigViewUpdate function
	return diags
}

func resourceConfigViewDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)
	_, err := c.DNSConfigurationAPI.View.ViewDelete(
		&view.ViewDeleteParams{ID: d.Id(), Context: ctx},
		nil,
	)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return nil
}

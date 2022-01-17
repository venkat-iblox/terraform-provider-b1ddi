package b1ddi

import (
	"context"
	"github.com/go-openapi/swag"
	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/dns_config/auth_zone"
	"github.com/infobloxopen/b1ddi-go-client/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ConfigAuthZone AuthZone
//
// Authoritative zone.
//
// swagger:model configAuthZone
func resourceConfigAuthZone() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceConfigAuthZoneCreate,
		ReadContext:   resourceConfigAuthZoneRead,
		UpdateContext: resourceConfigAuthZoneUpdate,
		DeleteContext: resourceConfigAuthZoneDelete,
		Schema: map[string]*schema.Schema{

			// Optional. Comment for zone configuration.
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Optional. Comment for zone configuration.",
			},

			// Time when the object has been created.
			// Read Only: true
			// Format: date-time
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Time when the object has been created.",
			},

			// Optional. _true_ to disable object. A disabled object is effectively non-existent when generating configuration.
			"disabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Optional. _true_ to disable object. A disabled object is effectively non-existent when generating configuration.",
			},

			// Optional. DNS primaries external to BloxOne DDI. Order is not significant.
			"external_primaries": {
				Type:        schema.TypeList,
				Elem:        schemaConfigExternalPrimary(),
				Optional:    true,
				Description: "Optional. DNS primaries external to BloxOne DDI. Order is not significant.",
			},

			// DNS secondaries external to BloxOne DDI. Order is not significant.
			"external_secondaries": {
				Type:        schema.TypeList,
				Elem:        schemaConfigExternalSecondary(),
				Optional:    true,
				Description: "DNS secondaries external to BloxOne DDI. Order is not significant.",
			},

			// Zone FQDN.
			// The FQDN supplied at creation will be converted to canonical form.
			//
			// Read-only after creation.
			// Required: true
			"fqdn": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Zone FQDN.\nThe FQDN supplied at creation will be converted to canonical form.\n\nRead-only after creation.",
			},

			// _gss_tsig_enabled_ enables/disables GSS-TSIG signed dynamic updates.
			//
			// Defaults to _false_.
			"gss_tsig_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "_gss_tsig_enabled_ enables/disables GSS-TSIG signed dynamic updates.\n\nDefaults to _false_.",
			},

			// The list of the inheritance assigned hosts of the object.
			// Read Only: true
			"inheritance_assigned_hosts": {
				Type:        schema.TypeList,
				Elem:        schemaInheritance2AssignedHost(),
				Computed:    true,
				Description: "The list of the inheritance assigned hosts of the object.",
			},

			// Optional. Inheritance configuration.
			"inheritance_sources": {
				Type:        schema.TypeList,
				Elem:        schemaConfigAuthZoneInheritance(),
				MaxItems:    1,
				Optional:    true,
				Description: "Optional. Inheritance configuration.",
			},

			// On-create-only. SOA serial is allowed to be set when the authoritative zone is created.
			"initial_soa_serial": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "On-create-only. SOA serial is allowed to be set when the authoritative zone is created.",
			},

			// Optional. BloxOne DDI hosts acting as internal secondaries. Order is not significant.
			"internal_secondaries": {
				Type:        schema.TypeList,
				Elem:        schemaConfigInternalSecondary(),
				Optional:    true,
				Description: "Optional. BloxOne DDI hosts acting as internal secondaries. Order is not significant.",
			},

			// Reverse zone network address in the following format: "ip-address/cidr".
			// Defaults to empty.
			// Read Only: true
			"mapped_subnet": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Reverse zone network address in the following format: \"ip-address/cidr\".\nDefaults to empty.",
			},

			// Zone mapping type.
			// Allowed values:
			//  * _forward_,
			//  * _ipv4_reverse_.
			//  * _ipv6_reverse_.
			//
			// Defaults to forward.
			// Read Only: true
			"mapping": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Zone mapping type.\nAllowed values:\n * _forward_,\n * _ipv4_reverse_.\n * _ipv6_reverse_.\n\nDefaults to forward.",
			},

			// Also notify all external secondary DNS servers if enabled.
			//
			// Defaults to _false_.
			"notify": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Also notify all external secondary DNS servers if enabled.\n\nDefaults to _false_.",
			},

			// The resource identifier.
			"nsgs": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "The resource identifier.",
			},

			// The resource identifier.
			"parent": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The resource identifier.",
			},

			// Primary type for an authoritative zone.
			// Read only after creation.
			// Allowed values:
			//  * _external_: zone data owned by an external nameserver,
			//  * _cloud_: zone data is owned by a BloxOne DDI host.
			// Required: true
			"primary_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Primary type for an authoritative zone.\nRead only after creation.\nAllowed values:\n * _external_: zone data owned by an external nameserver,\n * _cloud_: zone data is owned by a BloxOne DDI host.",
			},

			// Zone FQDN in punycode.
			// Read Only: true
			"protocol_fqdn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Zone FQDN in punycode.",
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

			// Tagging specifics.
			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Tagging specifics.",
			},

			// Optional. Clients must match this ACL to receive zone transfers.
			"transfer_acl": {
				Type:        schema.TypeList,
				Elem:        schemaConfigACLItem(),
				Optional:    true,
				Description: "Optional. Clients must match this ACL to receive zone transfers.",
			},

			// Optional. Specifies which hosts are allowed to submit Dynamic DNS updates for authoritative zones of _primary_type_ _cloud_.
			//
			// Defaults to empty.
			"update_acl": {
				Type:        schema.TypeList,
				Elem:        schemaConfigACLItem(),
				Optional:    true,
				Description: "Optional. Specifies which hosts are allowed to submit Dynamic DNS updates for authoritative zones of _primary_type_ _cloud_.\n\nDefaults to empty.",
			},

			// Time when the object has been updated. Equals to _created_at_ if not updated after creation.
			// Read Only: true
			// Format: date-time
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Time when the object has been updated. Equals to _created_at_ if not updated after creation.",
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

			// The resource identifier.
			"view": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The resource identifier.",
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

func resourceConfigAuthZoneCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	externalPrimaries := make([]*models.ConfigExternalPrimary, 0)
	for _, ep := range d.Get("external_primaries").([]interface{}) {
		if ep != nil {
			externalPrimaries = append(externalPrimaries, expandConfigExternalPrimary(ep.(map[string]interface{})))
		}
	}

	externalSecondaries := make([]*models.ConfigExternalSecondary, 0)
	for _, es := range d.Get("external_secondaries").([]interface{}) {
		if es != nil {
			externalSecondaries = append(externalSecondaries, expandConfigExternalSecondary(es.(map[string]interface{})))
		}
	}

	internalSecondaries := make([]*models.ConfigInternalSecondary, 0)
	for _, is := range d.Get("internal_secondaries").([]interface{}) {
		if is != nil {
			internalSecondaries = append(internalSecondaries, expandConfigInternalSecondary(is.(map[string]interface{})))
		}
	}

	nsgs := make([]string, 0)
	for _, n := range d.Get("nsgs").([]interface{}) {
		if n != nil {
			nsgs = append(nsgs, n.(string))
		}
	}

	queryACL := make([]*models.ConfigACLItem, 0)
	for _, aclItem := range d.Get("query_acl").([]interface{}) {
		if aclItem != nil {
			queryACL = append(queryACL, expandConfigACLItem(aclItem.(map[string]interface{})))
		}
	}

	transferACL := make([]*models.ConfigACLItem, 0)
	for _, aclItem := range d.Get("transfer_acl").([]interface{}) {
		if aclItem != nil {
			transferACL = append(transferACL, expandConfigACLItem(aclItem.(map[string]interface{})))
		}
	}

	updateACL := make([]*models.ConfigACLItem, 0)
	for _, aclItem := range d.Get("update_acl").([]interface{}) {
		if aclItem != nil {
			updateACL = append(updateACL, expandConfigACLItem(aclItem.(map[string]interface{})))
		}
	}

	az := &models.ConfigAuthZone{
		Comment:                  d.Get("comment").(string),
		Disabled:                 d.Get("disabled").(bool),
		ExternalPrimaries:        externalPrimaries,
		ExternalSecondaries:      externalSecondaries,
		Fqdn:                     swag.String(d.Get("fqdn").(string)),
		GssTsigEnabled:           d.Get("gss_tsig_enabled").(bool),
		InheritanceSources:       expandConfigAuthZoneInheritance(d.Get("inheritance_sources").([]interface{})),
		InitialSoaSerial:         int64(d.Get("initial_soa_serial").(int)),
		InternalSecondaries:      internalSecondaries,
		Notify:                   d.Get("notify").(bool),
		Nsgs:                     nsgs,
		Parent:                   d.Get("parent").(string),
		PrimaryType:              swag.String(d.Get("primary_type").(string)),
		QueryACL:                 queryACL,
		Tags:                     d.Get("tags"),
		TransferACL:              transferACL,
		UpdateACL:                updateACL,
		UseForwardersForSubzones: d.Get("use_forwarders_for_subzones").(bool),
		View:                     d.Get("view").(string),
		ZoneAuthority:            expandConfigZoneAuthority(d.Get("zone_authority").([]interface{})),
	}

	resp, err := c.DNSConfigurationAPI.AuthZone.AuthZoneCreate(
		&auth_zone.AuthZoneCreateParams{
			Body: az, Context: ctx,
		},
		nil,
	)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.Payload.Result.ID)

	return resourceConfigAuthZoneRead(ctx, d, m)
}

func resourceConfigAuthZoneRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	var diags diag.Diagnostics

	resp, err := c.DNSConfigurationAPI.AuthZone.AuthZoneRead(
		&auth_zone.AuthZoneReadParams{ID: d.Id(), Context: ctx},
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
	err = d.Set("disabled", resp.Payload.Result.Disabled)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	externalPrimaries := make([]interface{}, 0, len(resp.Payload.Result.ExternalPrimaries))
	for _, ep := range resp.Payload.Result.ExternalPrimaries {
		externalPrimaries = append(externalPrimaries, flattenConfigExternalPrimary(ep))
	}
	err = d.Set("external_primaries", externalPrimaries)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	externalSecondaries := make([]interface{}, 0, len(resp.Payload.Result.ExternalSecondaries))
	for _, es := range resp.Payload.Result.ExternalSecondaries {
		externalSecondaries = append(externalSecondaries, flattenConfigExternalSecondary(es))
	}
	err = d.Set("external_secondaries", externalSecondaries)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("fqdn", resp.Payload.Result.Fqdn)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("gss_tsig_enabled", resp.Payload.Result.GssTsigEnabled)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	inheritanceAssignedHosts := make([]interface{}, 0, len(resp.Payload.Result.InheritanceAssignedHosts))
	for _, iah := range resp.Payload.Result.InheritanceAssignedHosts {
		inheritanceAssignedHosts = append(inheritanceAssignedHosts, flattenInheritance2AssignedHost(iah))
	}
	err = d.Set("inheritance_assigned_hosts", inheritanceAssignedHosts)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("inheritance_sources", flattenConfigAuthZoneInheritance(resp.Payload.Result.InheritanceSources))
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("initial_soa_serial", resp.Payload.Result.InitialSoaSerial)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	internalSecondaries := make([]interface{}, 0, len(resp.Payload.Result.InternalSecondaries))
	for _, is := range resp.Payload.Result.InternalSecondaries {
		internalSecondaries = append(internalSecondaries, flattenConfigInternalSecondary(is))
	}
	err = d.Set("internal_secondaries", internalSecondaries)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("mapped_subnet", resp.Payload.Result.MappedSubnet)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("mapping", resp.Payload.Result.Mapping)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("notify", resp.Payload.Result.Notify)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("nsgs", resp.Payload.Result.Nsgs)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("parent", resp.Payload.Result.Parent)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("primary_type", resp.Payload.Result.PrimaryType)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("protocol_fqdn", resp.Payload.Result.ProtocolFqdn)
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
	err = d.Set("view", resp.Payload.Result.View)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("zone_authority", flattenConfigZoneAuthority(resp.Payload.Result.ZoneAuthority))
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	return diags
}

func flattenConfigAuthZone(r *models.ConfigAuthZone) []interface{} {
	if r == nil {
		return nil
	}

	externalPrimaries := make([]interface{}, 0, len(r.ExternalPrimaries))
	for _, ep := range r.ExternalPrimaries {
		externalPrimaries = append(externalPrimaries, flattenConfigExternalPrimary(ep))
	}

	externalSecondaries := make([]interface{}, 0, len(r.ExternalSecondaries))
	for _, es := range r.ExternalSecondaries {
		externalSecondaries = append(externalSecondaries, flattenConfigExternalSecondary(es))
	}

	inheritanceAssignedHosts := make([]interface{}, 0, len(r.InheritanceAssignedHosts))
	for _, iah := range r.InheritanceAssignedHosts {
		inheritanceAssignedHosts = append(inheritanceAssignedHosts, flattenInheritance2AssignedHost(iah))
	}

	internalSecondaries := make([]interface{}, 0, len(r.InternalSecondaries))
	for _, is := range r.InternalSecondaries {
		internalSecondaries = append(internalSecondaries, flattenConfigInternalSecondary(is))
	}

	queryACL := make([]interface{}, 0, len(r.QueryACL))
	for _, aclItem := range r.QueryACL {
		queryACL = append(queryACL, flattenConfigACLItem(aclItem))
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
			"disabled":                    r.Disabled,
			"external_primaries":          externalPrimaries,
			"external_secondaries":        externalSecondaries,
			"fqdn":                        r.Fqdn,
			"gss_tsig_enabled":            r.GssTsigEnabled,
			"inheritance_assigned_hosts":  inheritanceAssignedHosts,
			"inheritance_sources":         flattenConfigAuthZoneInheritance(r.InheritanceSources),
			"initial_soa_serial":          r.InitialSoaSerial,
			"internal_secondaries":        internalSecondaries,
			"mapped_subnet":               r.MappedSubnet,
			"mapping":                     r.Mapping,
			"notify":                      r.Notify,
			"nsgs":                        r.Nsgs,
			"parent":                      r.Parent,
			"primary_type":                r.PrimaryType,
			"protocol_fqdn":               r.ProtocolFqdn,
			"query_acl":                   queryACL,
			"tags":                        r.Tags,
			"transfer_acl":                transferACL,
			"update_acl":                  updateACL,
			"updated_at":                  r.UpdatedAt.String(),
			"use_forwarders_for_subzones": r.UseForwardersForSubzones,
			"view":                        r.View,
			"zone_authority":              flattenConfigZoneAuthority(r.ZoneAuthority),
		},
	}
}

func resourceConfigAuthZoneUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// ToDo Implement resourceConfigAuthZoneUpdate function
	return diags
}

func resourceConfigAuthZoneDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)
	_, err := c.DNSConfigurationAPI.AuthZone.AuthZoneDelete(
		&auth_zone.AuthZoneDeleteParams{ID: d.Id(), Context: ctx},
		nil,
	)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return nil
}

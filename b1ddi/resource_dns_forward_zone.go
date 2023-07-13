package b1ddi

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/dns_config/forward_zone"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// ConfigForwardZone ForwardZone
//
// # Forward zone
//
// swagger:model configForwardZone
func resourceConfigForwardZone() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceConfigForwardZoneCreate,
		ReadContext:   resourceConfigForwardZoneRead,
		UpdateContext: resourceConfigForwardZoneUpdate,
		DeleteContext: resourceConfigForwardZoneDelete,
		Schema: map[string]*schema.Schema{
			// Optional. Comment for zone configuration.
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Optional. Comment for zone configuration.",
			},

			// The timestamp when the object has been created.
			// Read Only: true
			// Format: date-time
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The timestamp when the object has been created.",
			},

			// Optional. _true_ to disable object. A disabled object is effectively non-existent when generating configuration.
			"disabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Optional. _true_ to disable object. A disabled object is effectively non-existent when generating configuration.",
			},

			// Optional. External DNS servers to forward to. Order is not significant.
			"external_forwarders": {
				Type:        schema.TypeList,
				Elem:        schemaConfigForwarder(),
				Optional:    true,
				Description: "Optional. External DNS servers to forward to. Order is not significant.",
			},

			// Optional. _true_ to only forward.
			"forward_only": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Optional. _true_ to only forward.",
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

			// The resource identifier.
			"hosts": {
				Type: schema.TypeList,
				Elem: schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				Description: "The resource identifier.",
			},

			// The resource identifier.
			"internal_forwarders": {
				Type: schema.TypeList,
				Elem: schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				Description: "The resource identifier.",
			},

			// Reverse zone network address in the following format: "ip-address/cidr".
			// Defaults to empty.
			// Read Only: true
			"mapped_subnet": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Reverse zone network address in the following format: \"ip-address/cidr\".\nDefaults to empty.",
			},

			// Read-only. Zone mapping type.
			// Allowed values:
			//  * _forward_,
			//  * _ipv4_reverse_.
			//  * _ipv6_reverse_.
			//
			// Defaults to _forward_.
			// Read Only: true
			"mapping": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Read-only. Zone mapping type.\nAllowed values:\n * _forward_,\n * _ipv4_reverse_.\n * _ipv6_reverse_.\n\nDefaults to _forward_.",
			},

			// The resource identifier.
			"nsgs": {
				Type: schema.TypeList,
				Elem: schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				Description: "The resource identifier.",
			},

			// The resource identifier.
			"parent": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The resource identifier.",
			},

			// Zone FQDN in punycode.
			// Read Only: true
			"protocol_fqdn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Zone FQDN in punycode.",
			},

			// Tagging specifics.
			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Tagging specifics.",
			},

			// The timestamp when the object has been updated. Equals to _created_at_ if not updated after creation.
			// Read Only: true
			// Format: date-time
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The timestamp when the object has been updated. Equals to _created_at_ if not updated after creation.",
			},

			// The resource identifier.
			"view": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The resource identifier.",
			},

			// The list of a forward zone warnings.
			// Read Only: true
			"warnings": {
				Type:        schema.TypeList,
				Elem:        schemaConfigWarning(),
				Computed:    true,
				Description: "The list of a forward zone warnings.",
			},
		},
	}
}

func resourceConfigForwardZoneCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	resp, err := c.DNSConfigurationAPI.ForwardZone.ForwardZoneCreate(
		&forward_zone.ForwardZoneCreateParams{
			Body:    nil,
			Context: ctx,
		},
		nil,
	)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(resp.Payload.Result.ID)

	time.Sleep(time.Second)

	return resourceConfigForwardZoneRead(ctx, d, m)
}

func resourceConfigForwardZoneRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)
	var diags diag.Diagnostics

	resp, err := c.DNSConfigurationAPI.ForwardZone.ForwardZoneRead(
		&forward_zone.ForwardZoneReadParams{
			ID:      d.Id(),
			Context: ctx,
		},
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
	err = d.Set("external_forwarders", flattenConfigForwarder(resp.Payload.Result.ExternalForwarders))
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("forward_only", resp.Payload.Result.ForwardOnly)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("hosts", resp.Payload.Result.Hosts)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("internal_forwarders", resp.Payload.Result.InternalForwarders)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("id", resp.Payload.Result.ID)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("fqdn", resp.Payload.Result.Fqdn)
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
	err = d.Set("protocol_fqdn", resp.Payload.Result.ProtocolFqdn)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("tags", resp.Payload.Result.Tags)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("updated_at", resp.Payload.Result.UpdatedAt.String())
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("view", resp.Payload.Result.View)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("warnings", flattenConfigWarning(resp.Payload.Result.Warnings))
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("mapping", resp.Payload.Result.Mapping)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("mapped_subnet", resp.Payload.Result.MappedSubnet)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	return diags
}

func flattenConfigForwardZone(r *models.ConfigForwardZone) []interface{} {
	if r == nil {
		return nil
	}

	return []interface{}{
		map[string]interface{}{
			"comment":             r.Comment,
			"created_at":          r.CreatedAt.String(),
			"disabled":            r.Disabled,
			"external_forwarders": flattenConfigForwarder(r.ExternalForwarders),
			"forward_only":        r.ForwardOnly,
			"hosts":               r.Hosts,
			"internal_forwarders": r.InternalForwarders,
			"nsgs":                r.Nsgs,
			"parent":              r.Parent,
			"protocol_fqdn":       r.ProtocolFqdn,
			"tags":                r.Tags,
			"updated_at":          r.UpdatedAt.String(),
			"view":                r.View,
			"warnings":            flattenConfigWarning(r.Warnings),
		},
	}
}

func resourceConfigForwardZoneUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	efs := make([]*models.ConfigForwarder, 0)

	for _, ef := range d.Get("external_forwarders").([]interface{}) {
		efs = append(efs, expandConfigForwarder(ef.(map[string]interface{})))
	}

	body := &models.ConfigForwardZone{
		Comment:            d.Get("comment").(string),
		Disabled:           d.Get("disabled").(bool),
		ExternalForwarders: efs,
		ForwardOnly:        d.Get("forward_only").(bool),
		Hosts:              d.Get("hosts").([]string),
		InternalForwarders: d.Get("internal_forwarders").([]string),
		Nsgs:               d.Get("nsgs").([]string),
		Parent:             d.Get("parent").(string),
		Tags:               d.Get("tags").(string),
		View:               d.Get("view").(string),
	}

	resp, err := c.DNSConfigurationAPI.ForwardZone.ForwardZoneUpdate(
		&forward_zone.ForwardZoneUpdateParams{
			Body:    body,
			ID:      d.Id(),
			Context: ctx,
		},
		nil,
	)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(resp.Payload.Result.ID)

	return resourceConfigForwardZoneRead(ctx, d, m)
}

func resourceConfigForwardZoneDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	_, err := c.DNSConfigurationAPI.ForwardZone.ForwardZoneDelete(
		&forward_zone.ForwardZoneDeleteParams{
			ID:      d.Id(),
			Context: ctx,
		},
		nil,
	)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}

package b1ddi

import (
	"context"
	"fmt"
	"time"

	"github.com/go-openapi/swag"
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
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "The resource identifier.",
			},

			// The resource identifier.
			"internal_forwarders": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
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
				Elem:        &schema.Schema{Type: schema.TypeString},
				Computed:    true,
				Description: "The list of a forward zone warnings.",
			},
		},
	}
}

func resourceConfigForwardZoneCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	externalForwarders := make([]*models.ConfigForwarder, 0)
	for _, ef := range d.Get("external_forwarders").([]interface{}) {
		if ef != nil {
			externalForwarders = append(externalForwarders, expandConfigForwarder(ef.(map[string]interface{})))
		}
	}

	nsgs := make([]string, 0)
	for _, n := range d.Get("nsgs").([]interface{}) {
		if n != nil {
			nsgs = append(nsgs, n.(string))
		}
	}

	hosts := make([]string, 0)
	for _, host := range d.Get("hosts").([]interface{}) {
		if host != nil {
			hosts = append(hosts, host.(string))
		}
	}

	internalForwarders := make([]string, 0)
	for _, intfwd := range d.Get("internal_forwarders").([]interface{}) {
		if intfwd != nil {
			internalForwarders = append(internalForwarders, intfwd.(string))
		}
	}

	fwdZone := &models.ConfigForwardZone{
		Comment:            d.Get("comment").(string),
		Disabled:           d.Get("disabled").(bool),
		ExternalForwarders: externalForwarders,
		ForwardOnly:        d.Get("forward_only").(bool),
		Fqdn:               swag.String(d.Get("fqdn").(string)),
		Hosts:              hosts,
		InternalForwarders: internalForwarders,
		Nsgs:               nsgs,
		Tags:               d.Get("tags"),
		View:               d.Get("view").(string),
	}
	resp, err := c.DNSConfigurationAPI.ForwardZone.ForwardZoneCreate(
		&forward_zone.ForwardZoneCreateParams{
			Body: fwdZone, Context: ctx,
		},
		nil,
	)
	if err != nil {
		return diag.FromErr(err)
	}

	// Wait for API to create the Forward Zone
	time.Sleep(time.Second * 2)

	d.SetId(resp.Payload.Result.ID)

	return resourceConfigForwardZoneRead(ctx, d, m)
}

func resourceConfigForwardZoneRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	var diags diag.Diagnostics

	resp, err := c.DNSConfigurationAPI.ForwardZone.ForwardZoneRead(
		&forward_zone.ForwardZoneReadParams{ID: d.Id(), Context: ctx},
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
	externalForwarders := make([]map[string]interface{}, 0, len(resp.Payload.Result.ExternalForwarders))
	for _, ef := range resp.Payload.Result.ExternalForwarders {
		externalForwarders = append(externalForwarders, flattenConfigForwarder(ef))
	}
	err = d.Set("forward_only", resp.Payload.Result.ForwardOnly)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("fqdn", resp.Payload.Result.Fqdn)
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
	err = d.Set("mapped_subnet", resp.Payload.Result.MappedSubnet)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("mapping", resp.Payload.Result.Mapping)
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
	return diags
}

func resourceConfigForwardZoneUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	if d.HasChange("fqdn") {
		d.Partial(true)
		return diag.FromErr(fmt.Errorf("changing the value of 'fqdn' field is not allowed"))
	}

	if d.HasChange("view") {
		d.Partial(true)
		return diag.FromErr(fmt.Errorf("changing the value of 'view' field is not allowed"))
	}

	externalForwarders := make([]*models.ConfigForwarder, 0)
	for _, ef := range d.Get("external_forwarders").([]interface{}) {
		if ef != nil {
			externalForwarders = append(externalForwarders, expandConfigForwarder(ef.(map[string]interface{})))
		}
	}

	nsgs := make([]string, 0)
	for _, n := range d.Get("nsgs").([]interface{}) {
		if n != nil {
			nsgs = append(nsgs, n.(string))
		}
	}

	internalForwarders := make([]string, 0)
	for _, n := range d.Get("internal_forwarders").([]interface{}) {
		if n != nil {
			internalForwarders = append(internalForwarders, n.(string))
		}
	}

	hosts := make([]string, 0)
	for _, host := range d.Get("hosts").([]interface{}) {
		if host != nil {
			hosts = append(hosts, host.(string))
		}
	}

	body := &models.ConfigForwardZone{
		Comment:            d.Get("comment").(string),
		Disabled:           d.Get("disabled").(bool),
		ExternalForwarders: externalForwarders,
		ForwardOnly:        d.Get("forward_only").(bool),
		Hosts:              hosts,
		InternalForwarders: internalForwarders,
		Nsgs:               nsgs,
		Tags:               d.Get("tags"),
	}

	resp, err := c.DNSConfigurationAPI.ForwardZone.ForwardZoneUpdate(
		&forward_zone.ForwardZoneUpdateParams{ID: d.Id(), Body: body, Context: ctx},
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
		&forward_zone.ForwardZoneDeleteParams{ID: d.Id(), Context: ctx},
		nil,
	)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return nil
}

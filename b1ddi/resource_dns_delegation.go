package b1ddi

import (
	"context"
	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/dns_config/delegation"
	"time"

	"github.com/infobloxopen/b1ddi-go-client/models"
)

// DNSDelegation Delegation
//
// DNS zone delegation.
//
// swagger:model configDelegation
func resourceDNSDelegation() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDNSDelegationCreate,
		ReadContext:   resourceDNSDelegationRead,
		UpdateContext: resourceDNSDelegationUpdate,
		DeleteContext: resourceDNSDelegationDelete,
		Schema: map[string]*schema.Schema{

			// Optional. Comment for zone delegation.
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Optional. Comment for zone delegation.",
			},

			// Required. DNS zone delegation servers. Order is not significant.
			"delegation_servers": {
				Type:        schema.TypeList,
				Elem:        resourceDNSDelegationServer(),
				Optional:    true,
				Description: "Required. DNS zone delegation servers. Order is not significant.",
			},

			// Optional. _true_ to disable object. A disabled object is effectively non-existent when generating resource records.
			"disabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Optional. _true_ to disable object. A disabled object is effectively non-existent when generating resource records.",
			},

			// Delegation FQDN.
			// The FQDN supplied at creation will be converted to canonical form.
			//
			// Read-only after creation.
			// Required: true
			"fqdn": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Delegation FQDN.\nThe FQDN supplied at creation will be converted to canonical form.\n\nRead-only after creation.",
			},

			// The resource identifier.
			"parent": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The resource identifier.",
			},

			// Delegation FQDN in punycode.
			// Read Only: true
			"protocol_fqdn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Delegation FQDN in punycode.",
			},

			// Tagging specifics.
			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Tagging specifics.",
			},

			// The resource identifier.
			"view": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The resource identifier.",
			},
		},
	}
}

func resourceDNSDelegationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	ds := make([]*models.ConfigDelegationServer, 0)

	for _, server := range d.Get("delegaton_servers").([]interface{}) {
		if server != nil {
			ds = append(ds, expandDNSDelegationServer(server.(map[string]interface{})))
		}
	}

	body := &models.ConfigDelegation{
		Comment:           d.Get("comment").(string),
		DelegationServers: ds,
		Disabled:          d.Get("disabled").(bool),
		Fqdn:              swag.String(d.Get("fqdn").(string)),
		Parent:            d.Get("parent").(string),
		Tags:              d.Get("tags").(string),
		View:              d.Get("view").(string),
	}
	resp, err := c.DNSConfigurationAPI.Delegation.DelegationCreate(
		&delegation.DelegationCreateParams{
			Body:    body,
			Context: ctx,
		},
		nil,
	)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(resp.Payload.Result.ID)

	time.Sleep(time.Second)

	return resourceDNSDelegationRead(ctx, d, m)
}

func resourceDNSDelegationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)
	var diags diag.Diagnostics

	resp, err := c.DNSConfigurationAPI.Delegation.DelegationRead(
		&delegation.DelegationReadParams{
			ID:      d.Id(),
			Context: ctx,
		},
		nil,
	)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("id", resp.Payload.Result.ID)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("comment", resp.Payload.Result.Comment)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("fqdn", resp.Payload.Result.Fqdn)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("prent", resp.Payload.Result.Parent)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("protocol_fqdn", resp.Payload.Result.ProtocolFqdn)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("view", resp.Payload.Result.View)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("tags", resp.Payload.Result.Tags)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("disabled", resp.Payload.Result.Disabled)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("delegation_servers", flattenDNSDelegationServer(resp.Payload.Result.DelegationServers))
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	return diags
}

func flattenDNSDelegation(r *models.ConfigDelegation) []interface{} {
	if r == nil {
		return nil
	}

	return []interface{}{
		map[string]interface{}{
			"comment":            r.Comment,
			"delegation_servers": flattenDNSDelegationServer(r.DelegationServers),
			"disabled":           r.Disabled,
			"fqdn":               r.Fqdn,
			"id":                 r.ID,
			"parent":             r.Parent,
			"protocol_fqdn":      r.ProtocolFqdn,
			"tags":               r.Tags,
			"view":               r.View,
		},
	}
}

func resourceDNSDelegationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	ds := make([]*models.ConfigDelegationServer, 0)
	for _, server := range d.Get("delegation_servers").([]interface{}) {
		ds = append(ds, expandDNSDelegationServer(server.(map[string]interface{})))
	}

	body := &models.ConfigDelegation{
		Comment:           d.Get("comment").(string),
		DelegationServers: ds,
		Disabled:          d.Get("disabled").(bool),
		Parent:            d.Get("parent").(string),
		Tags:              d.Get("tags").(string),
		View:              d.Get("view").(string),
	}
	resp, err := c.DNSConfigurationAPI.Delegation.DelegationUpdate(
		&delegation.DelegationUpdateParams{
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

	return resourceDNSDelegationRead(ctx, d, m)
}

func resourceDNSDelegationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	_, err := c.DNSConfigurationAPI.Delegation.DelegationDelete(
		&delegation.DelegationDeleteParams{
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

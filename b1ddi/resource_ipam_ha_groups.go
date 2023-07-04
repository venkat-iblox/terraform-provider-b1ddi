package b1ddi

import (
	"context"
	"fmt"
	"time"

	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/dns_data/record"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/ha_group"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// DataRecord Record
//
// A __Record__ object (_dns/record_) represents a DNS resource record in an authoritative zone.
//
// For creating a DNS resource record, one of the following pairs of fields is required:<ul><li>_name_in_zone_ and _zone_: The system creates the DNS resource record object within the specified zone. The value of the view field is automatically retrieved from the zone object.</li><li>_absolute_name_spec_ and _view_: The system looks for the appropriate zone in the provided view to create the DNS resource record object. The value of the zone field is automatically computed as part of this process.</li></ul>
//
// The _zone_ and _view_ fields cannot be modified while updating a DNS resource record. The _name_in_zone_ and _absolute_name_spec_ fields can be modified. If both fields are modified in the same update, they need to represent the same change.
//
// swagger:model dataRecord
func resourceHAGroups() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceHAGroupCreate,
		ReadContext:   resourceHAGroupRead,
		UpdateContext: resourceHAGroupUpdate,
		DeleteContext: resourceHAGroupDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{

			// Synthetic field, used to determine _zone_ and/or _name_in_zone_ field for records.
			"id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Synthetic field, used to determine _zone_ and/or _name_in_zone_ field for records.",
			},

			// The absolute domain name of the zone where this record belongs.
			// Read Only: true
			"ip_space": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The absolute domain name of the zone where this record belongs.",
			},

			// The description for the DNS resource record. May contain 0 to 1024 characters. Can include UTF-8.
			"mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description for the DNS resource record. May contain 0 to 1024 characters. Can include UTF-8.",
			},

			// The timestamp when the object has been created.
			// Read Only: true
			// Format: date-time
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The timestamp when the object has been created.",
			},

			// The resource identifier.
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The resource identifier.",
			},

			// The tags for the DNS resource record in JSON format.
			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "The tags for the DNS resource record in JSON format.",
			},

			// The timestamp when the object has been updated. Equals to _created_at_ if not updated after creation.
			// Read Only: true
			// Format: date-time
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The timestamp when the object has been updated. Equals to _created_at_ if not updated after creation.",
			},

			"comment": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "comment",
			},
			"anycast_config_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Anycast config ID",
			},

			"hosts": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     schemaHAGroupsHosts(),
			},
		},
	}
}

func resourceHAGroupCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	hosts := make([]*models.IpamsvcHAGroupHost, 0)
	for _, o := range d.Get("hosts").([]interface{}) {
		if o != nil {
			hosts = append(hosts, expandHAGroupsHosts(o.(map[string]interface{})))
		}
	}

	r := &models.IpamsvcHAGroup{
		Comment: d.Get("comment").(string),
		Tags:    d.Get("tags"),
		Mode:    d.Get("mode").(string),
		Name:    swag.String(d.Get("name").(string)),
		Hosts:   hosts,
	}

	resp, err := c.IPAddressManagementAPI.HaGroup.HaGroupCreate(
		&ha_group.HaGroupCreateParams{Body: r, Context: ctx},
		nil,
	)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.Payload.Result.ID)

	time.Sleep(time.Second)

	return resourceDataRecordRead(ctx, d, m)
}

func resourceHAGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	var diags diag.Diagnostics

	resp, err := c.IPAddressManagementAPI.HaGroup.HaGroupRead(
		&ha_group.HaGroupReadParams{ID: d.Id(), Context: ctx},
		nil,
	)
	if err != nil {
		return diag.FromErr(err)
	}

	err = d.Set("name", resp.Payload.Result.Name)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("mode", resp.Payload.Result.Mode)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("comment", resp.Payload.Result.Comment)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("created_at", resp.Payload.Result.CreatedAt.String())
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("ip_space", resp.Payload.Result.IPSpace)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("id", resp.Payload.Result.ID)
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
	err = d.Set("hosts", resp.Payload.Result.Hosts)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("status", resp.Payload.Result.Status)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	return diags
}

func resourceHAGroupUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	if d.HasChange("type") {
		d.Partial(true)
		return diag.FromErr(fmt.Errorf("changing the value of 'type' field is not allowed"))
	}

	if d.HasChange("view") {
		d.Partial(true)
		return diag.FromErr(fmt.Errorf("changing the value of 'view' field is not allowed"))
	}

	if d.HasChange("zone") {
		d.Partial(true)
		return diag.FromErr(fmt.Errorf("changing the value of 'zone' field is not allowed"))
	}

	body := &models.DataRecord{
		Comment:            d.Get("comment").(string),
		Delegation:         d.Get("delegation").(string),
		Disabled:           d.Get("disabled").(bool),
		InheritanceSources: expandDataRecordInheritance(d.Get("inheritance_sources").([]interface{})),
		NameInZone:         d.Get("name_in_zone").(string),
		Options:            d.Get("options"),
		Rdata:              d.Get("rdata"),
		Tags:               d.Get("tags"),
		TTL:                int64(d.Get("ttl").(int)),
	}

	resp, err := c.DNSDataAPI.Record.RecordUpdate(
		&record.RecordUpdateParams{ID: d.Id(), Body: body, Context: ctx},
		nil,
	)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.Payload.Result.ID)

	return resourceDataRecordRead(ctx, d, m)
}

func resourceHAGroupDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)
	_, err := c.DNSDataAPI.Record.RecordDelete(
		&record.RecordDeleteParams{ID: d.Id(), Context: ctx},
		nil,
	)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return nil
}

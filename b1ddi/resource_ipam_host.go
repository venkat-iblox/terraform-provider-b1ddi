package b1ddi

import (
	"context"
	"time"

	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/ipam_host"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

func resourceIpamHost() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIpamHostCreate,
		ReadContext:   resourceIpamHostRead,
		UpdateContext: resourceIpamHostUpdate,
		DeleteContext: resourceIpamHostDelete,
		Schema: map[string]*schema.Schema{

			// The list of all addresses associated with the IPAM host, which may be in different IP spaces.
			"addresses": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcHostAddress(),
				Optional:    true,
				Description: "The list of all addresses associated with the IPAM host, which may be in different IP spaces.",
			},

			// This flag specifies if resource records have to be auto generated for the host.
			"auto_generate_records": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "This flag specifies if resource records have to be auto generated for the host.",
			},

			// The description for the IPAM host. May contain 0 to 1024 characters. Can include UTF-8.
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description for the IPAM host. May contain 0 to 1024 characters. Can include UTF-8.",
			},

			// Time when the object has been created.
			// Read Only: true
			// Format: date-time
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Time when the object has been created.",
			},

			// The name records to be generated for the host.
			//
			// This field is required if _auto_generate_records_ is true.
			"host_names": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcHostName(),
				Optional:    true,
				Description: "The name records to be generated for the host.\n\nThis field is required if _auto_generate_records_ is true.",
			},

			// The name of the IPAM host. Must contain 1 to 256 characters. Can include UTF-8.
			// Required: true
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the IPAM host. Must contain 1 to 256 characters. Can include UTF-8.",
			},

			// The tags for the IPAM host in JSON format.
			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "The tags for the IPAM host in JSON format.",
			},

			// Time when the object has been updated. Equals to _created_at_ if not updated after creation.
			// Read Only: true
			// Format: date-time
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Time when the object has been updated. Equals to _created_at_ if not updated after creation.",
			},
		},
	}
}

func resourceIpamHostCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	addrs := make([]*models.IpamsvcHostAddress, 0)
	for _, a := range d.Get("addresses").([]interface{}) {
		if a != nil {
			addrs = append(addrs, expandIpamHostAddress(a.(map[string]interface{})))
		}
	}

	hnames := make([]*models.IpamsvcHostName, 0)
	for _, a := range d.Get("host_names").([]interface{}) {
		if a != nil {
			addrs = append(addrs, expandIpamHostAddress(a.(map[string]interface{})))
		}
	}

	body := &models.IpamsvcIpamHost{
		Addresses:           addrs,
		AutoGenerateRecords: d.Get("auto_generated_records").(bool),
		Comment:             d.Get("comment").(string),
		HostNames:           hnames,
		Name:                swag.String(d.Get("name").(string)),
		Tags:                d.Get("tags").(string),
	}
	resp, err := c.IPAddressManagementAPI.IpamHost.IpamHostCreate(&ipam_host.IpamHostCreateParams{
		Body:    body,
		Context: ctx,
	}, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.Payload.Result.ID)

	time.Sleep(time.Second)

	return resourceIpamHostRead(ctx, d, m)
}

func resourceIpamHostRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	var diags diag.Diagnostics

	resp, err := c.IPAddressManagementAPI.IpamHost.IpamHostRead(
		&ipam_host.IpamHostReadParams{
			ID:      d.Id(),
			Context: ctx,
		},
		nil,
	)
	if err != nil {
		return diag.FromErr(err)
	}

	err = d.Set("name", resp.Payload.Result.Name)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("comment", resp.Payload.Result.Comment)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("auto_generate_records", resp.Payload.Result.AutoGenerateRecords)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("created_at", resp.Payload.Result.CreatedAt.String())
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
	err = d.Set("addresses", flattenIpamHostAddress(resp.Payload.Result.Addresses))
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("host_names", flattenIpamHostName(resp.Payload.Result.HostNames))
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	return diags
}

func resourceIpamHostUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	addrs := make([]*models.IpamsvcHostAddress, 0)
	for _, a := range d.Get("addresses").([]interface{}) {
		if a != nil {
			addrs = append(addrs, expandIpamHostAddress(a.(map[string]interface{})))
		}
	}

	hnames := make([]*models.IpamsvcHostName, 0)
	for _, a := range d.Get("host_names").([]interface{}) {
		if a != nil {
			addrs = append(addrs, expandIpamHostAddress(a.(map[string]interface{})))
		}
	}
	body := &models.IpamsvcIpamHost{
		Addresses:           addrs,
		AutoGenerateRecords: d.Get("auto_generated_records").(bool),
		Comment:             d.Get("comment").(string),
		HostNames:           hnames,
		Name:                swag.String(d.Get("name").(string)),
		Tags:                d.Get("tags").(string),
	}

	resp, err := c.IPAddressManagementAPI.IpamHost.IpamHostUpdate(
		&ipam_host.IpamHostUpdateParams{
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

	return resourceIpamHostRead(ctx, d, m)
}

func resourceIpamHostDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)
	_, err := c.IPAddressManagementAPI.IpamHost.IpamHostDelete(&ipam_host.IpamHostDeleteParams{
		ID:      d.Id(),
		Context: ctx,
	}, nil)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return nil
}

func flattenIpamsvcIpamHost(r *models.IpamsvcIpamHost) []interface{} {
	if r == nil {
		return nil
	}

	return []interface{}{
		map[string]interface{}{
			"addresses":             flattenIpamHostAddress(r.Addresses),
			"auto_generate_records": r.AutoGenerateRecords,
			"comment":               r.Comment,
			"created_at":            r.CreatedAt.String(),
			"host_names":            flattenIpamHostName(r.HostNames),
			"id":                    r.ID,
			"name":                  r.Name,
			"tags":                  r.Tags,
			"updated_at":            r.UpdatedAt.String(),
		},
	}
}

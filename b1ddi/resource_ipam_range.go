package b1ddi

import (
	"context"
	"github.com/go-openapi/swag"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/range_operations"
	"github.com/infobloxopen/b1ddi-go-client/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// IpamsvcRange Range
//
// A __Range__ object (_ipam/range_) represents a set of contiguous IP addresses in the same IP space with no gap, expressed as a (start, end) pair within a given subnet that are grouped together for administrative purpose and protocol management. The start and end values are not required to align with CIDR boundaries.
//
// swagger:model ipamsvcRange
func resourceIpamsvcRange() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIpamsvcRangeCreate,
		ReadContext:   resourceIpamsvcRangeRead,
		UpdateContext: resourceIpamsvcRangeUpdate,
		DeleteContext: resourceIpamsvcRangeDelete,
		Schema: map[string]*schema.Schema{

			// The description for the range. May contain 0 to 1024 characters. Can include UTF-8.
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description for the range. May contain 0 to 1024 characters. Can include UTF-8.",
			},

			// Time when the object has been created.
			// Read Only: true
			// Format: date-time
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Time when the object has been created.",
			},

			// The resource identifier.
			"dhcp_host": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The resource identifier.",
			},

			// The list of DHCP options. May be either a specific option or a group of options.
			"dhcp_options": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcOptionItem(),
				Optional:    true,
				Description: "The list of DHCP options. May be either a specific option or a group of options.",
			},

			// The end IP address of the range.
			// Required: true
			"end": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The end IP address of the range.",
			},

			// The list of all exclusion ranges in the scope of the range.
			"exclusion_ranges": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcExclusionRange(),
				Optional:    true,
				Description: "The list of all exclusion ranges in the scope of the range.",
			},

			// The list of the inheritance assigned hosts of the object.
			// Read Only: true
			"inheritance_assigned_hosts": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceAssignedHost(),
				Computed:    true,
				Description: "The list of the inheritance assigned hosts of the object.",
			},

			// The resource identifier.
			"inheritance_parent": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The resource identifier.",
			},

			// The DHCP inheritance configuration for the range.
			"inheritance_sources": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcDHCPOptionsInheritance(),
				MaxItems:    1,
				Optional:    true,
				Description: "The DHCP inheritance configuration for the range.",
			},

			// The name of the range. May contain 1 to 256 characters. Can include UTF-8.
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the range. May contain 1 to 256 characters. Can include UTF-8.",
			},

			// The resource identifier.
			"parent": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The resource identifier.",
			},

			// The type of protocol (_ipv4_ or _ipv6_).
			// Read Only: true
			"protocol": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of protocol (_ipv4_ or _ipv6_).",
			},

			// The resource identifier.
			// Required: true
			"space": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The resource identifier.",
			},

			// The start IP address of the range.
			// Required: true
			"start": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The start IP address of the range.",
			},

			// The tags for the range in JSON format.
			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "The tags for the range in JSON format.",
			},

			// The utilization threshold settings for the range.
			"threshold": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcUtilizationThreshold(),
				MaxItems:    1,
				Optional:    true,
				Description: "The utilization threshold settings for the range.",
			},

			// Time when the object has been updated. Equals to _created_at_ if not updated after creation.
			// Read Only: true
			// Format: date-time
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Time when the object has been updated. Equals to _created_at_ if not updated after creation.",
			},

			// The utilization statistics for the range.
			// Read Only: true
			"utilization": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcUtilization(),
				Computed:    true,
				Description: "The utilization statistics for the range.",
			},
		},
	}
}

func resourceIpamsvcRangeCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ipamsvc.IPAddressManagementAPI)

	r := &models.IpamsvcRange{
		Comment: d.Get("comment").(string),
		End:     swag.String(d.Get("end").(string)),
		Name:    d.Get("name").(string),
		Space:   swag.String(d.Get("space").(string)),
		Start:   swag.String(d.Get("start").(string)),
	}

	resp, err := c.RangeOperations.RangeCreate(&range_operations.RangeCreateParams{Body: r, Context: ctx}, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.Payload.Result.ID)

	return resourceIpamsvcRangeRead(ctx, d, m)
}

func resourceIpamsvcRangeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ipamsvc.IPAddressManagementAPI)

	var diags diag.Diagnostics

	resp, err := c.RangeOperations.RangeRead(&range_operations.RangeReadParams{
		ID:      d.Id(),
		Context: ctx,
	}, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	err = d.Set("comment", resp.Payload.Result.Comment)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("end", resp.Payload.Result.End)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	inheritanceAssignedHosts := make([]interface{}, 0, len(resp.Payload.Result.InheritanceAssignedHosts))
	for _, inheritanceAssignedHost := range resp.Payload.Result.InheritanceAssignedHosts {
		inheritanceAssignedHosts = append(inheritanceAssignedHosts, flattenInheritanceAssignedHost(inheritanceAssignedHost))
	}
	err = d.Set("inheritance_assigned_hosts", inheritanceAssignedHosts)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	err = d.Set("name", resp.Payload.Result.Name)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("space", resp.Payload.Result.Space)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("start", resp.Payload.Result.Start)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("utilization", flattenIpamsvcUtilization(resp.Payload.Result.Utilization))
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	return diags
}

func resourceIpamsvcRangeUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// ToDo Implement resourceIpamsvcRangeUpdate function
	return diags
}

func resourceIpamsvcRangeDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ipamsvc.IPAddressManagementAPI)

	_, err := c.RangeOperations.RangeDelete(&range_operations.RangeDeleteParams{ID: d.Id(), Context: ctx}, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

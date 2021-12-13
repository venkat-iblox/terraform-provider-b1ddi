package b1ddi

import (
	"context"
	"github.com/go-openapi/swag"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/fixed_address"
	"github.com/infobloxopen/b1ddi-go-client/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// IpamsvcFixedAddress FixedAddress
//
// A __FixedAddress__ object (_dhcp/fixed_address_) reserves an address for a specific client. It must have a _match_type_ and a valid corresponding _match_value_ so it can match that client.
//
// swagger:model ipamsvcFixedAddress
func resourceIpamsvcFixedAddress() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIpamsvcFixedAddressCreate,
		ReadContext:   resourceIpamsvcFixedAddressRead,
		UpdateContext: resourceIpamsvcFixedAddressUpdate,
		DeleteContext: resourceIpamsvcFixedAddressDelete,
		Schema: map[string]*schema.Schema{

			// The reserved address.
			// Required: true
			"address": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The reserved address.",
			},

			// The description for the fixed address. May contain 0 to 1024 characters. Can include UTF-8.
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description for the fixed address. May contain 0 to 1024 characters. Can include UTF-8.",
			},

			// Time when the object has been created.
			// Read Only: true
			// Format: date-time
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Time when the object has been created.",
			},

			// The list of DHCP options. May be either a specific option or a group of options.
			"dhcp_options": {
				Type:        schema.TypeList,
				Elem:        dataSourceIpamsvcOptionItem(),
				Optional:    true,
				Description: "The list of DHCP options. May be either a specific option or a group of options.",
			},

			// The configuration for header option filename field.
			"header_option_filename": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The configuration for header option filename field.",
			},

			// The configuration for header option server address field.
			"header_option_server_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The configuration for header option server address field.",
			},

			// The configuration for header option server name field.
			"header_option_server_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The configuration for header option server name field.",
			},

			// The DHCP host name associated with this fixed address. It is of FQDN type and it defaults to empty.
			"hostname": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The DHCP host name associated with this fixed address. It is of FQDN type and it defaults to empty.",
			},

			// The list of the inheritance assigned hosts of the object.
			// Read Only: true
			// ToDo add inheritance_assigned_hosts
			//"inheritance_assigned_hosts": {
			//	Type:        schema.TypeList,
			//	Elem:        dataSourceInheritanceAssignedHost(),
			//	Computed:    true,
			//	Description: "The list of the inheritance assigned hosts of the object.",
			//},

			// The resource identifier.
			"inheritance_parent": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The resource identifier.",
			},

			// The inheritance configuration.
			"inheritance_sources": {
				Type:        schema.TypeList,
				Elem:        dataSourceIpamsvcFixedAddressInheritance(),
				MaxItems:    1,
				Optional:    true,
				Description: "The inheritance configuration.",
			},

			// The resource identifier.
			"ip_space": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The resource identifier.",
			},

			// Indicates how to match the client:
			//  * _mac_: match the client MAC address,
			//  * _client_text_ or _client_hex_: match the client identifier,
			//  * _relay_text_ or _relay_hex_: match the circuit ID or remote ID in the DHCP relay agent option (82).
			// Required: true
			"match_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Indicates how to match the client:\n * _mac_: match the client MAC address,\n * _client_text_ or _client_hex_: match the client identifier,\n * _relay_text_ or _relay_hex_: match the circuit ID or remote ID in the DHCP relay agent option (82).",
			},

			// The value to match.
			// Required: true
			"match_value": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The value to match.",
			},

			// The name of the fixed address. May contain 1 to 256 characters. Can include UTF-8.
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the fixed address. May contain 1 to 256 characters. Can include UTF-8.",
			},

			// The resource identifier.
			"parent": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The resource identifier.",
			},

			// The tags for the fixed address in JSON format.
			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "The tags for the fixed address in JSON format.",
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

func resourceIpamsvcFixedAddressCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ipamsvc.IPAddressManagementAPI)

	var diags diag.Diagnostics

	fa := &models.IpamsvcFixedAddress{
		Address:    swag.String(d.Get("address").(string)),
		Name:       d.Get("name").(string),
		MatchType:  swag.String(d.Get("match_type").(string)),
		MatchValue: swag.String(d.Get("match_value").(string)),
		IPSpace:    d.Get("ip_space").(string),
		Comment:    d.Get("comment").(string),
	}

	resp, err := c.FixedAddress.FixedAddressCreate(&fixed_address.FixedAddressCreateParams{Body: fa, Context: ctx}, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.Payload.Result.ID)

	resourceIpamsvcFixedAddressRead(ctx, d, m)

	return diags
}

func resourceIpamsvcFixedAddressRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ipamsvc.IPAddressManagementAPI)

	var diags diag.Diagnostics

	fa, err := c.FixedAddress.FixedAddressRead(&fixed_address.FixedAddressReadParams{ID: d.Id(), Context: ctx}, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	err = d.Set("address", fa.Payload.Result.Address)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("match_type", fa.Payload.Result.MatchType)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("match_value", fa.Payload.Result.MatchValue)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("comment", fa.Payload.Result.Comment)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("created_at", fa.Payload.Result.CreatedAt.String())
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("updated_at", fa.Payload.Result.UpdatedAt.String())
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceIpamsvcFixedAddressUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// ToDo Implement resourceIpamsvcFixedAddressUpdate function
	return diags
}

func resourceIpamsvcFixedAddressDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// ToDo Implement resourceIpamsvcFixedAddressDelete function
	return diags
}

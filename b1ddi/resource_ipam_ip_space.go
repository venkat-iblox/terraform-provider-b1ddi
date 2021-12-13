package b1ddi

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/ip_space"
	b1models "github.com/infobloxopen/b1ddi-go-client/models"
)

// IpamsvcIPSpace IPSpace
//
// An __IPSpace__ object (_ipam/ip_space_) allows customers to represent their entire managed address space with no collision. A collision arises when two or more block of addresses overlap partially or fully.
//
// swagger:model ipamsvcIPSpace
func resourceIpamsvcIPSpace() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIpamsvcIPSpaceCreate,
		ReadContext:   resourceIpamsvcIPSpaceRead,
		UpdateContext: resourceIpamsvcIPSpaceUpdate,
		DeleteContext: resourceIpamsvcIPSpaceDelete,
		Schema: map[string]*schema.Schema{

			// The Automated Scope Management configuration for the IP space.
			"asm_config": {
				Type:        schema.TypeList,
				Elem:        dataSourceIpamsvcASMConfig(),
				MaxItems:    1,
				Optional:    true,
				Description: "The Automated Scope Management configuration for the IP space.",
			},

			// The number of times the automated scope management usage limits have been exceeded for any of the subnets in this IP space.
			// Read Only: true
			"asm_scope_flag": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of times the automated scope management usage limits have been exceeded for any of the subnets in this IP space.",
			},

			// The description for the IP space. May contain 0 to 1024 characters. Can include UTF-8.
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description for the IP space. May contain 0 to 1024 characters. Can include UTF-8.",
			},

			// Time when the object has been created.
			// Read Only: true
			// Format: date-time
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Time when the object has been created.",
			},

			// Controls who does the DDNS updates.
			//
			// Valid values are:
			// * _client_: DHCP server updates DNS if requested by client.
			// * _server_: DHCP server always updates DNS, overriding an update request from the client, unless the client requests no updates.
			// * _ignore_: DHCP server always updates DNS, even if the client says not to.
			// * _over_client_update_: Same as _server_. DHCP server always updates DNS, overriding an update request from the client, unless the client requests no updates.
			// * _over_no_update_: DHCP server updates DNS even if the client requests that no updates be done. If the client requests to do the update, DHCP server allows it.
			//
			// Defaults to _client_.
			"ddns_client_update": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Controls who does the DDNS updates.\n\nValid values are:\n* _client_: DHCP server updates DNS if requested by client.\n* _server_: DHCP server always updates DNS, overriding an update request from the client, unless the client requests no updates.\n* _ignore_: DHCP server always updates DNS, even if the client says not to.\n* _over_client_update_: Same as _server_. DHCP server always updates DNS, overriding an update request from the client, unless the client requests no updates.\n* _over_no_update_: DHCP server updates DNS even if the client requests that no updates be done. If the client requests to do the update, DHCP server allows it.\n\nDefaults to _client_.",
			},

			// The domain suffix for DDNS updates. FQDN, may be empty.
			//
			// Defaults to empty.
			"ddns_domain": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The domain suffix for DDNS updates. FQDN, may be empty.\n\nDefaults to empty.",
			},

			// Indicates if DDNS needs to generate a hostname when not supplied by the client.
			//
			// Defaults to _false_.
			"ddns_generate_name": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates if DDNS needs to generate a hostname when not supplied by the client.\n\nDefaults to _false_.",
			},

			// The prefix used in the generation of an FQDN.
			//
			// When generating a name, DHCP server will construct the name in the format: [ddns-generated-prefix]-[address-text].[ddns-qualifying-suffix].
			// where address-text is simply the lease IP address converted to a hyphenated string.
			//
			// Defaults to "myhost".
			"ddns_generated_prefix": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The prefix used in the generation of an FQDN.\n\nWhen generating a name, DHCP server will construct the name in the format: [ddns-generated-prefix]-[address-text].[ddns-qualifying-suffix].\nwhere address-text is simply the lease IP address converted to a hyphenated string.\n\nDefaults to \"myhost\".",
			},

			// Determines if DDNS updates are enabled at the IP space level.
			// Defaults to _true_.
			"ddns_send_updates": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Determines if DDNS updates are enabled at the IP space level.\nDefaults to _true_.",
			},

			// Instructs the DHCP server to always update the DNS information when a lease is renewed even if its DNS information has not changed.
			//
			// Defaults to _false_.
			"ddns_update_on_renew": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Instructs the DHCP server to always update the DNS information when a lease is renewed even if its DNS information has not changed.\n\nDefaults to _false_.",
			},

			// When true, DHCP server will apply conflict resolution, as described in RFC 4703, when attempting to fulfill the update request.
			//
			// When false, DHCP server will simply attempt to update the DNS entries per the request, regardless of whether or not they conflict with existing entries owned by other DHCP4 clients.
			//
			// Defaults to _true_.
			"ddns_use_conflict_resolution": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "When true, DHCP server will apply conflict resolution, as described in RFC 4703, when attempting to fulfill the update request.\n\nWhen false, DHCP server will simply attempt to update the DNS entries per the request, regardless of whether or not they conflict with existing entries owned by other DHCP4 clients.\n\nDefaults to _true_.",
			},

			// The shared DHCP configuration for the IP space that controls how leases are issued.
			"dhcp_config": {
				Type:        schema.TypeList,
				Elem:        dataSourceIpamsvcDHCPConfig(),
				MaxItems:    1,
				Optional:    true,
				Description: "The shared DHCP configuration for the IP space that controls how leases are issued.",
			},

			// The list of DHCP options for the IP space. May be either a specific option or a group of options.
			"dhcp_options": {
				Type:        schema.TypeList,
				Elem:        dataSourceIpamsvcOptionItem(),
				Optional:    true,
				Description: "The list of DHCP options for the IP space. May be either a specific option or a group of options.",
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

			// The character to replace non-matching characters with, when hostname rewrite is enabled.
			//
			// Any single ASCII character.
			//
			// Defaults to "_".
			"hostname_rewrite_char": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The character to replace non-matching characters with, when hostname rewrite is enabled.\n\nAny single ASCII character.\n\nDefaults to \"_\".",
			},

			// Indicates if client supplied hostnames will be rewritten prior to DDNS update by replacing every character that does not match _hostname_rewrite_regex_ by _hostname_rewrite_char_.
			//
			// Defaults to _false_.
			"hostname_rewrite_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates if client supplied hostnames will be rewritten prior to DDNS update by replacing every character that does not match _hostname_rewrite_regex_ by _hostname_rewrite_char_.\n\nDefaults to _false_.",
			},

			// The regex bracket expression to match valid characters.
			//
			// Must begin with "[" and end with "]" and be a compilable POSIX regex.
			//
			// Defaults to "[^a-zA-Z0-9_.]".
			"hostname_rewrite_regex": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The regex bracket expression to match valid characters.\n\nMust begin with \"[\" and end with \"]\" and be a compilable POSIX regex.\n\nDefaults to \"[^a-zA-Z0-9_.]\".",
			},

			// The inheritance configuration.
			"inheritance_sources": {
				Type:        schema.TypeList,
				Elem:        dataSourceIpamsvcIPSpaceInheritance(),
				MaxItems:    1,
				Optional:    true,
				Description: "The inheritance configuration.",
			},

			// The name of the IP space. Must contain 1 to 256 characters. Can include UTF-8.
			// Required: true
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the IP space. Must contain 1 to 256 characters. Can include UTF-8.",
			},

			// The tags for the IP space in JSON format.
			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "The tags for the IP space in JSON format.",
			},

			// The utilization threshold settings for the IP space.
			// Read Only: true
			//"threshold": {
			//	Type:        schema.TypeList,
			//	Elem:        dataSourceIpamsvcUtilizationThreshold(),
			//	Computed:    true,
			//	Description: "The utilization threshold settings for the IP space.",
			//},

			// Time when the object has been updated. Equals to _created_at_ if not updated after creation.
			// Read Only: true
			// Format: date-time
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Time when the object has been updated. Equals to _created_at_ if not updated after creation.",
			},

			// The utilization of IP addresses in the IP space.
			// Read Only: true
			//"utilization": {
			//	Type:        schema.TypeList,
			//	Elem:        dataSourceIpamsvcUtilization(),
			//	Computed:    true,
			//	Description: "The utilization of IP addresses in the IP space.",
			//},

			// The resource identifier.
			"vendor_specific_option_option_space": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The resource identifier.",
			},
		},
	}
}

func resourceIpamsvcIPSpaceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ipamsvc.IPAddressManagementAPI)

	var diags diag.Diagnostics

	name := d.Get("name").(string)
	s := &b1models.IpamsvcIPSpace{
		Name:    &name,
		Comment: d.Get("comment").(string),
	}

	resp, err := c.IPSpace.IPSpaceCreate(&ip_space.IPSpaceCreateParams{Body: s, Context: ctx}, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.Payload.Result.ID)

	resourceIpamsvcIPSpaceRead(ctx, d, m)

	return diags
}

func resourceIpamsvcIPSpaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ipamsvc.IPAddressManagementAPI)

	var diags diag.Diagnostics

	s, err := c.IPSpace.IPSpaceRead(
		&ip_space.IPSpaceReadParams{
			ID:      d.Id(),
			Context: ctx,
		},
		nil,
	)
	if err != nil {
		return diag.FromErr(err)
	}

	err = d.Set("asm_scope_flag", s.Payload.Result.AsmScopeFlag)
	if err != nil {
		diag.FromErr(err)
	}

	err = d.Set("name", s.Payload.Result.Name)
	if err != nil {
		return diag.FromErr(err)
	}

	err = d.Set("comment", s.Payload.Result.Comment)
	if err != nil {
		return diag.FromErr(err)
	}

	err = d.Set("created_at", s.Payload.Result.CreatedAt.String())
	if err != nil {
		return diag.FromErr(err)
	}

	err = d.Set("updated_at", s.Payload.Result.UpdatedAt.String())
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceIpamsvcIPSpaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ipamsvc.IPAddressManagementAPI)

	var diags diag.Diagnostics

	if d.HasChange("comment") {
		name := d.Get("name").(string)
		instance := b1models.IpamsvcIPSpace{
			Name:    &name,
			Comment: d.Get("comment").(string),
		}

		resp, err := c.IPSpace.IPSpaceUpdate(&ip_space.IPSpaceUpdateParams{ID: d.Id(), Body: &instance, Context: ctx}, nil)
		if err != nil {
			diag.FromErr(err)
		}

		d.SetId(resp.Payload.Result.ID)
	}

	return diags
}

func resourceIpamsvcIPSpaceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ipamsvc.IPAddressManagementAPI)

	var diags diag.Diagnostics

	ipSpaceID := d.Id()

	_, err := c.IPSpace.IPSpaceDelete(&ip_space.IPSpaceDeleteParams{ID: ipSpaceID, Context: ctx}, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}

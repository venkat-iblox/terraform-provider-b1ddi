package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// IpamsvcIPSpaceInheritance IPSpaceInheritance
//
// The __IPSpaceInheritance__ object specifies how and which fields _IPSpace_ object inherits from the parent.
//
// swagger:model ipamsvcIPSpaceInheritance
func schemaIpamsvcIPSpaceInheritance() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// The inheritance configuration for _asm_config_ field.
			"asm_config": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcInheritedASMConfig(),
				MaxItems:    1,
				Optional:    true,
				Description: "The inheritance configuration for _asm_config_ field.",
			},

			// The inheritance configuration for _ddns_client_update_ field from _IPSpace_ object.
			"ddns_client_update": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedString(),
				MaxItems:    1,
				Optional:    true,
				Description: "The inheritance configuration for _ddns_client_update_ field from _IPSpace_ object.",
			},

			// The inheritance configuration for _ddns_enabled_ field. Only action allowed is 'inherit'.
			"ddns_enabled": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedBool(),
				MaxItems:    1,
				Optional:    true,
				Description: "The inheritance configuration for _ddns_enabled_ field. Only action allowed is 'inherit'.",
			},

			// The inheritance configuration for _ddns_generate_name_ and _ddns_generated_prefix_ fields from _IPSpace_ object.
			"ddns_hostname_block": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcInheritedDDNSHostnameBlock(),
				MaxItems:    1,
				Optional:    true,
				Description: "The inheritance configuration for _ddns_generate_name_ and _ddns_generated_prefix_ fields from _IPSpace_ object.",
			},

			// The inheritance configuration for _ddns_send_updates_ and _ddns_domain_ fields from _IPSpace_ object.
			"ddns_update_block": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcInheritedDDNSUpdateBlock(),
				MaxItems:    1,
				Optional:    true,
				Description: "The inheritance configuration for _ddns_send_updates_ and _ddns_domain_ fields from _IPSpace_ object.",
			},

			// The inheritance configuration for _ddns_update_on_renew_ field from _IPSpace_ object.
			"ddns_update_on_renew": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedBool(),
				MaxItems:    1,
				Optional:    true,
				Description: "The inheritance configuration for _ddns_update_on_renew_ field from _IPSpace_ object.",
			},

			// The inheritance configuration for _ddns_use_conflict_resolution_ field from _IPSpace_ object.
			"ddns_use_conflict_resolution": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedBool(),
				MaxItems:    1,
				Optional:    true,
				Description: "The inheritance configuration for _ddns_use_conflict_resolution_ field from _IPSpace_ object.",
			},

			// The inheritance configuration for _dhcp_config_ field.
			"dhcp_config": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcInheritedDHCPConfig(),
				MaxItems:    1,
				Optional:    true,
				Description: "The inheritance configuration for _dhcp_config_ field.",
			},

			// The inheritance configuration for _dhcp_options_ field.
			"dhcp_options": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcInheritedDHCPOptionList(),
				MaxItems:    1,
				Optional:    true,
				Description: "The inheritance configuration for _dhcp_options_ field.",
			},

			// The inheritance configuration for _header_option_filename_ field.
			"header_option_filename": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedString(),
				MaxItems:    1,
				Optional:    true,
				Description: "The inheritance configuration for _header_option_filename_ field.",
			},

			// The inheritance configuration for _header_option_server_address_ field.
			"header_option_server_address": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedString(),
				MaxItems:    1,
				Optional:    true,
				Description: "The inheritance configuration for _header_option_server_address_ field.",
			},

			// The inheritance configuration for _header_option_server_name_ field.
			"header_option_server_name": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedString(),
				MaxItems:    1,
				Optional:    true,
				Description: "The inheritance configuration for _header_option_server_name_ field.",
			},

			// The inheritance configuration for _hostname_rewrite_enabled_, _hostname_rewrite_regex_, and _hostname_rewrite_char_ fields from _IPSpace_ object.
			"hostname_rewrite_block": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcInheritedHostnameRewriteBlock(),
				MaxItems:    1,
				Optional:    true,
				Description: "The inheritance configuration for _hostname_rewrite_enabled_, _hostname_rewrite_regex_, and _hostname_rewrite_char_ fields from _IPSpace_ object.",
			},

			// The inheritance configuration for _vendor_specific_option_option_space_ field.
			"vendor_specific_option_option_space": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedIdentifier(),
				MaxItems:    1,
				Optional:    true,
				Description: "The inheritance configuration for _vendor_specific_option_option_space_ field.",
			},
		},
	}
}

func flattenIpamsvcIPSpaceInheritance(r *models.IpamsvcIPSpaceInheritance) []interface{} {
	if r == nil {
		return []interface{}{}
	}

	res := make(map[string]interface{})

	res["asm_config"] = flattenIpamsvcInheritedASMConfig(r.AsmConfig)
	res["ddns_client_update"] = flattenInheritanceInheritedString(r.DdnsClientUpdate)
	res["ddns_enabled"] = flattenInheritanceInheritedBool(r.DdnsEnabled)
	res["ddns_hostname_block"] = flattenIpamsvcInheritedDDNSHostnameBlock(r.DdnsHostnameBlock)
	res["ddns_update_block"] = flattenIpamsvcInheritedDDNSUpdateBlock(r.DdnsUpdateBlock)
	res["ddns_update_on_renew"] = flattenInheritanceInheritedBool(r.DdnsUpdateOnRenew)
	res["ddns_use_conflict_resolution"] = flattenInheritanceInheritedBool(r.DdnsUseConflictResolution)
	res["dhcp_config"] = flattenIpamsvcInheritedDHCPConfig(r.DhcpConfig)
	res["dhcp_options"] = flattenIpamsvcInheritedDHCPOptionList(r.DhcpOptions)
	res["header_option_filename"] = flattenInheritanceInheritedString(r.HeaderOptionFilename)
	res["header_option_server_address"] = flattenInheritanceInheritedString(r.HeaderOptionServerAddress)
	res["header_option_server_name"] = flattenInheritanceInheritedString(r.HeaderOptionServerName)
	res["hostname_rewrite_block"] = flattenIpamsvcInheritedHostnameRewriteBlock(r.HostnameRewriteBlock)
	res["vendor_specific_option_option_space"] = flattenInheritanceInheritedIdentifier(r.VendorSpecificOptionOptionSpace)

	return []interface{}{res}
}

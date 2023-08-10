package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// IpamsvcInheritedDHCPOption InheritedDHCPOption
//
// The inheritance configuration for a field of type of _OptionItem_.
//
// swagger:model ipamsvcInheritedDHCPOption
func schemaIpamsvcInheritedDHCPOptionItem() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			// The resource identifier.
			"overriding_group": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The resource identifier.",
			},

			// The inherited value for the DHCP option.
			// Read Only: true
			"option": {
				Type:        schema.TypeSet,
				Elem:        schemaIpamsvcOptionItem(),
				Optional:    true,
				Description: "The inherited value for the DHCP option.",
			},
		},
	}
}

func flattenIpamsvcInheritedDHCPOptionItem(r *models.IpamsvcInheritedDHCPOptionItem) map[string]interface{} {
	if r == nil {
		return nil
	}

	return map[string]interface{}{
		"overriding_group": r.OverridingGroup,
		"option":           flattenIpamsvcOptionItem(r.Option),
	}
}

func expandIpamsvcInheritedDHCPOptionItem(d map[string]interface{}) *models.IpamsvcInheritedDHCPOptionItem {
	if len(d) == 0 || d == nil {
		return nil
	}

	return &models.IpamsvcInheritedDHCPOptionItem{
		OverridingGroup: d["overriding_group"].(string),
		Option:          expandIpamsvcOptionItem(d["option"].(map[string]interface{})),
	}
}

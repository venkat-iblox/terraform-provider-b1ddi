package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// IpamsvcOptionItem OptionItem
//
// An item (_dhcp/option_item_) in a list of DHCP options. May be either a specific option or a group of options.
//
// swagger:model ipamsvcOptionItem
func dataSourceIpamsvcOptionItem() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// The resource identifier.
			"group": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The resource identifier.",
			},

			// The resource identifier.
			"option_code": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The resource identifier.",
			},

			// The option value.
			"option_value": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The option value.",
			},

			// The type of item.
			//
			// Valid values are:
			// * _group_
			// * _option_
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The type of item.\n\nValid values are:\n* _group_\n* _option_",
			},
		},
	}
}

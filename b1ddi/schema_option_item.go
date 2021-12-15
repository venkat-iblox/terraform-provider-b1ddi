package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// IpamsvcOptionItem OptionItem
//
// An item (_dhcp/option_item_) in a list of DHCP options. May be either a specific option or a group of options.
//
// swagger:model ipamsvcOptionItem
func schemaIpamsvcOptionItem() *schema.Resource {
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

func flattenIpamsvcOptionItem(r *models.IpamsvcOptionItem) []interface{} {
	if r == nil {
		return []interface{}{}
	}

	res := make(map[string]interface{})

	res["group"] = r.Group
	res["option_code"] = r.OptionCode
	res["option_value"] = r.OptionValue
	res["type"] = r.Type

	return []interface{}{res}
}

func expandIpamsvcOptionItem(d []interface{}) *models.IpamsvcOptionItem {
	if len(d) == 0 || d[0] == nil {
		return nil
	}
	in := d[0].(map[string]interface{})

	return &models.IpamsvcOptionItem{
		Group:       in["group"].(string),
		OptionCode:  in["option_code"].(string),
		OptionValue: in["option_value"].(string),
		Type:        in["type"].(string),
	}
}

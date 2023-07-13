package b1ddi

import (
	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/infobloxopen/b1ddi-go-client/models"
)

func schemaIpamsvcHostAddress() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// Field usage depends on the operation:
			//  * For read operation, _address_ of the _Address_ corresponding to the _ref_ resource.
			//  * For write operation, _address_ to be created if the _Address_ does not exist. Required if _ref_ is not set on write:
			//     * If the _Address_ already exists and is already pointing to the right _Host_, the operation proceeds.
			//     * If the _Address_ already exists and is pointing to a different _Host, the operation must abort.
			//     * If the _Address_ already exists and is not pointing to any _Host_, it is linked to the _Host_.
			// Required: true
			"address": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Field usage depends on the operation:\n * For read operation, _address_ of the _Address_ corresponding to the _ref_ resource.\n * For write operation, _address_ to be created if the _Address_ does not exist. Required if _ref_ is not set on write:\n    * If the _Address_ already exists and is already pointing to the right _Host_, the operation proceeds.\n    * If the _Address_ already exists and is pointing to a different _Host, the operation must abort.\n    * If the _Address_ already exists and is not pointing to any _Host_, it is linked to the _Host_.",
			},

			// The resource identifier.
			// Required: true
			"ref": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The resource identifier.",
			},

			// The resource identifier.
			// Required: true
			"space": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The resource identifier.",
			},
		},
	}
}

func expandIpamHostAddress(d map[string]interface{}) *models.IpamsvcHostAddress {
	if d == nil || len(d) == 0 {
		return nil
	}

	return &models.IpamsvcHostAddress{
		Address: swag.String(d["address"].(string)),
		Ref:     swag.String(d["ref"].(string)),
		Space:   swag.String(d["space"].(string)),
	}
}

func flattenIpamHostAddress(r []*models.IpamsvcHostAddress) []interface{} {
	if r == nil {
		return nil
	}
	var ha []interface{}

	for _, address := range r {
		ha = append(ha, map[string]interface{}{
			"address": address.Address,
			"ref":     address.Ref,
			"space":   address.Space,
		})
	}
	return ha
}

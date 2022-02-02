// Code generated by go-swagger; DO NOT EDIT.

package b1ddi

import (
	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// IpamsvcName Name
//
// The __Name__ object represents a name associated with an address.
//
// swagger:model ipamsvcName
func schemaIpamsvcName() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// The name expressed as a single label or FQDN.
			// Required: true
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name expressed as a single label or FQDN.",
			},

			// The origin of the name.
			// Required: true
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The origin of the name.",
			},
		},
	}
}

func flattenIpamsvcName(r *models.IpamsvcName) map[string]interface{} {
	if r == nil {
		return nil
	}

	return map[string]interface{}{
		"name": r.Name,
		"type": r.Type,
	}
}

func expandIpamsvcName(d map[string]interface{}) *models.IpamsvcName {
	if d == nil || len(d) == 0 {
		return nil
	}

	return &models.IpamsvcName{
		Name: swag.String(d["name"].(string)),
		Type: swag.String(d["type"].(string)),
	}
}

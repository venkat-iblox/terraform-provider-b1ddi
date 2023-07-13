package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// ConfigWarning Warning
//
// Warning message related to the containing object.
//
// swagger:model configWarning
func schemaConfigWarning() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// Server IP address.
			// Required: true
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the warning",
			},

			// Server FQDN.
			// Required: true
			"message": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Warning message",
			},
		},
	}
}

func flattenConfigWarning(r []*models.ConfigWarning) []interface{} {
	if r == nil {
		return nil
	}

	var cw []interface{}

	for _, warning := range r {
		cw = append(cw, map[string]interface{}{
			"name":    warning.Name,
			"message": warning.Message,
		})
	}
	return cw
}

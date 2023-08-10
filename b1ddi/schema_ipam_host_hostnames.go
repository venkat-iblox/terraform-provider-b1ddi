package b1ddi

import (
	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/infobloxopen/b1ddi-go-client/models"
)

func schemaIpamsvcHostNames() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// When _true_, the name is treated as an alias.
			"alias": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "When _true_, the name is treated as an alias.",
			},

			// A name for the host.
			// Required: true
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A name for the host.",
			},

			// When _true_, the name field is treated as primary name.
			// There must be one and only one primary name in the list of host names.
			// The primary name will be treated as the canonical name for all the aliases.
			// PTR record will be generated only for the primary name.
			"primary_name": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "When _true_, the name field is treated as primary name.\nThere must be one and only one primary name in the list of host names.\nThe primary name will be treated as the canonical name for all the aliases.\nPTR record will be generated only for the primary name.",
			},

			// The resource identifier.
			// Required: true
			"zone": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The resource identifier.",
			},
		},
	}
}

func expandIpamHostName(d map[string]interface{}) *models.IpamsvcHostName {
	if d == nil || len(d) == 0 {
		return nil
	}

	return &models.IpamsvcHostName{
		Alias:       d["alias"].(bool),
		Name:        swag.String(d["name"].(string)),
		PrimaryName: d["primary_name"].(bool),
		Zone:        swag.String(d["zone"].(string)),
	}
}

func flattenIpamHostHostNames(r []*models.IpamsvcHostName) []interface{} {
	if r == nil {
		return nil
	}
	var ha []interface{}

	for _, names := range r {
		ha = append(ha, map[string]interface{}{
			"alias":        names.Alias,
			"name":         names.Name,
			"primary_name": names.PrimaryName,
			"zone":         names.Zone,
		})
	}
	return ha
}

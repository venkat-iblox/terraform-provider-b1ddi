package b1ddi

import (
	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/infobloxopen/b1ddi-go-client/models"
)

// DataRecordInheritance RecordInheritance
//
// The inheritance configuration specifies how the _Record_ object inherits the _ttl_ field.
//
// swagger:model dataRecordInheritance
func schemaHAGroupsHosts() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Description: "HA group host",
				Required:    true,
			},
			"role": {
				Type:        schema.TypeString,
				Description: "HA group host role(active or passive)",
				Required:    true,
			},
		},
	}
}

func expandHAGroupsHosts(d map[string]interface{}) *models.IpamsvcHAGroupHost {
	if d == nil || len(d) == 0 {
		return nil
	}

	return &models.IpamsvcHAGroupHost{
		Host: swag.String(d["host"].(string)),
		Role: d["role"].(string),
	}
}

func flattenHAGroupsHosts(r []*models.IpamsvcHAGroupHost) []interface{} {
	if r == nil {
		return []interface{}{}
	}
	var hosts []interface{}

	for _, host := range r {
		hosts = append(hosts, map[string]interface{}{
			"host": host.Host,
			"port": host.Port,
			"role": host.Role,
		})
	}
	return hosts
}

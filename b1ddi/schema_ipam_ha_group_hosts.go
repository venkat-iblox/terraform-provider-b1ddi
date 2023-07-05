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
			"address": {
				Type:        schema.TypeString,
				Description: "IP address of the HA group host",
				Computed:    true,
			},
			"state": {
				Type:        schema.TypeString,
				Description: "HA group state",
				Computed:    true,
			},
			"port": {
				Type:        schema.TypeInt,
				Description: "HA group host port",
				Computed:    true,
			},
			"heartbeats": {
				Type:        schema.TypeList,
				Description: "Heartbeat info between the hosts in HA Group",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"peer": {
							Type:        schema.TypeString,
							Description: "heartbeat from the peer",
							Computed:    true,
						},
						"successful_heartbeat": {
							Type:        schema.TypeString,
							Description: "timestamp of the last successful timestamp",
							Computed:    true,
						},
					}},
			},
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
			"address":    host.Address,
			"host":       host.Host,
			"port":       host.Port,
			"role":       host.Role,
			"state":      host.State,
			"heartbeats": flattenHAGroupsHeartbeats(host.Heartbeats),
		})
	}
	return hosts
}

func flattenHAGroupsHeartbeats(r []*models.IpamsvcHAGroupHeartbeats) []interface{} {
	if r == nil {
		return []interface{}{}
	}
	var hb []interface{}

	for _, heartbeats := range r {
		hb = append(hb, map[string]interface{}{
			"peer":                 heartbeats.Peer,
			"successful_heartbeat": heartbeats.SuccessfulHeartbeat,
		})
	}
	return hb
}

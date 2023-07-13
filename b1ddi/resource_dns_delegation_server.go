package b1ddi

import (
	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/infobloxopen/b1ddi-go-client/models"
)

// ConfigDelegationServer DelegationServer
//
// DNS zone delegation server.
//
// swagger:model configDelegationServer
func resourceDNSDelegationServer() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// Optional. IP Address of nameserver.
			//
			// Only required when fqdn of a delegation server falls under delegation fqdn
			"address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Optional. IP Address of nameserver.\n\nOnly required when fqdn of a delegation server falls under delegation fqdn",
			},

			// Required. FQDN of nameserver.
			// Required: true
			"fqdn": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Required. FQDN of nameserver.",
			},

			// FQDN of nameserver in punycode.
			// Read Only: true
			"protocol_fqdn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "FQDN of nameserver in punycode.",
			},
		},
	}
}

func expandDNSDelegationServer(d map[string]interface{}) *models.ConfigDelegationServer {
	if d == nil || len(d) == 0 {
		return nil
	}

	return &models.ConfigDelegationServer{
		Address:      d["address"].(string),
		Fqdn:         swag.String(d["fqdn"].(string)),
		ProtocolFqdn: d["protocol_fqdn"].(string),
	}
}

func flattenDNSDelegationServer(r []*models.ConfigDelegationServer) []interface{} {
	if r == nil {
		return nil
	}
	var ds []interface{}

	for _, server := range r {
		ds = append(ds, map[string]interface{}{
			"address":       server.Address,
			"fqdn":          server.Fqdn,
			"protocol_fqdn": server.ProtocolFqdn,
		})
	}

	return ds
}

package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// IpamsvcDHCPConfig DHCPConfig
//
// A DHCP Config object (_dhcp/dhcp_config_) represents a shared DHCP configuration that controls how leases are issued.
//
// swagger:model ipamsvcDHCPConfig
func dataSourceIpamsvcDHCPConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// Disable to allow leases only for known clients, those for which a fixed address is configured.
			"allow_unknown": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Disable to allow leases only for known clients, those for which a fixed address is configured.",
			},

			// The resource identifier.
			"filters": {
				Type:        schema.TypeList,
				Elem:        schema.TypeString,
				Optional:    true,
				Description: "The resource identifier.",
			},

			// The list of clients to ignore requests from.
			"ignore_list": {
				Type:        schema.TypeList,
				Elem:        dataSourceIpamsvcIgnoreItem(),
				Optional:    true,
				Description: "The list of clients to ignore requests from.",
			},

			// The lease duration in seconds.
			"lease_time": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The lease duration in seconds.",
			},
		},
	}
}

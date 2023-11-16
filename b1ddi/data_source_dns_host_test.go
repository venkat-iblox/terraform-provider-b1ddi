package b1ddi

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccDataSourceDnsHost(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					data "b1ddi_dns_hosts" "dns_hosts" {}
				`),
			},
		},
	})
}

func TestAccDataSourceDnsHostByName(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					data "b1ddi_dns_hosts" "dns_host" {
						filters = {
							"name" = "%s"
						}
					}
				`, testAccReadDnsHost(t)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.b1ddi_dns_hosts.dns_host", "results.#", "1"),
					resource.TestCheckResourceAttr("data.b1ddi_dns_hosts.dns_host", "results.0.name", testAccReadDnsHost(t)),
				),
			},
		},
	})
}

func TestAccDataSourceDnsHostByTags(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			//As there are no host with corresponding filters , we expect a null value for the result
			{
				Config: `
					data "b1ddi_dns_hosts" "dns_host" {
						tfilters = {
							TestType = "Acceptance"
						}
					}
						`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.b1ddi_dns_hosts.dns_host", "results.#", "0"),
				),
			},
		},
	})
}

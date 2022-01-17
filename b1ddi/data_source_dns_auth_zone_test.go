package b1ddi

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccDataSourceConfigAuthZone(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceDnsAuthZoneConfig(t),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccDnsViewExists("b1ddi_dns_view.tf_acc_test_dns_view"),
					testAccDnsAuthZoneExists("b1ddi_dns_auth_zone.tf_acc_test_auth_zone"),
				),
			},
			{
				Config: fmt.Sprintf(`
					data "b1ddi_dns_auth_zones" "tf_acc_auth_zones" {
						filters = {
							fqdn = "tf-acc-test.com."
						}
					}
				`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.b1ddi_dns_auth_zones.tf_acc_auth_zones", "results.#", "1"),
					resource.TestCheckResourceAttr("data.b1ddi_dns_auth_zones.tf_acc_auth_zones", "results.0.fqdn", "tf-acc-test.com."),
				),
			},
		},
	})
}

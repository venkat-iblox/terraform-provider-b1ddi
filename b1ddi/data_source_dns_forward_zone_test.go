package b1ddi

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceConfigForwardZone_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceDnsForwardZoneBasicTestStep(t),
			{
				Config: fmt.Sprintf(`
					data "b1ddi_dns_forward_zones" "tf_acc_forward_zones" {
						filters = {
							fqdn = "tf-acc-test.com."
						}
					}
				`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.b1ddi_dns_forward_zones.tf_acc_forward_zones", "results.#", "1"),
					resource.TestCheckResourceAttrSet("data.b1ddi_dns_forward_zones.tf_acc_forward_zones", "results.0.id"),
					resource.TestCheckResourceAttr("data.b1ddi_dns_forward_zones.tf_acc_forward_zones", "results.0.fqdn", "tf-acc-test.com."),
				),
			},
		},
	})
}

func TestAccDataSourceConfigForwardZone_FullConfigCloud(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceDnsForwardZoneFullConfigCloudTestStep(t),
			{
				Config: fmt.Sprintf(`
					data "b1ddi_dns_forward_zones" "tf_acc_forward_zones" {
						filters = {
							fqdn = "tf-acc-test.com."
						}
					}
				`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.b1ddi_dns_forward_zones.tf_acc_forward_zones", "results.#", "1"),
					resource.TestCheckResourceAttrSet("data.b1ddi_dns_forward_zones.tf_acc_forward_zones", "results.0.id"),
					resource.TestCheckResourceAttr("data.b1ddi_dns_forward_zones.tf_acc_forward_zones", "results.0.fqdn", "tf-acc-test.com."),
				),
			},
		},
	})
}

func TestAccDataSourceConfigForwardZone_FullConfigExternal(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceDnsForwardZoneFullConfigExternalTestStep(t),
			{
				Config: fmt.Sprintf(`
					data "b1ddi_dns_forward_zones" "tf_acc_forward_zones" {
						filters = {
							fqdn = "tf-acc-test.com."
						}
					}
				`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.b1ddi_dns_forward_zones.tf_acc_forward_zones", "results.#", "1"),
					resource.TestCheckResourceAttrSet("data.b1ddi_dns_forward_zones.tf_acc_forward_zones", "results.0.id"),
					resource.TestCheckResourceAttr("data.b1ddi_dns_forward_zones.tf_acc_forward_zones", "results.0.fqdn", "tf-acc-test.com."),
				),
			},
		},
	})
}

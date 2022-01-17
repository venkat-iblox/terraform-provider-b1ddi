package b1ddi

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccDataSourceConfigView(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					resource "b1ddi_dns_view" "tf_acc_test_dns_view" {
						name = "tf_acc_test_dns_view"
					}
				`),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccDnsViewExists("b1ddi_dns_view.tf_acc_test_dns_view"),
				),
			},
			{
				Config: fmt.Sprintf(`
					data "b1ddi_dns_views" "tf_acc_dns_views" {
						filters = {
							# Check string filter
							"name" = "tf_acc_test_dns_view"
						}
					}
				`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.b1ddi_dns_views.tf_acc_dns_views", "results.#", "1"),
					resource.TestCheckResourceAttr("data.b1ddi_dns_views.tf_acc_dns_views", "results.0.name", "tf_acc_test_dns_view"),
				),
			},
		},
	})
}

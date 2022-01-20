package b1ddi

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccDataSourceIpamsvcSubnet(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			// Create Subnet resource
			resourceSubnetBasicTestStep(),
			// Check Subnet data source
			{
				Config: fmt.Sprintf(`
					data "b1ddi_subnets" "tf_acc_subnets" {
						filters = {
							# Check string filter
							"name" = "tf_acc_test_subnet"
							"address" = "192.168.1.0"
							"cidr" = 24
						}
					}
				`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.b1ddi_subnets.tf_acc_subnets", "results.#", "1"),
					resource.TestCheckResourceAttr("data.b1ddi_subnets.tf_acc_subnets", "results.0.comment", "This Subnet is created by terraform provider acceptance test"),
				),
			},
		},
	})
}

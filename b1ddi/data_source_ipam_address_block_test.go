package b1ddi

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccDataSourceIpamsvcAddressBlock(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					resource "b1ddi_ip_space" "tf_acc_test_space" {
  						name = "tf_acc_test_space"
  						comment = "This IP Space is created by terraform provider acceptance test"
					}
					resource "b1ddi_address_block" "tf_acc_test_address_block" {
						address = "192.168.1.0"
						cidr = 24
  						name = "tf_acc_test_address_block"
						space = b1ddi_ip_space.tf_acc_test_space.id 
  						comment = "This Address Block is created by terraform provider acceptance test"
					}
				`),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccIPSpaceExists("b1ddi_ip_space.tf_acc_test_space"),
					testAccAddressBlockExists("b1ddi_address_block.tf_acc_test_address_block"),
				),
			},
			{
				Config: fmt.Sprintf(`
					data "b1ddi_address_blocks" "tf_acc_address_blocks" {
						filters = {
							# Check string filter
							"name" = "tf_acc_test_address_block"
						}
					}
				`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.b1ddi_address_blocks.tf_acc_address_blocks", "results.#", "1"),
					resource.TestCheckResourceAttr("data.b1ddi_address_blocks.tf_acc_address_blocks", "results.0.comment", "This Address Block is created by terraform provider acceptance test"),
				),
			},
		},
	})
}

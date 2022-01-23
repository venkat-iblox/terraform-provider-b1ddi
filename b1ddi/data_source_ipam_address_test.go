package b1ddi

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccDataSourceIpamsvcAddress(t *testing.T) {
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
					resource "b1ddi_subnet" "tf_acc_test_subnet" {
						name = "tf_acc_test_subnet"						
						address = "192.168.1.0"
						space = b1ddi_ip_space.tf_acc_test_space.id
						cidr = 24
  						comment = "This Subnet is created by terraform provider acceptance test"
					}
					resource "b1ddi_address" "tf_acc_test_address" {
						address = "192.168.1.15"
						comment = "This Address is created by terraform provider acceptance test"
						space = b1ddi_ip_space.tf_acc_test_space.id
						depends_on = [b1ddi_subnet.tf_acc_test_subnet]
					}
				`),
				Check: resource.ComposeAggregateTestCheckFunc(
					testCheckIPSpaceExists("b1ddi_ip_space.tf_acc_test_space"),
					testCheckSubnetExists("b1ddi_subnet.tf_acc_test_subnet"),
					testAccAddressExists("b1ddi_address.tf_acc_test_address"),
				),
			},
			{
				Config: fmt.Sprintf(`
					data "b1ddi_addresses" "tf_acc_addresses" {
						filters = {
							# Check string filter
							"address" = "192.168.1.15"
						}
					}
				`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.b1ddi_addresses.tf_acc_addresses", "results.#", "1"),
					resource.TestCheckResourceAttr("data.b1ddi_addresses.tf_acc_addresses", "results.0.comment", "This Address is created by terraform provider acceptance test"),
				),
			},
		},
	})
}

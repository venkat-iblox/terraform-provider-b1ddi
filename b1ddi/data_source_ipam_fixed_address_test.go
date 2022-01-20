package b1ddi

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccDataSourceIpamsvcFixedAddress(t *testing.T) {
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
						cidr = 24
						space = b1ddi_ip_space.tf_acc_test_space.id
  						comment = "This Subnet is created by terraform provider acceptance test"
					}
					resource "b1ddi_fixed_address" "tf_acc_test_fixed_address" {
						name = "tf_acc_test_fixed_address"						
						address = "192.168.1.15"
						ip_space = b1ddi_ip_space.tf_acc_test_space.id
						match_type = "mac"
						match_value = "00:00:00:00:00:00"
						comment = "This Fixed Address is created by terraform provider acceptance test"
						depends_on = [b1ddi_subnet.tf_acc_test_subnet]
					}
				`),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccIPSpaceExists("b1ddi_ip_space.tf_acc_test_space"),
					testCheckSubnetExists("b1ddi_subnet.tf_acc_test_subnet"),
					testAccFixedAddressExists("b1ddi_fixed_address.tf_acc_test_fixed_address"),
				),
			},
			{
				Config: fmt.Sprintf(`
					data "b1ddi_fixed_addresses" "tf_acc_fixed_addresses" {
						filters = {
							"name" = "tf_acc_test_fixed_address"
						}
					}
				`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.b1ddi_fixed_addresses.tf_acc_fixed_addresses", "results.#", "1"),
					resource.TestCheckResourceAttr("data.b1ddi_fixed_addresses.tf_acc_fixed_addresses", "results.0.comment", "This Fixed Address is created by terraform provider acceptance test"),
				),
			},
		},
	})
}

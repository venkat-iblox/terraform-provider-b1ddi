package b1ddi

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIpamNextAvailableSubnet_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			// Check Next Available Subnet data source
			{
				Config: `
					resource "b1ddi_ip_space" "tf_acc_test_space" {
  						name = "tf_acc_test_space"
  						comment = "This IP Space is created by terraform provider acceptance test"
					}
					resource "b1ddi_address_block" "tf_acc_test_address_block" {
						address = "192.168.1.0"
  						name = "tf_acc_test_address_block"
						cidr = 24
						space = b1ddi_ip_space.tf_acc_test_space.id 
  						comment = "This Address Block is created by terraform provider acceptance test"
						tags = {
							TestType = "Acceptance"
						}
					}
					data "b1ddi_ipam_next_available_subnets" "tf_subnet_nas" {
						id = b1ddi_address_block.tf_acc_test_address_block.id
						cidr = 26
						comment = "This is subnet that can be created using the next available subnet"
						name = "tf_subnet_nas"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.b1ddi_ipam_next_available_subnets.tf_subnet_nas", "results.#", "1"),
					resource.TestCheckResourceAttrSet("data.b1ddi_ipam_next_available_subnets.tf_subnet_nas", "results.0.id"),
					resource.TestCheckResourceAttr("data.b1ddi_ipam_next_available_subnets.tf_subnet_nas", "results.0.comment", "This is subnet that can be created using the next available subnet"),
					resource.TestCheckResourceAttr("data.b1ddi_ipam_next_available_subnets.tf_subnet_nas", "results.0.name", "tf_subnet_nas"),
					resource.TestCheckResourceAttr("data.b1ddi_ipam_next_available_subnets.tf_subnet_nas", "results.0.cidr", "26"),
					resource.TestCheckResourceAttrSet("data.b1ddi_ipam_next_available_subnets.tf_subnet_nas", "results.0.address"),
				),
			},
		},
	})
}

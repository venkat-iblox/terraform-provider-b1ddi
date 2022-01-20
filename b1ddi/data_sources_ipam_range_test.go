package b1ddi

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccDataSourceIpamsvcRange(t *testing.T) {
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
					resource "b1ddi_range" "tf_acc_test_range" {
						start = "192.168.1.15"
						end = "192.168.1.30"
  						name = "tf_acc_test_range"
						space = b1ddi_ip_space.tf_acc_test_space.id
  						comment = "This Range is created by terraform provider acceptance test"
						depends_on = [b1ddi_subnet.tf_acc_test_subnet]
					}
				`),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccIPSpaceExists("b1ddi_ip_space.tf_acc_test_space"),
					testCheckSubnetExists("b1ddi_subnet.tf_acc_test_subnet"),
					testAccRangeExists("b1ddi_range.tf_acc_test_range"),
				),
			},
			{
				Config: fmt.Sprintf(`
					data "b1ddi_ranges" "tf_acc_ranges" {
						filters = {
							# Check string filter
							"name" = "tf_acc_test_range"
							"start" = "192.168.1.15"
							"end" = "192.168.1.30"
						}
					}
				`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.b1ddi_ranges.tf_acc_ranges", "results.#", "1"),
					resource.TestCheckResourceAttr("data.b1ddi_ranges.tf_acc_ranges", "results.0.comment", "This Range is created by terraform provider acceptance test"),
				),
			},
		},
	})
}

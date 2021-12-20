package b1ddi

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/range_operations"
	"testing"
)

func TestAccResourceRange_basic(t *testing.T) {
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
  						comment = "This Address Block is created by terraform provider acceptance test"
						depends_on = [b1ddi_subnet.tf_acc_test_subnet]
					}`),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccRangeExists("b1ddi_range.tf_acc_test_range"),
					resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "start", "192.168.1.15"),
					resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "end", "192.168.1.30"),
					resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "name", "tf_acc_test_range"),
					resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "comment", "This Address Block is created by terraform provider acceptance test"),
				),
			},
		},
	})
}

func testAccRangeExists(resPath string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		res, found := s.RootModule().Resources[resPath]
		if !found {
			return fmt.Errorf("not found %s", resPath)
		}
		if res.Primary.ID == "" {
			return fmt.Errorf("ID for %s is not set", resPath)
		}

		cli := testAccProvider.Meta().(*ipamsvc.IPAddressManagementAPI)

		_, err := cli.RangeOperations.RangeRead(
			&range_operations.RangeReadParams{ID: res.Primary.ID, Context: context.TODO()},
			nil,
		)
		if err != nil {
			return err
		}
		return nil
	}
}

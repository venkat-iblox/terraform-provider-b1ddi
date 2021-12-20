package b1ddi

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/address_block"
	"testing"
)

func TestAccResourceAddressBlock_basic(t *testing.T) {
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
						address = "192.168.1.0/24"
  						name = "tf_acc_test_address_block"
						space = b1ddi_ip_space.tf_acc_test_space.id 
  						comment = "This Address Block is created by terraform provider acceptance test"
						tags = {
							TestType = "Acceptance"
						}
					}`),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccAddressBlockExists("b1ddi_address_block.tf_acc_test_address_block"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "address", "192.168.1.0"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "cidr", "24"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "name", "tf_acc_test_address_block"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "comment", "This Address Block is created by terraform provider acceptance test"),
				),
			},
		},
	})
}

func testAccAddressBlockExists(resPath string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		res, found := s.RootModule().Resources[resPath]
		if !found {
			return fmt.Errorf("not found %s", resPath)
		}
		if res.Primary.ID == "" {
			return fmt.Errorf("ID for %s is not set", resPath)
		}

		cli := testAccProvider.Meta().(*ipamsvc.IPAddressManagementAPI)

		_, err := cli.AddressBlock.AddressBlockRead(
			&address_block.AddressBlockReadParams{ID: res.Primary.ID, Context: context.TODO()},
			nil,
		)
		if err != nil {
			return err
		}
		return nil
	}
}

package b1ddi

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/fixed_address"
	"github.com/infobloxopen/b1ddi-go-client/models"
	"testing"
)

func TestAccResourceFixedAddress_basic(t *testing.T) {
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
					}`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.ComposeTestCheckFunc(
						testAccFixedAddressCompare(t, "b1ddi_fixed_address.tf_acc_test_fixed_address", models.IpamsvcFixedAddress{
							Name:    "tf_acc_test_fixed_address",
							Comment: "This Fixed Address is created by terraform provider acceptance test",
						}),
					),
					resource.TestCheckResourceAttrSet("b1ddi_fixed_address.tf_acc_test_fixed_address", "created_at"),
					resource.TestCheckResourceAttrSet("b1ddi_fixed_address.tf_acc_test_fixed_address", "updated_at"),
				),
			},
		},
	})
}

func testAccFixedAddressCompare(t *testing.T, resPath string, expected models.IpamsvcFixedAddress) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		res, found := s.RootModule().Resources[resPath]
		if !found {
			return fmt.Errorf("not found %s", resPath)
		}
		if res.Primary.ID == "" {
			return fmt.Errorf("ID for %s is not set", resPath)
		}

		cli := testAccProvider.Meta().(*ipamsvc.IPAddressManagementAPI)

		resp, err := cli.FixedAddress.FixedAddressRead(
			&fixed_address.FixedAddressReadParams{ID: res.Primary.ID, Context: context.TODO()},
			nil,
		)
		if err != nil {
			return err
		}

		if resp.Payload.Result.Name != expected.Name {
			return fmt.Errorf(
				"'name' does not match: \n got: '%s', \nexpected: '%s'",
				resp.Payload.Result.Name,
				expected.Name)
		}

		if resp.Payload.Result.Comment != expected.Comment {
			return fmt.Errorf(
				"'comment' does not match: \n got: '%s', \nexpected: '%s'",
				resp.Payload.Result.Comment,
				expected.Comment)
		}

		return nil
	}
}

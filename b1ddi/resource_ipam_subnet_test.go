package b1ddi

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/subnet"
	"github.com/infobloxopen/b1ddi-go-client/models"
	"testing"
)

func TestAccResourceSubnet_basic(t *testing.T) {
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
					}`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.ComposeTestCheckFunc(
						testAccSubnetCompare(t, "b1ddi_subnet.tf_acc_test_subnet", models.IpamsvcSubnet{
							Name:    "tf_acc_test_subnet",
							Comment: "This Subnet is created by terraform provider acceptance test",
						}),
					),
					resource.TestCheckResourceAttrSet("b1ddi_subnet.tf_acc_test_subnet", "protocol"),
					resource.TestCheckResourceAttrSet("b1ddi_subnet.tf_acc_test_subnet", "created_at"),
					resource.TestCheckResourceAttrSet("b1ddi_subnet.tf_acc_test_subnet", "updated_at"),
				),
			},
		},
	})
}

func testAccSubnetCompare(t *testing.T, resPath string, expected models.IpamsvcSubnet) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		res, found := s.RootModule().Resources[resPath]
		if !found {
			return fmt.Errorf("not found %s", resPath)
		}
		if res.Primary.ID == "" {
			return fmt.Errorf("ID for %s is not set", resPath)
		}

		cli := testAccProvider.Meta().(*ipamsvc.IPAddressManagementAPI)

		resp, err := cli.Subnet.SubnetRead(
			&subnet.SubnetReadParams{ID: res.Primary.ID, Context: context.TODO()},
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

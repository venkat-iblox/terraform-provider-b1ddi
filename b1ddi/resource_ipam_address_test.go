package b1ddi

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/address"
	"testing"
)

func TestAccResourceAddress_basic(t *testing.T) {
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
					resource "b1ddi_address" "tf_acc_test_address" {
						address = "192.168.1.15"
						comment = "This Address is created by terraform provider acceptance test"
						space = b1ddi_ip_space.tf_acc_test_space.id
					}`),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccAddressExists("b1ddi_address.tf_acc_test_address"),
					resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "address", "192.168.1.15"),
					resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "comment", "This Address is created by terraform provider acceptance test"),
					resource.TestCheckResourceAttrSet("b1ddi_address.tf_acc_test_address", "created_at"),
					// ToDo Add checks for dhcp_info
					resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "", ""),
				),
			},
		},
	})
}

func testAccAddressExists(resPath string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		res, found := s.RootModule().Resources[resPath]
		if !found {
			return fmt.Errorf("not found %s", resPath)
		}
		if res.Primary.ID == "" {
			return fmt.Errorf("ID for %s is not set", resPath)
		}

		cli := testAccProvider.Meta().(*b1ddiclient.Client)

		resp, err := cli.IPAddressManagementAPI.Address.AddressRead(
			&address.AddressReadParams{ID: res.Primary.ID, Context: context.TODO()},
			nil,
		)
		if err != nil {
			return err
		}
		if resp.Payload.Result.ID != res.Primary.ID {
			return fmt.Errorf(
				"'id' does not match: \n got: '%s', \nexpected: '%s'",
				resp.Payload.Result.ID,
				res.Primary.ID)
		}
		return nil
	}
}

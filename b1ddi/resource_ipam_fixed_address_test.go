package b1ddi

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/fixed_address"
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
					testAccFixedAddressExists("b1ddi_fixed_address.tf_acc_test_fixed_address"),
					resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "address", "192.168.1.15"),
					resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "comment", "This Fixed Address is created by terraform provider acceptance test"),
					resource.TestCheckResourceAttrSet("b1ddi_fixed_address.tf_acc_test_fixed_address", "created_at"),
					resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "dhcp_options.%", "0"),
					resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "header_option_filename", ""),
					resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "header_option_server_address", ""),
					resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "header_option_server_name", ""),
					// ToDo Check hostname
					resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "hostname", ""),
					resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "inheritance_assigned_hosts.%", "0"),
					resource.TestCheckResourceAttrSet("b1ddi_fixed_address.tf_acc_test_fixed_address", "inheritance_parent"),
					resource.TestCheckNoResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "inheritance_sources"),
					resource.TestCheckResourceAttrSet("b1ddi_fixed_address.tf_acc_test_fixed_address", "ip_space"),
					resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "match_type", "mac"),
					resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "match_value", "00:00:00:00:00:00"),
					resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "name", "tf_acc_test_fixed_address"),
					resource.TestCheckResourceAttrSet("b1ddi_fixed_address.tf_acc_test_fixed_address", "parent"),
					resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "tags.%", "0"),
					resource.TestCheckResourceAttrSet("b1ddi_fixed_address.tf_acc_test_fixed_address", "updated_at"),
				),
			},
		},
	})
}

func testAccFixedAddressExists(resPath string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		res, found := s.RootModule().Resources[resPath]
		if !found {
			return fmt.Errorf("not found %s", resPath)
		}
		if res.Primary.ID == "" {
			return fmt.Errorf("ID for %s is not set", resPath)
		}

		cli := testAccProvider.Meta().(*b1ddiclient.Client)

		resp, err := cli.IPAddressManagementAPI.FixedAddress.FixedAddressRead(
			&fixed_address.FixedAddressReadParams{ID: res.Primary.ID, Context: context.TODO()},
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

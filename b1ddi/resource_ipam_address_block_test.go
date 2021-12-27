package b1ddi

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
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

					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "asm_config.0.asm_threshold", "90"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "asm_config.0.enable", "true"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "asm_config.0.enable_notification", "true"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "asm_config.0.forecast_period", "14"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "asm_config.0.growth_factor", "20"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "asm_config.0.growth_type", "percent"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "asm_config.0.history", "30"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "asm_config.0.min_total", "10"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "asm_config.0.min_unused", "10"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "asm_config.0.reenable_date", "1970-01-01T00:00:00.000Z"),

					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "asm_scope_flag", "0"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "cidr", "24"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "comment", "This Address Block is created by terraform provider acceptance test"),
					resource.TestCheckResourceAttrSet("b1ddi_subnet.tf_acc_test_subnet", "created_at"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "ddns_client_update", "client"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "ddns_domain", ""),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "ddns_generate_name", "false"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "ddns_generated_prefix", "myhost"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "ddns_send_updates", "true"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "ddns_update_on_renew", "false"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "ddns_use_conflict_resolution", "true"),

					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "dhcp_config.0.allow_unknown", "true"),
					resource.TestCheckNoResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "dhcp_config.0.filters"),
					resource.TestCheckNoResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "dhcp_config.0.ignore_list"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "dhcp_config.0.lease_time", "3600"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "dhcp_host", ""),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "dhcp_options.%", "0"),

					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "dhcp_utilization.0.dhcp_free", "0"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "dhcp_utilization.0.dhcp_total", "0"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "dhcp_utilization.0.dhcp_used", "0"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "dhcp_utilization.0.dhcp_utilization", "0"),

					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "header_option_filename", ""),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "header_option_server_address", ""),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "header_option_server_name", ""),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "hostname_rewrite_char", "_"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "hostname_rewrite_enabled", "false"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "hostname_rewrite_regex", "[^a-zA-Z0-9_.]"),

					resource.TestCheckResourceAttrSet("b1ddi_subnet.tf_acc_test_subnet", "inheritance_parent"),
					resource.TestCheckNoResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "inheritance_sources"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "name", "tf_acc_test_address_block"),
					resource.TestCheckResourceAttrSet("b1ddi_subnet.tf_acc_test_subnet", "parent"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "protocol", "ip4"),
					resource.TestCheckResourceAttrSet("b1ddi_subnet.tf_acc_test_subnet", "space"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "tags.%", "0"),

					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "threshold.0.enabled", "false"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "threshold.0.high", "0"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "threshold.0.low", "0"),

					resource.TestCheckResourceAttrSet("b1ddi_subnet.tf_acc_test_subnet", "updated_at"),

					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "utilization.0.abandon_utilization", "0"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "utilization.0.abandoned", "0"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "utilization.0.dynamic", "0"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "utilization.0.free", "254"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "utilization.0.static", "2"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "utilization.0.total", "256"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "utilization.0.used", "2"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "utilization.0.utilization", "1"),
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

		cli := testAccProvider.Meta().(*b1ddiclient.Client)

		resp, err := cli.IPAddressManagementAPI.AddressBlock.AddressBlockRead(
			&address_block.AddressBlockReadParams{ID: res.Primary.ID, Context: context.TODO()},
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

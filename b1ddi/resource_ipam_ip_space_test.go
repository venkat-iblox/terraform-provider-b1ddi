package b1ddi

import (
	"context"
	"fmt"
	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/ip_space"
	"github.com/infobloxopen/b1ddi-go-client/models"
	"reflect"
	"testing"
)

// ToDo add test case for IP Space with DHCP options
// ToDo add test case for IP Space with Inheritance Sources
// ToDo add check deleted

func TestAccResourceIPSpace_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					resource "b1ddi_ip_space" "tf_acc_test_space" {
  						name = "tf_acc_test_space"
  						comment = "This IP Space is created by terraform provider acceptance test"
						tags = {
							TestType = "Acceptance"
						}
					}`),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccIPSpaceExists(t, "b1ddi_ip_space.tf_acc_test_space", models.IpamsvcIPSpace{
						Name:    swag.String("tf_acc_test_space"),
						Comment: "This IP Space is created by terraform provider acceptance test",
						Tags:    map[string]interface{}{"TestType": "Acceptance"},
					}),
					// Check default values
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.asm_threshold", "90"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.enable", "true"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.enable_notification", "true"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.forecast_period", "14"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.growth_factor", "20"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.growth_type", "percent"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.history", "30"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.min_total", "10"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.min_unused", "10"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.reenable_date", "1970-01-01T00:00:00.000Z"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_scope_flag", "0"),

					resource.TestCheckResourceAttrSet("b1ddi_ip_space.tf_acc_test_space", "created_at"),

					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "ddns_client_update", "client"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "ddns_domain", ""),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "ddns_generate_name", "false"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "ddns_generated_prefix", "myhost"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "ddns_send_updates", "true"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "ddns_update_on_renew", "false"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "ddns_use_conflict_resolution", "true"),

					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "dhcp_config.0.allow_unknown", "true"),
					resource.TestCheckNoResourceAttr("b1ddi_ip_space.tf_acc_test_space", "dhcp_config.0.filters"),
					resource.TestCheckNoResourceAttr("b1ddi_ip_space.tf_acc_test_space", "dhcp_config.0.ignore_list"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "dhcp_config.0.lease_time", "3600"),
				),
			},
		},
	})
}

func TestAccResourceIPSpace_full_config(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				resource "b1ddi_ip_space" "tf_acc_test_space" {
					asm_config {
						asm_threshold = 80
						forecast_period = 9
						growth_type = "count"
						history = 50
						min_total = 20
						min_unused = 20
					}
					name = "tf_acc_test_space"
					comment = "This IP Space is created by terraform provider acceptance test"
					ddns_client_update = "ignore"
					ddns_domain = "domain"
					ddns_generate_name = true
					ddns_generated_prefix = "tf_acc_host"
					ddns_send_updates = false
					ddns_update_on_renew = true
					ddns_use_conflict_resolution = false

					dhcp_config {
						allow_unknown = false
						#filters = ["filter1"]
						lease_time = 1800
					}

					tags = {
						TestType = "Acceptance"
					}
				}`),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccIPSpaceExists(t, "b1ddi_ip_space.tf_acc_test_space", models.IpamsvcIPSpace{
						Name:    swag.String("tf_acc_test_space"),
						Comment: "This IP Space is created by terraform provider acceptance test",
						Tags:    map[string]interface{}{"TestType": "Acceptance"},
					}),
					// Check default values
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.asm_threshold", "80"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.enable", "true"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.enable_notification", "true"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.forecast_period", "9"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.growth_factor", "20"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.growth_type", "count"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.history", "50"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.min_total", "20"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.min_unused", "20"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.reenable_date", "1970-01-01T00:00:00.000Z"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_scope_flag", "0"),

					resource.TestCheckResourceAttrSet("b1ddi_ip_space.tf_acc_test_space", "created_at"),

					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "ddns_client_update", "ignore"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "ddns_domain", "domain"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "ddns_generate_name", "true"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "ddns_generated_prefix", "tf_acc_host"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "ddns_send_updates", "false"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "ddns_update_on_renew", "true"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "ddns_use_conflict_resolution", "false"),

					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "dhcp_config.0.allow_unknown", "false"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "dhcp_config.0.filters", "[]"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "dhcp_config.0.ignore_list", "[]"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "dhcp_config.0.lease_time", "1800"),

					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "dhcp_options.0.type", "group"),
				),
			},
		},
	})
}

func testAccIPSpaceExists(t *testing.T, resPath string, expectedSpace models.IpamsvcIPSpace) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		res, found := s.RootModule().Resources[resPath]
		if !found {
			return fmt.Errorf("not found %s", resPath)
		}
		if res.Primary.ID == "" {
			return fmt.Errorf("ID for %s is not set", resPath)
		}

		cli := testAccProvider.Meta().(*ipamsvc.IPAddressManagementAPI)

		resultSpace, err := cli.IPSpace.IPSpaceRead(
			&ip_space.IPSpaceReadParams{ID: res.Primary.ID, Context: context.TODO()},
			nil,
		)
		if err != nil {
			return err
		}

		if *resultSpace.Payload.Result.Name != *expectedSpace.Name {
			return fmt.Errorf(
				"'name' does not match: \n got: '%s', \nexpected: '%s'",
				*resultSpace.Payload.Result.Name,
				*expectedSpace.Name)
		}

		if resultSpace.Payload.Result.Comment != expectedSpace.Comment {
			return fmt.Errorf(
				"'comment' does not match: \n got: '%s', \nexpected: '%s'",
				resultSpace.Payload.Result.Comment,
				expectedSpace.Comment)
		}

		if !reflect.DeepEqual(resultSpace.Payload.Result.Tags, expectedSpace.Tags) {
			return fmt.Errorf(
				"'tags' does not match: \n got: '%s', \nexpected: '%s'",
				resultSpace.Payload.Result.Tags,
				expectedSpace.Tags,
			)
		}

		return nil
	}
}

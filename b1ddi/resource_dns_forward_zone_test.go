package b1ddi

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/dns_config/forward_zone"
)

func TestAccResourceDnsForwardZone_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceDnsForwardZoneBasicTestStep(t),
			{
				ResourceName:      "b1ddi_dns_forward_zone.tf_acc_test_forward_zone",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func resourceDnsForwardZoneBasicTestStep(t *testing.T) resource.TestStep {
	return resource.TestStep{
		Config: fmt.Sprintf(`
		resource "b1ddi_dns_view" "tf_acc_test_dns_view" {
			name = "tf_acc_test_dns_view"
		}
		resource "b1ddi_dns_forward_zone" "tf_acc_test_forward_zone" {
			fqdn = "tf-acc-test.com."
			view = b1ddi_dns_view.tf_acc_test_dns_view.id
		}`),
		Check: resource.ComposeAggregateTestCheckFunc(
			testAccDnsForwardZoneExists("b1ddi_dns_forward_zone.tf_acc_test_forward_zone"),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "comment", ""),
			resource.TestCheckResourceAttrSet("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "created_at"),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "disabled", "false"),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "external_forwarders.#", "0"),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "forwarders_only", "false"),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "fqdn", "tf-acc-test.com."),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "hosts.#", "0"),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "internal_forwarders.#", "0"),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "mapped_subnet", ""),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "mapping", "forward"),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "nsgs.#", "0"),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "parent", ""),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "protocol_fqdn", "tf-acc-test.com."),
			resource.TestCheckNoResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "tags"),
			resource.TestCheckResourceAttrSet("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "updated_at"),
			resource.TestCheckResourceAttrSet("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "view"),
		),
	}
}

func testAccDnsForwardZoneExists(resPath string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		res, found := s.RootModule().Resources[resPath]
		if !found {
			return fmt.Errorf("not found %s", resPath)
		}
		if res.Primary.ID == "" {
			return fmt.Errorf("ID for %s is not set", resPath)
		}

		cli := testAccProvider.Meta().(*b1ddiclient.Client)

		resp, err := cli.DNSConfigurationAPI.ForwardZone.ForwardZoneRead(
			&forward_zone.ForwardZoneReadParams{ID: res.Primary.ID, Context: context.TODO()},
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

func TestAccResourceDnsForwardZone_FullConfigCloud(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceDnsForwardZoneFullConfigCloudTestStep(t),
			{
				ResourceName:      "b1ddi_dns_forward_zone.tf_acc_test_forward_zone",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func resourceDnsForwardZoneFullConfigCloudTestStep(t *testing.T) resource.TestStep {
	return resource.TestStep{
		Config: fmt.Sprintf(`
					resource "b1ddi_dns_view" "tf_acc_test_dns_view" {
						name = "tf_acc_test_dns_view"
					}
					resource "b1ddi_dns_forward_zone" "tf_acc_test_forward_zone" {
						comment = "This Forward Zone is created by the terraform provider acceptance test"
						disabled = true

						external_forwarders {
							address = "192.168.1.50"
							fqdn = "tf_test_external_forwarder."
						}
						fqdn = "tf-acc-test.com."
						tags = {
							TestType = "Acceptance"
						}
						view = b1ddi_dns_view.tf_acc_test_dns_view.id
					}`,
		),
		Check: resource.ComposeAggregateTestCheckFunc(
			testAccDnsForwardZoneExists("b1ddi_dns_forward_zone.tf_acc_test_forward_zone"),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "comment", "This Forward Zone is created by the terraform provider acceptance test"),
			resource.TestCheckResourceAttrSet("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "created_at"),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "disabled", "true"),

			resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "external_forwarders.#", "1"),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "external_forwarders.0.address", "192.168.1.50"),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "external_forwarders.0.fqdn", "tf_test_external_forwarder."),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "external_forwarders.0.protocol_fqdn", "tf_test_external_forwarder."),

			resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "fqdn", "tf-acc-test.com."),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "mapped_subnet", ""),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "mapping", "forward"),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "nsgs.#", "0"),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "parent", ""),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "protocol_fqdn", "tf-acc-test.com."),
			resource.TestCheckNoResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "tags.#"),
			resource.TestCheckResourceAttrSet("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "updated_at"),
			resource.TestCheckResourceAttrSet("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "view"),
		),
	}
}

func TestAccResourceDnsForwardZone_FullConfigExternal(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceDnsForwardZoneFullConfigExternalTestStep(t),
			{
				ResourceName:      "b1ddi_dns_forward_zone.tf_acc_test_forward_zone",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func resourceDnsForwardZoneFullConfigExternalTestStep(t *testing.T) resource.TestStep {
	return resource.TestStep{
		Config: fmt.Sprintf(`
					resource "b1ddi_dns_view" "tf_acc_test_dns_view" {
						name = "tf_acc_test_dns_view"
					}
					resource "b1ddi_dns_forward_zone" "tf_acc_test_forward_zone" {
						comment = "This Forward Zone is created by the terraform provider acceptance test"
						disabled = true
						external_forwarders {
							address = "192.168.1.60"
							fqdn = "tf_test_external_forwarder."
						}
						fqdn = "tf-acc-test.com."
						tags = {
							TestType = "Acceptance"
						}
						view = b1ddi_dns_view.tf_acc_test_dns_view.id
						
					}`,
		),
		Check: resource.ComposeAggregateTestCheckFunc(
			testAccDnsForwardZoneExists("b1ddi_dns_forward_zone.tf_acc_test_forward_zone"),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "comment", "This Forward Zone is created by the terraform provider acceptance test"),
			resource.TestCheckResourceAttrSet("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "created_at"),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "disabled", "true"),

			resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "external_forwarders.#", "1"),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "external_forwarders.0.address", "192.168.1.60"),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "external_forwarders.0.fqdn", "tf_test_external_forwarder."),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "external_forwarders.0.protocol_fqdn", "tf_test_external_forwarder."),

			resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "fqdn", "tf-acc-test.com."),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "mapped_subnet", ""),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "mapping", "forward"),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "nsgs.#", "0"),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "parent", ""),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "protocol_fqdn", "tf-acc-test.com."),

			resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "tags.TestType", "Acceptance"),

			resource.TestCheckResourceAttrSet("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "updated_at"),

			resource.TestCheckResourceAttrSet("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "view"),
		),
	}
}

func TestAccResourceDnsForwardZone_UpdateFQDNExpectError(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceDnsForwardZoneBasicTestStep(t),
			{
				Config: fmt.Sprintf(`
					resource "b1ddi_dns_view" "tf_acc_test_dns_view" {
						name = "tf_acc_test_dns_view"
					}
					resource "b1ddi_dns_forward_zone" "tf_acc_test_forward_zone" {
						fqdn = "tf-acc-test2.com."
						view = b1ddi_dns_view.tf_acc_test_dns_view.id
					}`,
				),
				ExpectError: regexp.MustCompile("changing the value of 'fqdn' field is not allowed"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccDnsForwardZoneExists("b1ddi_dns_forward_zone.tf_acc_test_forward_zone"),
					resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "fqdn", "tf-acc-test.com."),
				),
			},
			{
				ResourceName:      "b1ddi_dns_forward_zone.tf_acc_test_forward_zone",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceDnsForwardZone_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceDnsForwardZoneBasicTestStep(t),
			{
				Config: fmt.Sprintf(`
					resource "b1ddi_dns_view" "tf_acc_test_dns_view" {
						name = "tf_acc_test_dns_view"
					}
					resource "b1ddi_dns_forward_zone" "tf_acc_test_forward_zone" {
						comment = "This Forward Zone is created by the terraform provider acceptance test"
						disabled = true
						external_forwarders {
							address = "192.168.1.50"
							fqdn = "tf_test_external_forwarder."
						}
						fqdn = "tf-acc-test.com."
						tags = {
							TestType = "Acceptance"
						}	view = b1ddi_dns_view.tf_acc_test_dns_view.id
					}`,
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccDnsForwardZoneExists("b1ddi_dns_forward_zone.tf_acc_test_forward_zone"),
					resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "comment", "This Forward Zone is created by the terraform provider acceptance test"),
					resource.TestCheckResourceAttrSet("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "created_at"),
					resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "disabled", "true"),
					resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "external_forwarders.#", "1"),
					resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "external_forwarders.0.address", "192.168.1.50"),
					resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "external_forwarders.0.fqdn", "tf_test_external_forwarder."),
					resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "external_forwarders.0.protocol_fqdn", "tf_test_external_forwarder."),
					resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "fqdn", "tf-acc-test.com."),
					resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "mapped_subnet", ""),
					resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "mapping", "forward"),
					resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "nsgs.#", "0"),
					resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "parent", ""),
					resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "protocol_fqdn", "tf-acc-test.com."),
					resource.TestCheckResourceAttr("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "tags.TestType", "Acceptance"),
					resource.TestCheckResourceAttrSet("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "updated_at"),
					resource.TestCheckResourceAttrSet("b1ddi_dns_forward_zone.tf_acc_test_forward_zone", "view"),
				),
			},
			{
				ResourceName:      "b1ddi_dns_forward_zone.tf_acc_test_forward_zone",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

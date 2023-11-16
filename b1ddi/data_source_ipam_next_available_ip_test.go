package b1ddi

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIpamsvcNaIP_AB_Count(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceSubnetNAIPBasicTestStep(),
			{
				Config: `
					data "b1ddi_address_blocks" "ab" {
						filters = {
							name = "tf_address_block"
						}
					}
					data "b1ddi_next_available_ip" "next_ip_ab" {
  						id = "/test/12345"
  						ip_count = 5
					}`,
				ExpectError: regexp.MustCompile("invalid resource ID specified"),
			},
			{
				Config: `
					data "b1ddi_address_blocks" "ab" {
						filters = {
							name = "tf_address_block"
						}
					}
					data "b1ddi_next_available_ip" "next_ip_ab" {
  						id = data.b1ddi_address_blocks.ab.results.0.id
  						ip_count = 5
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.b1ddi_next_available_ip.next_ip_ab", "results.#", "5"),
					resource.TestCheckResourceAttr("data.b1ddi_next_available_ip.next_ip_ab", "results.0.address", "192.168.0.1"),
					resource.TestCheckResourceAttr("data.b1ddi_next_available_ip.next_ip_ab", "results.1.address", "192.168.0.2"),
					resource.TestCheckResourceAttr("data.b1ddi_next_available_ip.next_ip_ab", "results.2.address", "192.168.0.3"),
					resource.TestCheckResourceAttr("data.b1ddi_next_available_ip.next_ip_ab", "results.3.address", "192.168.0.4"),
					resource.TestCheckResourceAttr("data.b1ddi_next_available_ip.next_ip_ab", "results.4.address", "192.168.0.5"),
				),
			},
		},
	})
}

func TestAccDataSourceIpamsvcNaIP_AB_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceSubnetNAIPBasicTestStep(),
			{
				Config: `
					data "b1ddi_address_blocks" "ab" {
						filters = {
							name = "tf_address_block"
						}
					}	
					data "b1ddi_next_available_ip" "next_ip_ab" {
  						id = data.b1ddi_address_blocks.ab.results.0.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.b1ddi_next_available_ip.next_ip_ab", "results.#", "1"),
					resource.TestCheckResourceAttr("data.b1ddi_next_available_ip.next_ip_ab", "results.0.address", "192.168.0.1"),
				),
			},
		},
	})
}

func TestAccDataSourceIpamsvcNaIP_Subnet_Count(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceSubnetNAIPBasicTestStep(),
			{
				Config: `
					data "b1ddi_subnets" "sub" {
						filters = {
							name = "tf_acc_test_subnet"
						}
					}
					data "b1ddi_next_available_ip" "next_ip_sub" {
  						id = data.b1ddi_subnets.sub.results.0.id
  						ip_count = 5
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.b1ddi_next_available_ip.next_ip_sub", "results.#", "5"),
					resource.TestCheckResourceAttr("data.b1ddi_next_available_ip.next_ip_sub", "results.0.address", "192.168.1.1"),
					resource.TestCheckResourceAttr("data.b1ddi_next_available_ip.next_ip_sub", "results.1.address", "192.168.1.2"),
					resource.TestCheckResourceAttr("data.b1ddi_next_available_ip.next_ip_sub", "results.2.address", "192.168.1.3"),
					resource.TestCheckResourceAttr("data.b1ddi_next_available_ip.next_ip_sub", "results.3.address", "192.168.1.4"),
					resource.TestCheckResourceAttr("data.b1ddi_next_available_ip.next_ip_sub", "results.4.address", "192.168.1.5"),
				),
			},
		},
	})
}

func TestAccDataSourceIpamsvcNaIP_Subnet_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceSubnetNAIPBasicTestStep(),
			{
				Config: `
					data "b1ddi_subnets" "sub" {
						filters = {
							name = "tf_acc_test_subnet"
						}
					}	
					data "b1ddi_next_available_ip" "next_ip_sub" {
  						id = data.b1ddi_subnets.sub.results.0.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.b1ddi_next_available_ip.next_ip_sub", "results.#", "1"),
					resource.TestCheckResourceAttr("data.b1ddi_next_available_ip.next_ip_sub", "results.0.address", "192.168.1.1"),
				),
			},
		},
	})
}

func TestAccDataSourceIpamsvcNaIP_Range_Count(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceSubnetNAIPBasicTestStep(),
			{
				Config: `
					data "b1ddi_address_blocks" "ab" {
						filters = {
							name = "tf_address_block"
						}
					}
					data "b1ddi_next_available_ip" "next_ip_ab" {
  						id = ""
  						ip_count = 5
					}`,
				ExpectError: regexp.MustCompile("invalid resource ID specified"),
			},
			{
				Config: `
					data "b1ddi_ranges" "rng" {
						filters = {
							name = "tf_acc_test_range"
						}
					}
					data "b1ddi_next_available_ip" "next_ip_rng" {
  						id = data.b1ddi_ranges.rng.results.0.id
  						ip_count = 5
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.b1ddi_next_available_ip.next_ip_rng", "results.#", "5"),
					resource.TestCheckResourceAttr("data.b1ddi_next_available_ip.next_ip_rng", "results.0.address", "192.168.1.15"),
					resource.TestCheckResourceAttr("data.b1ddi_next_available_ip.next_ip_rng", "results.1.address", "192.168.1.16"),
					resource.TestCheckResourceAttr("data.b1ddi_next_available_ip.next_ip_rng", "results.2.address", "192.168.1.17"),
					resource.TestCheckResourceAttr("data.b1ddi_next_available_ip.next_ip_rng", "results.3.address", "192.168.1.18"),
					resource.TestCheckResourceAttr("data.b1ddi_next_available_ip.next_ip_rng", "results.4.address", "192.168.1.19"),
				),
			},
		},
	})
}

func TestAccDataSourceIpamsvcNaIP_Range_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceSubnetNAIPBasicTestStep(),
			{
				Config: `
					data "b1ddi_ranges" "rng" {
						filters = {
							name = "tf_acc_test_range"
						}
					}
					data "b1ddi_next_available_ip" "next_ip_rng" {
  						id = data.b1ddi_ranges.rng.results.0.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.b1ddi_next_available_ip.next_ip_rng", "results.#", "1"),
					resource.TestCheckResourceAttr("data.b1ddi_next_available_ip.next_ip_rng", "results.0.address", "192.168.1.15"),
				),
			},
		},
	})
}

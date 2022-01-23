package b1ddi

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccDataSourceIpamsvcRange(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceRangeBasicTestStep(),
			{
				Config: fmt.Sprintf(`
					data "b1ddi_ranges" "tf_acc_ranges" {
						filters = {
							# Check string filter
							"name" = "tf_acc_test_range"
							"start" = "192.168.1.15"
							"end" = "192.168.1.30"
						}
					}
				`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.b1ddi_ranges.tf_acc_ranges", "results.#", "1"),
					resource.TestCheckResourceAttr("data.b1ddi_ranges.tf_acc_ranges", "results.0.comment", "This Range is created by terraform provider acceptance test"),
				),
			},
		},
	})
}

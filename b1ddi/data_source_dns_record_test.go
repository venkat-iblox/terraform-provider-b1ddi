package b1ddi

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccDataSourceDataRecord(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceDnsRecordConfig(t),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccDnsViewExists("b1ddi_dns_view.tf_acc_test_dns_view"),
					testAccDnsAuthZoneExists("b1ddi_dns_auth_zone.tf_acc_test_auth_zone"),
					testAccDnsRecordExists("b1ddi_dns_record.tf_acc_test_dns_record"),
				),
			},
			{
				Config: fmt.Sprintf(`
					data "b1ddi_dns_records" "tf_acc_dns_records" {
						filters = {
							# Check string filter
							"name_in_zone" = "tf_acc_test_a_record"
							"type" = "A"
						}
					}
				`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.b1ddi_dns_records.tf_acc_dns_records", "results.#", "1"),
					resource.TestCheckResourceAttr("data.b1ddi_dns_records.tf_acc_dns_records", "results.0.name_in_zone", "tf_acc_test_a_record"),
					resource.TestCheckResourceAttr("data.b1ddi_dns_records.tf_acc_dns_records", "results.0.type", "A"),
					resource.TestCheckResourceAttr("data.b1ddi_dns_records.tf_acc_dns_records", "results.0.rdata.address", "192.168.1.15"),
				),
			},
		},
	})
}

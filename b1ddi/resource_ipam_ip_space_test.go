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
	"testing"
)

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
					}`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.ComposeTestCheckFunc(
						testAccIPSpaceCompare(t, "b1ddi_ip_space.tf_acc_test_space", models.IpamsvcIPSpace{
							Name:    swag.String("tf_acc_test_space"),
							Comment: "This IP Space is created by terraform provider acceptance test",
						}),
					),
					resource.TestCheckResourceAttrSet("b1ddi_ip_space.tf_acc_test_space", "created_at"),
				),
			},
		},
	})
}

func testAccIPSpaceCompare(t *testing.T, resPath string, expectedSpace models.IpamsvcIPSpace) resource.TestCheckFunc {
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

		return nil
	}
}

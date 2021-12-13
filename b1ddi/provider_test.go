package b1ddi

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"os"
	"testing"
)

var testAccProvider *schema.Provider
var testAccProviderFactories map[string]func() (*schema.Provider, error)

func init() {
	testAccProvider = Provider()

	testAccProviderFactories = map[string]func() (*schema.Provider, error){
		"b1ddi": func() (*schema.Provider, error) {
			return Provider(), nil
		},
	}
}

func testAccPreCheck(t *testing.T) {
	if host := os.Getenv("B1DDI_HOST"); host == "" {
		t.Fatal("B1DDI_HOST must be set for acceptance tests")
	}

	if token := os.Getenv("B1DDI_TOKEN"); token == "" {
		t.Fatal("B1DDI_TOKEN must be set for acceptance tests")
	}

	err := testAccProvider.Configure(context.TODO(), terraform.NewResourceConfigRaw(nil))
	if err != nil {
		t.Fatal(err)
	}
}

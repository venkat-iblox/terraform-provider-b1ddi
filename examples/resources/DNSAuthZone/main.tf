terraform {
  required_providers {
    b1ddi = {
      version = "0.1"
      source  = "github.com/infobloxopen/b1ddi"
    }
  }
}

resource "b1ddi_dns_view" "tf_example_dns_view" {
  name = "example_tf_dns_view"
}

resource "b1ddi_dns_auth_zone" "tf_example_auth_zone" {
  internal_secondaries {
    host = "{{ On Prem Host Name }}"
  }
  fqdn = "tf-example.com."
  primary_type = "cloud"
  view = b1ddi_dns_view.tf_example_dns_view.id
}

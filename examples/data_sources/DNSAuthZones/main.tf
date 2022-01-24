terraform {
  required_providers {
    b1ddi = {
      version = "0.1"
      source  = "github.com/infobloxopen/b1ddi"
    }
  }
}

# Get DNS Auth Zone with the specified FQDN
data "b1ddi_dns_auth_zones" "example_auth_zone" {
  filters = {
    fqdn = "tf-example.com."
  }
}

# Get all DNS Auth Zones
data "b1ddi_dns_auth_zones" "all_dns_auth_zones" {}

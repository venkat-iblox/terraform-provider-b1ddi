terraform {
  required_providers {
    b1ddi = {
      source  = "infobloxopen/b1ddi"
    }
  }
}

# Get DNS Forward Zone with the specified FQDN
data "b1ddi_dns_forward_zones" "example_forward_zone" {
  filters = {
    fqdn = "tf-example.com."
  }
}

# Get all DNS Forward Zones
data "b1ddi_dns_forward_zones" "all_dns_forward_zones" {}

terraform {
  required_providers {
    b1ddi = {
      source  = "infobloxopen/b1ddi"
    }
  }
}

# Get DNS Auth NSG with the specified name
data "b1ddi_dns_auth_nsgs" "tf_example_auth_nsg" {
  filters = {
    name = "tf_example_auth_nsg"
  }
}

# Get all DNS Auth NSGs
data "b1ddi_dns_auth_nsgs" "all_dns_auth_nsgs" {}

# Get DNS Auth NSGs with the specified tags
data "b1ddi_dns_auth_nsgs" "all_dns_auth_nsgs_with_tags" {
  tfilters = {
    location = "site1"
  }
}

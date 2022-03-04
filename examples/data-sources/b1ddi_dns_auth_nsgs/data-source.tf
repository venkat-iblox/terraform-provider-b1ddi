terraform {
  required_providers {
    b1ddi = {
      version = "0.1"
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

terraform {
  required_providers {
    b1ddi = {
      version = "0.1"
      source  = "infobloxopen/b1ddi"
    }
  }
}

# Get DNS Forward NSG with the specified name
data "b1ddi_dns_forward_nsgs" "tf_example_forward_nsg" {
  filters = {
    name = "tf_example_forward_nsg"
  }
}

# Get all DNS Forward NSGs
data "b1ddi_dns_forward_nsgs" "all_dns_forward_nsgs" {}

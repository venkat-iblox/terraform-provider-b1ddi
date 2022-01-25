terraform {
  required_providers {
    b1ddi = {
      version = "0.1"
      source  = "github.com/infobloxopen/b1ddi"
    }
  }
}

# Select specific fixed address
data "b1ddi_fixed_addresses" "tf_example_fixed_address" {
  filters = {
    "address" = "192.168.1.15"
  }
}

# Get all Fixed Addresses
data "b1ddi_fixed_addresses" "all_fixed_addresses" {}

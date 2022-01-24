terraform {
  required_providers {
    b1ddi = {
      version = "0.1"
      source  = "github.com/infobloxopen/b1ddi"
    }
  }
}

# Get all Addresses
data "b1ddi_addresses" "all_addresses" {}



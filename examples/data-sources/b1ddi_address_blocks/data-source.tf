terraform {
  required_providers {
    b1ddi = {
      version = "0.1"
      source  = "infobloxopen/b1ddi"
    }
  }
}

# Get Address Block by subnet address
data "b1ddi_address_blocks" "address_block_by_addr" {
  filters = {
    "address" = "192.168.1.0"
    "cidr" = 24
  }
}

# Get all Address Blocks
data "b1ddi_address_blocks" "all_address_blocks" {}

# Get Address Block by tag
data "b1ddi_address_blocks" "address_block_by_tag" {
  tfilters = {
    location  = "site1"
  }
}
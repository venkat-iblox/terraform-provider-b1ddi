terraform {
  required_providers {
    b1ddi = {
      version = "0.1.4"
      source  = "infobloxopen/b1ddi"
    }
  }
}

# Get Address Block by subnet address
data "b1ddi_address_blocks" "address_block_by_addr" {
  filters = {
    "address" = "10.0.0.0"
    "cidr" = 8
  }
}

# Get all Address Blocks
#data "b1ddi_address_blocks" "all_address_blocks" {}

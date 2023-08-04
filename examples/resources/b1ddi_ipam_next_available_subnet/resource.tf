terraform {
  required_providers {
    b1ddi = {
      version = "0.1.4"
      source  = "infobloxopen/b1ddi"
    }
  }
}

data "b1ddi_address_blocks" "tf_example_ab" {
  filters = {
    name = "AB1"
  }
}

resource "b1ddi_ipam_next_available_subnet" "tf_example_nas" {
  address_block_id = data.b1ddi_address_blocks.tf_example_ab.results[0].id
}
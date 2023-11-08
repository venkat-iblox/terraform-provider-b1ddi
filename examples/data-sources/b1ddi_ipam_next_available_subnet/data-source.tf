terraform {
  required_providers {
    b1ddi = {
      version = "0.1.5"
      source  = "infobloxopen/b1ddi"
    }
  }
}

resource "b1ddi_ip_space" "space" {
  name = "tf_ip_space"
}

resource "b1ddi_address_block" "tf_example_address_block" {
  address    = "192.168.1.0"
  cidr       = 24
  name       = "tf_example_address_block"
  space      = b1ddi_ip_space.space.id
  depends_on = [b1ddi_ip_space.space]
}

// List the subnets available in the above address block
// subnet_count = number of subnets to be created, if not specified defaults to 1
// cidr = size of subnet
data "b1ddi_ipam_next_available_subnets" "example_tf_subs" {
  id = b1ddi_address_block.tf_example_address_block.id
  name = "tf_subnet"
  cidr = 29
  subnet_count = 5
}


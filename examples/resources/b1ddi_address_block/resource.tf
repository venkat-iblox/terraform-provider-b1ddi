terraform {
  required_providers {
    b1ddi = {
      version = "0.1"
      source  = "infobloxopen/b1ddi"
    }
  }
}

resource "b1ddi_ip_space" "tf_example_space" {
  name = "tf_example_space"
  tags = {
    location   = "site1"
  }
}

resource "b1ddi_address_block" "tf_example_address_block" {
  address = "192.168.1.0"
  cidr = 24
  name = "tf_example_address_block"
  space = b1ddi_ip_space.tf_example_space.id
}
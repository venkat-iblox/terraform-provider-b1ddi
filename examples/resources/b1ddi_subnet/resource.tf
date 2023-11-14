terraform {
  required_providers {
    b1ddi = {
      source = "infobloxopen/b1ddi"
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

// Create a subnet dynamically under the Address Block 'tf_example_address_block.'
// Next available subnet under an Address Block example
resource "b1ddi_subnet" "example_tf_subnet_1" {
  name    = "tf_subnet_dynamic_nas"
  parent  = b1ddi_address_block.tf_example_address_block.id
  cidr    = 28
  comment = "subnet created through Terraform using Address block ID"
  space   = b1ddi_ip_space.space.id
  tags = {
    "TestType" = "Acceptance"
    "Cluster"  = "stg-1"
  }
}

// Create a static subnet
resource "b1ddi_subnet" "example_tf_subnet" {
  name       = "tf_subnet"
  address    = "192.168.3.0"
  cidr       = 28
  space      = b1ddi_ip_space.space.id
  comment    = "subnet created through Terraform"
  depends_on = [b1ddi_ip_space.space]
}

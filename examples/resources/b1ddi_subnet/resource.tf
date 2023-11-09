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

// Create a subnet dynamically under the Address Block 'tf_example_address_block.'
// Next available subnet under an Address Block example
resource "b1ddi_subnet" "example_tf_subnet_1" {
  name = "tf_subnet_dynamic_nas"
  address = b1ddi_address_block.tf_example_address_block.id
  cidr = 28
  comment = "subnet created through Terraform using Address block ID"
  space = b1ddi_ip_space.space.id
}

// List the subnets available in the above address block
// subnet_count = number of subnets to be created
// cidr = size of subnet
data "b1ddi_subnets" "nas" {
  id           = b1ddi_address_block.tf_example_address_block.id
  cidr         = 27
  subnet_count = 3
  depends_on = [b1ddi_subnet.example_tf_subnet_1]
}

// Create the subnet dynamically listed from the above data source
resource "b1ddi_subnet" "example_tf_subnet_nas" {
  count   = data.b1ddi_subnets.nas.subnet_count
  name    = "tf_subnet-${count.index}"
  address = data.b1ddi_subnets.nas.results[count.index].address
  cidr    = data.b1ddi_subnets.nas.results[count.index].cidr
  space   = b1ddi_ip_space.space.id
  comment = "subnet created through Terraform using Next available subnet"
  lifecycle {
    ignore_changes = all
  }
  depends_on = [data.b1ddi_subnets.nas]
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

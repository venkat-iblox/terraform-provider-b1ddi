terraform {
  required_providers {
    b1ddi = {
      version = "0.1"
      source  = "infobloxopen/b1ddi"
    }
  }
}

resource "b1ddi_ip_space" "tf_example_space" {
  name    = "tf_example_space"
  comment = "This is the example IP Space created by the B1DDI terraform provider"
  tags = {
    location   = "site1"
  }
}

resource "b1ddi_subnet" "tf_example_subnet" {
  name    = "tf_example_subnet"
  address = "192.168.1.0"
  cidr    = 24
  space   = b1ddi_ip_space.tf_example_space.id
  tags = {
    location   = "site1"
  }
}

resource "b1ddi_fixed_address" "tf_example_fixed_address" {
  name        = "tf_example_fixed_address"
  address     = "192.168.1.15"
  ip_space    = b1ddi_ip_space.tf_example_space.id
  match_type  = "mac"
  match_value = "00:00:00:00:00:00"
  comment     = "This is the example Fixed Address created by the B1DDI terraform provider"
  tags = {
    location   = "site1"
  }
  depends_on  = [b1ddi_subnet.tf_example_subnet]
}
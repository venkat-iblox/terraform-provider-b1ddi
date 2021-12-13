terraform {
  required_providers {
    b1ddi = {
      version = "0.1"
      source  = "github.com/infobloxopen/b1ddi"
    }
  }
}

resource "b1ddi_ip_space" "example_tf_space" {
  name = "example_tf_space"
  comment = "Example IP space for the terraform provider"
}

resource "b1ddi_subnet" "example_tf_subnet" {
  name = "example_tf_subnet"
  space = b1ddi_ip_space.example_tf_space.id
  address = "192.168.1.0"
  cidr = 24
  comment = "Example Subnet for the terraform provider"
}

resource "b1ddi_fixed_address" "example_tf_fixed_address" {
  name = "example_tf_fixed_address"
  ip_space = b1ddi_ip_space.example_tf_space.id
  address = "192.168.1.15"
  match_type = "mac"
  match_value = "00:00:00:00:00:00"
  comment = "Example Fixed Address for the terraform provider"
  depends_on = [b1ddi_subnet.example_tf_subnet]
}

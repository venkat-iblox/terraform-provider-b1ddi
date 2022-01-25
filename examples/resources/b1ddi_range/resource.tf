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
  comment = "Example IP space created by the terraform provider"
}

resource "b1ddi_subnet" "example_tf_subnet" {
  name = "example_tf_subnet"
  space = b1ddi_ip_space.example_tf_space.id
  address = "192.168.1.0"
  cidr = 24
  comment = "Example Subnet created by the terraform provider"
}

resource "b1ddi_range" "tf_acc_test_range" {
  start = "192.168.1.15"
  end = "192.168.1.30"
  name = "example_tf_range"
  space = b1ddi_ip_space.example_tf_space.id
  comment = "Example Range created by the terraform provider"
  depends_on = [b1ddi_subnet.example_tf_subnet]
}
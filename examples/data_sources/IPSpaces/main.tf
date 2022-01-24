terraform {
  required_providers {
    b1ddi = {
      version = "0.1"
      source  = "github.com/infobloxopen/b1ddi"
    }
  }
}

# Select IP Space with specified name
data "b1ddi_ip_spaces" "example_tf_space" {
  filters = {
    "name" = "example_tf_space"
  }
}

# Get all IP Spaces
data "b1ddi_ip_spaces" "all_ip_spaces" {}

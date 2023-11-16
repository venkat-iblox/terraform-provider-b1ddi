terraform {
  required_providers {
    b1ddi = {
      version = "0.1"
      source  = "infobloxopen/b1ddi"
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

# Get IP Space by tag
data "b1ddi_ip_spaces" "ip_space_by_tag" {
  tfilters = {
    location  = "site1"
  }
}
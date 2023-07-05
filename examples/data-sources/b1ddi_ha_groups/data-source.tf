terraform {
  required_providers {
    b1ddi = {
      version = "0.1.4"
      source  = "infobloxopen/b1ddi"
    }
  }
}

# Select HA Group by name
data "b1ddi_ha_groups" "example_ha_group" {
  filters = {
    "name" = "example_tf_ha_group"
  }
}

# Get all HA Groups
data "b1ddi_ha_groups" "all_ha_groups" {}

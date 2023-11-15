terraform {
  required_providers {
    b1ddi = {
      version = "0.1"
      source  = "infobloxopen/b1ddi"
    }
  }
}

# Select Subnet with specified name
data "b1ddi_subnets" "example_tf_subnet" {
  filters = {
    "name" = "example_tf_subnet"
  }
}

# Get all Subnets
data "b1ddi_subnets" "all_subnets" {}

# Get Subnet by tag
data "b1ddi_subnets" "subnet_by_tag"{
  tfilter = {
    location = "site1"
  }
}
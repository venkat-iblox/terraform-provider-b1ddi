terraform {
  required_providers {
    b1ddi = {
      source  = "infobloxopen/b1ddi"
    }
  }
}

# Select DNS View by name
data "b1ddi_dns_views" "tf_example_dns_view" {
  filters = {
    "name" = "example_tf_dns_view"
  }
}

# Get all DNS Views
data "b1ddi_dns_views" "all_dns_views" {}

# Get all DNS Views with specific tags
data "b1ddi_dns_views" "all_dns_views_with_tags" {
  tfilters = {
    location = "site1"
  }
}

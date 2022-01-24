terraform {
  required_providers {
    b1ddi = {
      version = "0.1"
      source  = "github.com/infobloxopen/b1ddi"
    }
  }
}

# Select DNS View by name
data "b1ddi_dns_views" "example_dns_view" {
  filters = {
    "name" = "example_tf_dns_view"
  }
}

# Get all DNS Views
data "b1ddi_dns_views" "all_dns_views" {}

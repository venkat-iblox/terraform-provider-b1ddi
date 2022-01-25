terraform {
  required_providers {
    b1ddi = {
      version = "0.1"
      source  = "github.com/infobloxopen/b1ddi"
    }
  }
}

# Get all Ranges
data "b1ddi_ranges" "all_ranges" {}

## Get specific Range by start and end values
data "b1ddi_ranges" "range_by_start_end" {
  filters = {
    "start" = "192.168.1.15",
    "end" = "192.168.1.30"
  }
}

## Get specific Range by name
data "b1ddi_ranges" "range_by_name" {
  filters = {
    "name" = "example_tf_range"
  }
}

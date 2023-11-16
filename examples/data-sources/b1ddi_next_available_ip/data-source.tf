terraform {
  required_providers {
    b1ddi = {
      source  = "infobloxopen/b1ddi"
    }
  }
}

// This example uses name as a filter, you can use other filters here too
data "b1ddi_address_blocks" "abs" {
  filters = {
    name = "<name of address block>"
  }
}

data "b1ddi_subnets" "subs" {
  filters = {
    name = "<name of the subnet>"
  }
}

data "b1ddi_ranges" "rgs" {
  filters = {
    name = "<name of the range>"
  }
}

// 'ip_count' allows you to get the number of next available ips in the resource specified by 'id'
// If not defined, count would default to 1
data "b1ddi_next_available_ip" "next_ip_ab" {
  id = data.b1ddi_address_blocks.abs.results.0.id
  ip_count = 5
}

data "b1ddi_next_available_ip" "next_ip_ab_default_count" {
  id = data.b1ddi_address_blocks.abs.results.0.id
}

data "b1ddi_next_available_ip" "next_ip_sub" {
  id = data.b1ddi_subnets.subs.results.0.id
  ip_count = 5
}

data "b1ddi_next_available_ip" "next_ip_range" {
  id = data.b1ddi_ranges.rgs.results.0.id
  ip_count = 5
}

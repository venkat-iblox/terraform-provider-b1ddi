terraform {
  required_providers {
    b1ddi = {
      version = "0.1"
      source  = "infobloxopen/b1ddi"
    }
  }
}

# Get all DHCP hosts
data "b1ddi_dhcp_hosts" "all_hosts" {}

# Get DHCP Host by name
data "b1ddi_dhcp_hosts" "dns_host_by_name" {
  filters = {
    name = "dhcp_host_name"
  }
}

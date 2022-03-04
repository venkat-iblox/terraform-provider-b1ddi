terraform {
  required_providers {
    b1ddi = {
      version = "0.1"
      source  = "infobloxopen/b1ddi"
    }
  }
}

variable "internal_secondary_name" {
  type = string
  description = "Internal secondary DNS Host name for the DNS Auth NSG configuration"
}

# Get DNS Host by name
data "b1ddi_dns_hosts" "dns_host_by_name" {
  filters = {
    "name" = var.internal_secondary_name
  }
}

# Create DNS Auth NSG
resource "b1ddi_dns_auth_nsg" "tf_example_auth_nsg" {
  name = "tf_example_auth_nsg"
  internal_secondaries {
    host = data.b1ddi_dns_hosts.dns_host_by_name.results.0.id
  }
}
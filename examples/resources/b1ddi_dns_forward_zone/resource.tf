terraform {
  required_providers {
    b1ddi = {
      source  = "infobloxopen/b1ddi"
    }
  }
}

variable "dns_host_name" {
  type = string
  description = "DNS Host name for the DNS Forward Zone configuration"
}

# Get DNS Host by name
data "b1ddi_dns_hosts" "dns_host_by_name" {
  filters = {
    "name" = var.dns_host_name
  }
}

resource "b1ddi_dns_view" "tf_example_dns_view" {
  name = "example_tf_dns_view"
}

resource "b1ddi_dns_forward_zone" "tf_example_forward_zone" {
  hosts = [data.b1ddi_dns_hosts.dns_host_by_name.results.0.id]
  fqdn = "tf-example.com."
  view = b1ddi_dns_view.tf_example_dns_view.id
}

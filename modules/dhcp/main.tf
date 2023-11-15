terraform {
  required_providers {
    b1ddi = {
      version = "0.1.5"
      source = "infobloxopen/b1ddi"
    }
  }
}

provider "b1ddi" {
  /*host    = "env-2a.test.infoblox.com"
  api_key = "5cd715820635a05c9e2eb552057a428f745609df45e6ea42158e16a444cc9b62"*/
  host = "env-6.test.infoblox.com"
  api_key = "f0ef42a4eda2badb05d15a867cb4787fb0ce99050086f808376f7708392677ff"
}

resource "b1ddi_ip_space" "tf_space" {
  name    = var.ip_space_name
  comment = "Example IP space create by the terraform provider"
}

resource "b1ddi_address_block" "tf_address_block" {
  space      = b1ddi_ip_space.tf_space.id
  address    = var.address_block.address
  cidr       = var.address_block.cidr
  name       = var.address_block.name
  depends_on = [b1ddi_ip_space.tf_space]
}

data "b1ddi_dhcp_hosts" "dhcp_host_by_name" {
  filters = {
    name = var.dhcp_host
  }
}

/*resource "b1ddi_ipam_next_available_subnet" "tf_subnet" {
  ab_id        = b1ddi_address_block.tf_address_block.id
  cidr         = var.subnet.next_available.cidr
  name         = var.subnet.next_available.name
  subnet_count = var.subnet.next_available.count
  comment      = "Example Subnet created using Next Available Subnet by the terraform provider"
  dhcp_host    = data.b1ddi_dhcp_hosts.dhcp_host_by_name.results.0.id
  depends_on   = [b1ddi_address_block.tf_address_block]
}*/

resource "b1ddi_subnet" "tf_subnet" {
  address    = b1ddi_address_block.tf_address_block.id
  space      = b1ddi_ip_space.tf_space.id
  cidr       = var.subnet.next_available.cidr
  name       = var.subnet.next_available.name
  comment    = "Example Subnet created using Next Available Subnet by the terraform provider"
  dhcp_host  = data.b1ddi_dhcp_hosts.dhcp_host_by_name.results.0.id
  depends_on = [b1ddi_address_block.tf_address_block]
  lifecycle {
    ignore_changes = [
      address
    ]
  }
}

resource "b1ddi_subnet" "tf_subnet_1" {
  name       = var.subnet.static.name
  space      = b1ddi_ip_space.tf_space.id
  address    = var.subnet.static.address
  cidr       = var.subnet.static.cidr
  dhcp_host  = data.b1ddi_dhcp_hosts.dhcp_host_by_name.results.0.id
  depends_on = [b1ddi_address_block.tf_address_block]
}

// We will need the next available IP for ranges for this
resource "b1ddi_range" "tf_range" {
  start      = var.range.range_start
  end        = var.range.range_end
  name       = var.range.name
  space      = b1ddi_ip_space.tf_space.id
  depends_on = [b1ddi_subnet.tf_subnet_1]
  lifecycle {
    ignore_changes = [
      dhcp_host
    ]
  }
}

resource "b1ddi_fixed_address" "tf_fixed_address" {
  name        = var.fixed_address.name
  address     = var.fixed_address.address
  ip_space    = b1ddi_ip_space.tf_space.id
  match_type  = "mac"
  match_value = "00:00:00:00:00:00"
  depends_on  = [b1ddi_subnet.tf_subnet_1]
}

# DNS views/zone + record creation

resource "b1ddi_dns_view" "tf_example_dns_view" {
  name = "tf_dns_view"
}

data "b1ddi_dns_hosts" "dns_host_by_name" {
  filters = {
    name = var.dhcp_host
  }
}

// Primary Auth zone
resource "b1ddi_dns_auth_zone" "tf_example_auth_zone" {
  internal_secondaries {
    host = data.b1ddi_dns_hosts.dns_host_by_name.results.0.id
  }
  fqdn         = "tf-example.com."
  primary_type = "cloud"
  view         = b1ddi_dns_view.tf_example_dns_view.id
}

// Reverse mapping zone
resource "b1ddi_dns_auth_zone" "tf_ptr_auth_zone" {
  fqdn         = "3.168.192.in-addr.arpa."
  primary_type = "cloud"
  view         = b1ddi_dns_view.tf_example_dns_view.id
}

resource "b1ddi_dns_record" "tf_a_record_gateway" {
  zone         = b1ddi_dns_auth_zone.tf_example_auth_zone.id
  name_in_zone = "tf-demo-gateway"
  rdata = {
    "address" = "192.168.3.1"
  }
  type = "A"
  options = {
    create_ptr = false
  }
  lifecycle {
    ignore_changes = [
      options
    ]
  }
}

resource "b1ddi_dns_record" "tf_a_record" {
  zone         = b1ddi_dns_auth_zone.tf_example_auth_zone.id
  name_in_zone = "tf-demo"
  rdata = {
    "address" = "192.168.3.15"
  }
  type = "A"
  options = {
    create_ptr = true
  }
  lifecycle {
    ignore_changes = [
      options
    ]
  }
}

resource "b1ddi_dns_record" "tf_a_record_1" {
  zone         = b1ddi_dns_auth_zone.tf_example_auth_zone.id
  name_in_zone = "tf-demo-1"
  rdata = {
    "address" = "192.168.3.16"
  }
  type = "A"
  options = {
    create_ptr = true
  }
  lifecycle {
    ignore_changes = [
      options
    ]
  }
}
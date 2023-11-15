terraform {
  required_providers {
    b1ddi = {
      source = "infobloxopen/b1ddi"
    }

  }
}

provider "b1ddi" {
  host    = "env-6.test.infoblox.com"
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

resource "b1ddi_ipam_next_available_subnet" "tf_subnet" {
  ab_id        = b1ddi_address_block.tf_address_block.id
  cidr         = var.subnet.next_available.cidr
  name         = var.subnet.next_available.name
  subnet_count = var.subnet.next_available.count
  comment      = "Example Subnet created using Next Available Subnet by the terraform provider"
  dhcp_host    = data.b1ddi_dhcp_hosts.dhcp_host_by_name.results.0.id
  depends_on   = [b1ddi_address_block.tf_address_block]
}

resource "b1ddi_subnet" "tf_subnet" {
  name    = var.subnet.static.name
  space   = b1ddi_ip_space.tf_space.id
  address = var.subnet.static.address
  cidr    = var.subnet.static.cidr
  depends_on   = [b1ddi_address_block.tf_address_block]
}

// We will need the next available IP for ranges for this
resource "b1ddi_range" "tf_range" {
  start      = var.range.range_start
  end        = var.range.range_end
  name       = var.range.name
  space      = b1ddi_ip_space.tf_space.id
  depends_on = [b1ddi_subnet.tf_subnet]
}

resource "b1ddi_fixed_address" "tf_fixed_address" {
  name        = var.fixed_address.name
  address     = var.fixed_address.address
  ip_space    = b1ddi_ip_space.tf_space.id
  match_type  = "mac"
  match_value = "00:00:00:00:00:00"
  depends_on  = [b1ddi_subnet.tf_subnet]
}

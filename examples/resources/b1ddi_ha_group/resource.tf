terraform {
  required_providers {
    b1ddi = {
      version = "0.1.4"
      source  = "infobloxopen/b1ddi"
    }
  }
}

data "b1ddi_dhcp_hosts" "host1" {
  filters = {
    name = "Venkat-01"
  }
}

data "b1ddi_dhcp_hosts" "host2" {
  filters = {
    name = "ZTP_venkathost_6681338607527706568"
  }
}

data "b1ddi_dhcp_hosts" "host3" {
  filters = {
    name = "ZTP_venkathost_2348760024534794691"
  }
}

data "b1ddi_dhcp_hosts" "host4" {
  filters = {
    name = "Naveen-OnPrem-172-28-5-186"
  }
}

resource "b1ddi_ha_group" "tf_example_ha_group" {
  name    = "example_tf_ha_group"
  mode    = "active-active"
  hosts {
    host = data.b1ddi_dhcp_hosts.host1.results.0.id
    role = "active"
  }
  hosts {
    host = data.b1ddi_dhcp_hosts.host2.results.0.id
    role = "active"
  }
}

resource "b1ddi_ha_group" "tf_example_ha_group_1" {
  name    = "example_tf_ha_group_1"
  mode    = "advanced-active-passive"
  hosts {
    host = data.b1ddi_dhcp_hosts.host4.results.0.id
    role = "active"
  }
  hosts {
    host = data.b1ddi_dhcp_hosts.host3.results.0.id
    role = "passive"
  }
}


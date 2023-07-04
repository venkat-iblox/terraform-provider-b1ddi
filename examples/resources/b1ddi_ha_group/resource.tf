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
    name = "ZTP_venkathost_2348760024534794691"
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

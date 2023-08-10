terraform {
  required_providers {
    b1ddi = {
      version = "0.1.4"
      source  = "infobloxopen/b1ddi"
    }
  }
}
data "b1ddi_ip_spaces" "example_tf_space" {
  filters = {
    "name" = "tf_test_space"
  }
}

data "b1ddi_dns_auth_zones" "example_auth_zone" {
  filters = {
    fqdn = "tf_test_zone."
  }
}

# Create IPAM Host
resource "b1ddi_ipam_host" "tf_example_ipam_host" {
  name = "tf_ipam_host"
  comment = "IPAM host created through terraform provider"
  auto_generate_records = true
  host_names {
    name = "tf_ipam_host_1"
    zone = data.b1ddi_dns_auth_zones.example_auth_zone.id
  }
  addresses {
    address = "192.168.0.3"
    space = data.b1ddi_ip_spaces.example_tf_space.id
    ref = ""
  }
}

terraform {
  required_providers {
    b1ddi = {
      version = "0.1"
      source  = "infobloxopen/b1ddi"
    }
  }
}

resource "b1ddi_dns_view" "tf_example_dns_view" {
  name = "example_tf_dns_view"
}

resource "b1ddi_dns_auth_zone" "tf_example_auth_zone" {
  internal_secondaries {
    host = "<On Prem Host Name>"
  }
  fqdn = "tf-example.com."
  primary_type = "cloud"
  view = b1ddi_dns_view.tf_example_dns_view.id
}

resource "b1ddi_dns_record" "a_record" {
  zone = b1ddi_dns_auth_zone.tf_example_auth_zone.id
  name_in_zone = "tf_example_a_record"
  rdata = {
    "address" = "192.168.1.15"
  }
  type = "A"
}

resource "b1ddi_dns_record" "ptr_record" {
  zone = b1ddi_dns_auth_zone.tf_example_auth_zone.id
  name_in_zone = "192.168.1.15"
  rdata = {
    "dname" = "tf_example_ptr_record"
  }
  type = "PTR"
}

resource "b1ddi_dns_record" "cname_record" {
  zone = b1ddi_dns_auth_zone.tf_example_auth_zone.id
  name_in_zone = "tf_example_cname_record"
  rdata = {
    "cname" = "canonical"
  }
  type = "CNAME"
}

resource "b1ddi_dns_record" "ns_record" {
  zone = b1ddi_dns_auth_zone.tf_example_auth_zone.id
  rdata = {
    "dname" = "ns1.tf-example.com."
  }
  type = "NS"
}

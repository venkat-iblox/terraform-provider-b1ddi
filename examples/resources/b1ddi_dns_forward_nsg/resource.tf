terraform {
  required_providers {
    b1ddi = {
      version = "0.1"
      source  = "infobloxopen/b1ddi"
    }
  }
}

# Create DNS Forward NSG
resource "b1ddi_dns_forward_nsg" "tf_example_forward_nsg" {
  name = "tf_example_forward_nsg"
}
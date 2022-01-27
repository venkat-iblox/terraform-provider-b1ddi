# b1ddi_dns_auth_zones (Data Source)

## Example Usage

```terraform
terraform {
  required_providers {
    b1ddi = {
      version = "0.1"
      source  = "github.com/infobloxopen/b1ddi"
    }
  }
}

# Get DNS Auth Zone with the specified FQDN
data "b1ddi_dns_auth_zones" "example_auth_zone" {
  filters = {
    fqdn = "tf-example.com."
  }
}

# Get all DNS Auth Zones
data "b1ddi_dns_auth_zones" "all_dns_auth_zones" {}
```

## Schema

### Optional

- **filters** (Map of String) Configure a map of filters to be applied on the search result.
- **id** (String) The ID of this resource.

### Read-Only

- **results** (List of Object) List of DNS Auth Zones matching filters. The schema of each element is identical to the b1ddi_dns_auth_zone resource schema. (see [nested schema](../resources/dns_auth_zone.md))

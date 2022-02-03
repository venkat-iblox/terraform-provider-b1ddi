# b1ddi_addresses (Data Source)

## Example Usage

```terraform
terraform {
  required_providers {
    b1ddi = {
      version = "0.1"
      source  = "infobloxopen/b1ddi"
    }
  }
}

# Get all Addresses
data "b1ddi_addresses" "all_addresses" {}
```

## Schema

### Optional

- **filters** (Map of String) Configure a map of filters to be applied on the search result.
- **id** (String) The ID of this resource.

### Read-Only

- **results** (List of Object) List of Addresses matching filters. The schema of each element is identical to the b1ddi_address resource schema. (see [nested schema](../resources/address.md))

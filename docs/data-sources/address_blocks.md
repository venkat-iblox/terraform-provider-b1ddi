# b1ddi_address_blocks (Data Source)



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

# Get Address Block by subnet address
data "b1ddi_address_blocks" "address_block_by_addr" {
  filters = {
    "address" = "192.168.1.0"
    "cidr" = 24
  }
}

# Get all Address Blocks
data "b1ddi_address_blocks" "all_address_blocks" {}
```

## Schema

### Optional

- **filters** (Map of String) Configure a map of filters to be applied on the search result.
- **id** (String) The ID of this resource.

### Read-Only

- **results** (List of Object) List of Address Blocks matching filters. The schema of each element is identical to the b1ddi_address_block resource schema. (see [nested schema](../resources/address_block.md))

# b1ddi_dns_records (Data Source)

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

# Select all PTR DNS Records
data "b1ddi_dns_records" "ptr_dns_records" {
  filters = {
    "type" = "PTR"
  }
}

# Select all A DNS Records
data "b1ddi_dns_records" "a_dns_records" {
  filters = {
    "type" = "A"
  }
}

# Select DNS Record with specified name in zone
data "b1ddi_dns_records" "tf_example_a_dns_record" {
  filters = {
    "name_in_zone" = "tf_example_a_record"
  }
}

# Get all DNS Records
data "b1ddi_dns_records" "all_dns_records" {}
```

## Schema

### Optional

- **filters** (Map of String) Configure a map of filters to be applied on the search result.
- **id** (String) The ID of this resource.

### Read-Only

- **results** (List of Object) List of DNS Records matching filters. The schema of each element is identical to the b1ddi_dns_record resource schema. (see [nested schema](../resources/dns_record.md))

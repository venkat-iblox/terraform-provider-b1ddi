# BloxOne DDI Terraform Provider

Terraform provider for the BloxOne DDI that enables lifecycle management of BloxOne DDI resources.

## Example Usage

Follow the "[BloxOne DDI API Guide](https://docs.infoblox.com/display/BloxOneDDI/BloxOne+DDI+API+Guide)" 
to get the API Key. 

```terraform
provider "b1ddi" {
  host = "b1ddi.infoblox.com"
  api_key = "<BloxOne DDI API Key>"
}
```

You can also configure the credentials using environment variables as follows:
```shell
export B1DDI_HOST=<B1DDI HOSTNAME> or <B1DDI IP ADDR>
export B1DDI_API_KEY=<B1DDI API KEY>
``` 

## Provider Features

The provider plugin has BloxOne DDI resources represented as Terraform resources and data sources. The consolidated
list of supported resources and data sources is as follows:

### Resources

- DNS View
- DNS Record
- DNS Auth Zone
- IPAM IP Space
- IPAM Address Block
- IPAM Subnet
- IPAM Range
- IPAM Fixed Address
- IPAM Address

### Data Sources

- DNS Views
- DNS Records
- DNS Auth Zones
- IPAM IP Spaces
- IPAM Address Blocks
- IPAM Subnets
- IPAM Ranges
- IPAM Fixed Addresses
- IPAM Addresses

## Provider Limitations

- Utilization data can be outdated after the respective IPAM resource is created. If you need the latest utilization
  data for previously created IPAM resources, you can run `terraform refresh` command.

## Schema

### Required

- **host** (String) BloxOne DDI host URL.
- **api_key** (String, Sensitive) API token for authentication against the Infoblox BloxOne DDI platform.

### Optional

- **base_path** (String) The base path is to indicate the API version and the product name.


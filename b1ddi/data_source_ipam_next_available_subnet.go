package b1ddi

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/address_block"
)

func dataSourceIpamNextAvailableSubnet() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIpamNextAvailableSubnetList,
		Schema: map[string]*schema.Schema{
			"filters": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Configure a map of filters to be applied on the search result.",
			},
			"results": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        dataSourceSchemaFromResource(resourceIpamNextAvailableSubnet),
				Description: "List of IP Spaces matching filters. The schema of each element is identical to the b1ddi_ip_space resource schema.",
			},
		},
	}
}

func dataSourceIpamNextAvailableSubnetList(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	var diags diag.Diagnostics
	filtersMap := d.Get("filters").(map[string]interface{})

	resp, err := c.IPAddressManagementAPI.AddressBlock.AddressBlockListNextAvailableSubnet(
		&address_block.AddressBlockListNextAvailableSubnetParams{
			ID:      filtersMap["address_block_id"].(string),
			Context: ctx,
		}, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	err = d.Set("results", flattenIpamNextAvailableSubnet(resp.Payload))
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

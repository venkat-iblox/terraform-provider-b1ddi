package b1ddi

import (
	"context"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/address_block"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/range_operations"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/subnet"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

func dataSourceIpamsvcNaIP() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIpamsvcNaIPRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`^ipam\/(range|subnet|address_block)\/[0-9a-f-].*$`), "invalid resource ID specified"),
				Description:  "An application specific resource identity of a resource",
			},
			// Query parameter
			"contiguous": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Indicates whether the IP addresses should belong to a contiguous block.\n\nDefaults to false.",
			},

			// Query parameter
			"ip_count": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     1,
				Description: "The number of IP addresses requested.\n\nDefaults to 1.",
			},

			"results": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        dataSourceSchemaOverrideFromResource(resourceIpamsvcAddress),
				Description: "List of IPs available under the resource defined by 'id'.\n\nThe schema of each element is identical to the b1ddi_address resource.",
			},
		},
	}
}

func dataSourceIpamsvcNaIPRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var (
		diags   diag.Diagnostics
		results []*models.IpamsvcAddress
	)

	addressStr := d.Get("id").(string)

	c := m.(*b1ddiclient.Client)

	switch addressStr[:strings.LastIndex(addressStr, "/")] {
	case "ipam/address_block":
		resp, err := c.IPAddressManagementAPI.AddressBlock.AddressBlockListNextAvailableIP(
			&address_block.AddressBlockListNextAvailableIPParams{
				Contiguous: swag.Bool(d.Get("contiguous").(bool)),
				Count:      swag.Int32(int32(d.Get("ip_count").(int))),
				ID:         addressStr,
				Context:    ctx,
			},
			nil,
		)
		if err != nil {
			return diag.FromErr(err)
		}
		results = resp.Payload.Results

	case "ipam/subnet":
		resp, err := c.IPAddressManagementAPI.Subnet.SubnetListNextAvailableIP(
			&subnet.SubnetListNextAvailableIPParams{
				Contiguous: swag.Bool(d.Get("contiguous").(bool)),
				Count:      swag.Int32(int32(d.Get("ip_count").(int))),
				ID:         addressStr,
				Context:    ctx,
			},
			nil,
		)
		if err != nil {
			return diag.FromErr(err)
		}
		results = resp.Payload.Results

	case "ipam/range":
		resp, err := c.IPAddressManagementAPI.RangeOperations.RangeListNextAvailableIP(
			&range_operations.RangeListNextAvailableIPParams{
				Contiguous: swag.Bool(d.Get("contiguous").(bool)),
				Count:      swag.Int32(int32(d.Get("ip_count").(int))),
				ID:         addressStr,
				Context:    ctx,
			},
			nil,
		)
		if err != nil {
			return diag.FromErr(err)
		}
		results = resp.Payload.Results

	}

	r := make([]interface{}, 0, len(results))
	for _, addr := range results {
		r = append(r, flattenIpamsvcAddress(addr)...)
	}

	err := d.Set("results", r)
	if err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

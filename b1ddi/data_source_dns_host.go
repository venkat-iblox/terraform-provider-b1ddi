package b1ddi

import (
	"context"
	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/dns_config/host"
	"strconv"
	"time"
)

func dataSourceDnsHost() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDnsHostRead,
		Schema: map[string]*schema.Schema{
			"filters": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Configure a map of filters to be applied on the search result.",
			},
			"results": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        schemaConfigHost(),
				Description: "List of DNS Hosts matching filters.",
			},
			"tfilters": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Configure a map of tag filters to be applied on the search result.",
			},
		},
	}
}

func dataSourceDnsHostRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	var diags diag.Diagnostics

	filtersMap := d.Get("filters").(map[string]interface{})
	filterStr := filterFromMap(filtersMap)

	tfilterMap := d.Get("tfilters").(map[string]interface{})
	tfilterStr := filterFromMap(tfilterMap)

	resp, err := c.DNSConfigurationAPI.Host.HostList(&host.HostListParams{
		Filter:  swag.String(filterStr),
		Tfilter: swag.String(tfilterStr),
		Context: ctx,
	}, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	results := make([]interface{}, 0, len(resp.Payload.Results))
	for _, h := range resp.Payload.Results {
		results = append(results, flattenConfigHost(h)...)
	}
	err = d.Set("results", results)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

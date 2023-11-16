package b1ddi

import (
	"context"
	"strconv"
	"time"

	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/dns_config/forward_zone"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

func dataSourceConfigForwardZone() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceConfigForwardZoneRead,
		Schema: map[string]*schema.Schema{
			"filters": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Configure a map of filters to be applied on the search result.",
			},
			"results": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        dataSourceSchemaFromResource(resourceConfigForwardZone),
				Description: "List of DNS Forward Zones matching filters. The schema of each element is identical to the b1ddi_dns_forward_zone resource schema.",
			},
		},
	}
}

func dataSourceConfigForwardZoneRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	var diags diag.Diagnostics

	filtersMap := d.Get("filters").(map[string]interface{})
	filterStr := filterFromMap(filtersMap)

	resp, err := c.DNSConfigurationAPI.ForwardZone.ForwardZoneList(&forward_zone.ForwardZoneListParams{
		Filter:  swag.String(filterStr),
		Context: ctx,
	}, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	results := make([]interface{}, 0, len(resp.Payload.Results))
	for _, fwdZone := range resp.Payload.Results {
		results = append(results, flattenConfigForwardZone(fwdZone)...)
	}
	err = d.Set("results", results)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenConfigForwardZone(r *models.ConfigForwardZone) []interface{} {
	if r == nil {
		return nil
	}

	externalForwarders := make([]map[string]interface{}, 0, len(r.ExternalForwarders))
	for _, ef := range r.ExternalForwarders {
		externalForwarders = append(externalForwarders, flattenConfigForwarder(ef))
	}

	return []interface{}{
		map[string]interface{}{
			"id":                  r.ID,
			"comment":             r.Comment,
			"created_at":          r.CreatedAt.String(),
			"disabled":            r.Disabled,
			"external_forwarders": externalForwarders,
			"forward_only":        r.ForwardOnly,
			"fqdn":                r.Fqdn,
			"hosts":               r.Hosts,
			"internal_forwarders": r.InternalForwarders,
			"mapped_subnet":       r.MappedSubnet,
			"mapping":             r.Mapping,
			"nsgs":                r.Nsgs,
			"parent":              r.Parent,
			"protocol_fqdn":       r.ProtocolFqdn,
			"tags":                r.Tags,
			"updated_at":          r.UpdatedAt.String(),
			"view":                r.View,
		},
	}
}

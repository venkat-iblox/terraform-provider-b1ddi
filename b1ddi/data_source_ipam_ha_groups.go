package b1ddi

import (
	"context"
	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/ha_group"
	"github.com/infobloxopen/b1ddi-go-client/models"
	"strconv"
	"time"
)

func dataSourceHAGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHAGroupRead,
		Schema: map[string]*schema.Schema{
			"filters": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Configure a map of filters to be applied on the search result.",
			},
			"results": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        dataSourceSchemaFromResource(resourceDataRecord),
				Description: "List of HA Groups matching filters. The schema of each element is identical to the b1ddi_ipam_ha_group resource schema.",
			},
		},
	}
}

func dataSourceHAGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	var diags diag.Diagnostics

	filtersMap := d.Get("filters").(map[string]interface{})
	filterStr := filterFromMap(filtersMap)

	resp, err := c.IPAddressManagementAPI.HaGroup.HaGroupList(&ha_group.HaGroupListParams{
		Filter:  swag.String(filterStr),
		Context: ctx,
	}, nil)

	if err != nil {
		return diag.FromErr(err)
	}

	results := make([]interface{}, 0, len(resp.Payload.Results))
	for _, ab := range resp.Payload.Results {
		results = append(results, flattenHAGroup(ab)...)
	}
	err = d.Set("results", results)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenHAGroup(r *models.IpamsvcHAGroup) []interface{} {
	if r == nil {
		return nil
	}

	return []interface{}{
		map[string]interface{}{
			"id":         r.ID,
			"comment":    r.Comment,
			"created_at": r.CreatedAt.String(),
			"tags":       r.Tags,
			"updated_at": r.UpdatedAt.String(),
			"mode":       r.Mode,
			"name":       r.Name,
			"ip_space":   r.IPSpace,
			"hosts":      flattenHAGroupsHosts(r.Hosts),
		},
	}
}

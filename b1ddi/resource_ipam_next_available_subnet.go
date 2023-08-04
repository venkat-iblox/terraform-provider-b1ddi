package b1ddi

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/address_block"
	"github.com/infobloxopen/b1ddi-go-client/models"
	"time"
)

// IpamsvcCreateNextAvailableSubnetResponse CreateNextAvailableABResponse
//
// The Next Available Subnet object create response format.
//
// swagger:model ipamsvcCreateNextAvailableSubnetResponse
func resourceIpamNextAvailableSubnet() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIpamNextAvailableSubnetCreate,
		ReadContext:   resourceIpamNextAvailableSubnetList,
		DeleteContext: resourceIpamNextAvailableSubnetDelete,
		Schema: map[string]*schema.Schema{
			"address_block_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the Address block",
			},
		},
	}
}

func resourceIpamNextAvailableSubnetCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	resp, err := c.IPAddressManagementAPI.AddressBlock.AddressBlockCreateNextAvailableSubnet(
		&address_block.AddressBlockCreateNextAvailableSubnetParams{
			ID:      d.Get("address_block_id").(string),
			Context: ctx,
		},
		nil,
	)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(resp.Payload.Results[0].ID)
	time.Sleep(time.Second)

	return resourceIpamNextAvailableSubnetList(ctx, d, m)
}

func resourceIpamNextAvailableSubnetList(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)
	var diags diag.Diagnostics

	resp, err := c.IPAddressManagementAPI.AddressBlock.AddressBlockListNextAvailableSubnet(
		&address_block.AddressBlockListNextAvailableSubnetParams{
			ID:      d.Id(),
			Context: ctx,
		},
		nil,
	)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("results", flattenIpamNextAvailableSubnet(resp.Payload))
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	return diags
}

func resourceIpamNextAvailableSubnetDelete(_ context.Context, _ *schema.ResourceData, _ interface{}) diag.Diagnostics {
	return nil
}

func flattenIpamNextAvailableSubnet(r *models.IpamsvcNextAvailableSubnetResponse) []interface{} {
	if r == nil {
		return nil
	}

	var sb []interface{}

	for _, result := range r.Results {
		sb = append(sb, flattenIpamsvcSubnet(result))
	}

	return sb
}

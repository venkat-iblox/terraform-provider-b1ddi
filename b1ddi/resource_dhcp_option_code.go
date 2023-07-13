package b1ddi

import (
	"context"
	"github.com/go-openapi/swag"
	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/option_code"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/infobloxopen/b1ddi-go-client/models"
)

// IpamsvcCreateOptionCodeResponse CreateOptionCodeResponse
//
// The response format to create the __OptionCode__ object.
//
// swagger:model ipamsvcCreateOptionCodeResponse
func resourceDhcpOptionCodeResponse() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDhcpOptionCodeResponseCreate,
		ReadContext:   resourceDhcpOptionCodeResponseRead,
		UpdateContext: resourceDhcpOptionCodeResponseUpdate,
		DeleteContext: resourceDhcpOptionCodeResponseDelete,
		Schema: map[string]*schema.Schema{

			// The created OptionCode object.
			"result": {
				Type:        schema.TypeList,
				Elem:        resourceIpamsvcOptionCode(),
				MaxItems:    1,
				Optional:    true,
				Description: "The created OptionCode object.",
			},
		},
	}
}

func resourceDhcpOptionCodeResponseCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(b1ddiclient.Client)

	body := &models.IpamsvcOptionCode{
		Array:       d.Get("array").(bool),
		Code:        swag.Int64(d.Get("code").(int64)),
		Comment:     d.Get("comment").(string),
		Name:        swag.String(d.Get("name").(string)),
		OptionSpace: swag.String(d.Get("option_space").(string)),
		Type:        swag.String(d.Get("type").(string)),
	}
	resp, err := c.IPAddressManagementAPI.OptionCode.OptionCodeCreate(
		&option_code.OptionCodeCreateParams{
			Body:    body,
			Context: ctx,
		},
		nil,
	)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.Payload.Result.ID)

	time.Sleep(time.Second)

	return resourceDhcpOptionCodeResponseRead(ctx, d, m)
}

func resourceDhcpOptionCodeResponseRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(b1ddiclient.Client)

	var diags diag.Diagnostics

	resp, err := c.IPAddressManagementAPI.OptionCode.OptionCodeRead(
		&option_code.OptionCodeReadParams{
			Fields:  swag.String(d.Get("fields").(string)),
			Context: ctx,
		},
		nil,
	)
	if err != nil {
		return diag.FromErr(err)
	}

	err = d.Set("result", flattenIpamsvcOptionCode(resp.Payload.Result))
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenDhcpOptionCodeResponse(r *models.IpamsvcCreateOptionCodeResponse) []interface{} {
	if r == nil {
		return nil
	}

	return []interface{}{
		map[string]interface{}{

			"result": flattenIpamsvcOptionCode(r.Result),
		},
	}
}

func resourceDhcpOptionCodeResponseUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)
	var diags diag.Diagnostics

	body := &models.IpamsvcOptionCode{
		Array:       d.Get("array").(bool),
		Code:        swag.Int64(d.Get("code").(int64)),
		Comment:     d.Get("comment").(string),
		Name:        swag.String(d.Get("name").(string)),
		OptionSpace: swag.String(d.Get("option_space").(string)),
		Type:        swag.String(d.Get("type").(string)),
	}

	c.IPAddressManagementAPI.OptionCode.OptionCodeUpdate(
		&option_code.OptionCodeUpdateParams{
			Body:    body,
			ID:      d.Id(),
			Context: ctx,
		},
		nil,
	)
	return diags
}

func resourceDhcpOptionCodeResponseDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	_, err := c.IPAddressManagementAPI.OptionCode.OptionCodeDelete(
		&option_code.OptionCodeDeleteParams{
			ID:      d.Id(),
			Context: ctx,
		},
		nil,
	)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")

	return nil
}

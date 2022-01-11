package b1ddi

import (
	"context"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("B1DDI_HOST", nil),
			},
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("B1DDI_TOKEN", nil),
			},
			"base_path": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "api/ddi/v1",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"b1ddi_ip_space":      resourceIpamsvcIPSpace(),
			"b1ddi_subnet":        resourceIpamsvcSubnet(),
			"b1ddi_fixed_address": resourceIpamsvcFixedAddress(),
			"b1ddi_address_block": resourceIpamsvcAddressBlock(),
			"b1ddi_range":         resourceIpamsvcRange(),
			"b1ddi_address":       resourceIpamsvcAddress(),
			"b1ddi_dns_view":      resourceConfigView(),
			"b1ddi_dns_record":    resourceDataRecord(),
		},
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	host := d.Get("host").(string)
	token := d.Get("token").(string)
	basePath := d.Get("base_path").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if host == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to initialise B1DDI client without API host",
		})
		return nil, diags
	}

	if token == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to initialise B1DDI client without API token",
		})
		return nil, diags
	}

	if basePath == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to initialise B1DDI client without API base path",
		})
		return nil, diags
	}

	// create the transport
	transport := httptransport.New(
		host, basePath, nil,
	)

	// Create default auth header for all API requests
	tokenAuth := b1ddiclient.B1DDIToken(token)
	transport.DefaultAuthentication = tokenAuth

	// create the API client
	c := b1ddiclient.NewClient(transport, strfmt.Default)

	return c, diags
}

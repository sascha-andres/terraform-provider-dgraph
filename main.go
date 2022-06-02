package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"livingit.de/code/tf-dgraph/resources"
)

// Provider returns a provider for terraform
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"server": {
				Type:        schema.TypeString,
				Optional:    false,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("DGRAPH_SERVER", "localhost:9080"),
				Description: "Connect to this dgraph server",
			},
			"insecure_skip_verify": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether the client should verify the server's certificate chain and host name",
			},
			"client_certificate": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "PEM-encoded client certificate for TLS authentication.",
			},
			"client_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "PEM-encoded client certificate key for TLS authentication.",
			},
			"ca_certificate": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "PEM-encoded root certificates bundle for TLS authentication.",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"dgraph_predicate": resourcePredicate(),
			"dgraph_type":      resourceType(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	m := resources.Meta{
		Client: resources.DeferredGetClient(d),
	}
	return m, nil
}

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return Provider()
		},
	})
}

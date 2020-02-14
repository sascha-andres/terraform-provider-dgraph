package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// Provider returns a provider for terraform
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"server": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    false,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("DGRAPH_SERVER", ""),
				Description: "Connect to this dgraph server",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"dgraph_predicate": resourcePredicate(),
			"dgraph_type":      resourceType(),
		},
	}
}

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return Provider()
		},
	})
}

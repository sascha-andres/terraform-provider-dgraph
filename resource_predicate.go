package main

import (
	"github.com/dgraph-io/dgo/v2"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"livingit.de/code/tf-dgraph/resources/predicate"
)

var client *dgo.Dgraph

func resourcePredicate() *schema.Resource {
	return &schema.Resource{
		Create: predicate.Create,
		Read:   predicate.Read,
		Update: predicate.Update,
		Delete: predicate.Delete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"array": &schema.Schema{
				Type:     schema.TypeBool,
				Required: false,
				Optional: true,
			},
			"lang": &schema.Schema{
				Type:     schema.TypeBool,
				Required: false,
				Optional: true,
			},
			"index": &schema.Schema{
				Type:     schema.TypeBool,
				Required: false,
				Optional: true,
			},
			"tokenizer": &schema.Schema{
				Type:        schema.TypeString,
				Required:    false,
				Optional:    true,
				Description: "Required when index is true",
			},
			"edge_count": &schema.Schema{
				Type:     schema.TypeBool,
				Required: false,
				Optional: true,
			},
		},
	}
}

package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"livingit.de/code/tf-dgraph/resources/dtype"
)

func resourceType() *schema.Resource {
	return &schema.Resource{
		Create: dtype.Create,
		Read:   dtype.Read,
		Update: dtype.Update,
		Delete: dtype.Delete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"fields": &schema.Schema{
				Type:     schema.TypeMap,
				Required: true,
			},
		},
	}
}

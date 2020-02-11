package main

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func resourcePredicateUpdate(d *schema.ResourceData, m interface{}) error {
	return resourcePredicateCreate(d, m)
}

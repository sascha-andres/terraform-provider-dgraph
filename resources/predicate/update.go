package predicate

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func Update(d *schema.ResourceData, m interface{}) error {
	return Create(d, m)
}

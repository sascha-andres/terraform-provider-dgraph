package predicate

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

// Update pushes changes for a type to Dgraph
func Update(d *schema.ResourceData, m interface{}) error {
	return Create(d, m)
}

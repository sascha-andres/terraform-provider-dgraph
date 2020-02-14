package dtype

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

// Update updates a type within Dgraph
func Update(d *schema.ResourceData, m interface{}) error {
	return Create(d, m)
}

package predicate

import (
	"context"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Delete removes a predicate from Dgraph
func Delete(d *schema.ResourceData, m interface{}) error {
	predicateName := d.Get("name").(string)

	err := client.Alter(context.Background(), &api.Operation{
		DropAttr: predicateName,
	})
	if err != nil {
		return err
	}

	return nil
}

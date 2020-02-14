package dtype

import (
	"context"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Delete removes a type from Dgraph
func Delete(d *schema.ResourceData, m interface{}) error {
	typeName := d.Get("name").(string)

	err := client.Alter(context.Background(), &api.Operation{
		DropOp:   api.Operation_TYPE,
		DropAttr: typeName,
	})
	if err != nil {
		return err
	}

	return nil
}

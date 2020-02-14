package predicate

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Read returns data about a predicate from Dgraph
func Read(d *schema.ResourceData, m interface{}) error {
	predicateName := d.Id()

	if len(predicateName) < 11 {
		return fmt.Errorf("id has wrong format: [%s] should be predicate_<id>", predicateName)
	}
	predicateName = predicateName[10:]

	data, err := GetPredicate(predicateName)
	if err != nil {
		return err
	}

	if nil == data.Schema {
		return nil
	}

	d.Set("name", data.Schema[0].Name)
	d.Set("type", data.Schema[0].Type)
	d.Set("array", data.Schema[0].List)
	d.Set("index", data.Schema[0].Index)
	if data.Schema[0].Index {
		d.Set("tokenizer", data.Schema[0].Tokenizer)
	}
	d.Set("edge_count", data.Schema[0].Count)
	d.Set("lang", data.Schema[0].Lang)

	return nil
}

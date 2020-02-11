package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"io/ioutil"
)

func resourcePredicateRead(d *schema.ResourceData, m interface{}) error {
	predicateName := d.Id()

	schemaQuery := fmt.Sprintf(`schema(pred: [%s]) {
  type
  index
  reverse
  tokenizer
  list
  count
  upsert
  lang
}`, predicateName)

	resp, err := client.NewReadOnlyTxn().Do(context.Background(), &api.Request{
		Query: schemaQuery,
	})
	if err != nil {
		return err
	}

	ioutil.WriteFile("out.txt", resp.Json, 0600)

	var data resourceData
	err = json.Unmarshal(resp.Json, &data)
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

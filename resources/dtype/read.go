package dtype

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"io/ioutil"
	"livingit.de/code/tf-dgraph/resources"
	"livingit.de/code/tf-dgraph/resources/predicate"
)

func Read(d *schema.ResourceData, m interface{}) error {

	typeName := d.Id()

	if len(typeName) < 6 {
		return fmt.Errorf("id has wrong format: [%s] should be type_<id>", typeName)
	}
	typeName = typeName[5:]

	schemaQuery := fmt.Sprintf(`schema(type: [%s]){}`, typeName)

	resp, err := client.NewReadOnlyTxn().Do(context.Background(), &api.Request{
		Query: schemaQuery,
	})
	if err != nil {
		return err
	}

	var data resources.ResourceTypeData
	err = json.Unmarshal(resp.Json, &data)
	if err != nil {
		return err
	}

	if nil == data.Types {
		return nil
	}

	d.Set("name", data.Types[0].Name)

	fields := d.Get("fields").(map[string]interface{})
	for _, v := range data.Types[0].Fields {
		if _, ok := fields[v.Name]; !ok { // expected field not found
			t, err := getPredicateType(v.Name)
			if err != nil {
				return err
			}
			fields[v.Name] = t
		}
	}
	for k, _ := range fields {
		n := ""
		for _, f := range data.Types[0].Fields {
			if f.Name == k {
				n = f.Name
				break
			}
		}
		if n == "" { // locally field is expected, remote non existent
			delete(fields, n)
		} else {
			t, err := getPredicateType(n)
			if err != nil {
				return err
			}
			fields[n] = t
		}
	}
	d.Set("fields", fields)

	ioutil.WriteFile(fmt.Sprintf("out-%s-fields.txt", typeName), []byte(fmt.Sprintf("%#v", fields)), 0600)

	return nil

	/*
		d.Set("type", data.Schema[0].Type)
		d.Set("array", data.Schema[0].List)
		d.Set("index", data.Schema[0].Index)
		if data.Schema[0].Index {
			d.Set("tokenizer", data.Schema[0].Tokenizer)
		}
		d.Set("edge_count", data.Schema[0].Count)
		d.Set("lang", data.Schema[0].Lang)

		return nil*/
}

// getPredicateType is a helper to get the type of a predicate
func getPredicateType(predicateName string) (string, error) {
	data, err := predicate.GetPredicate(predicateName)
	if err != nil {
		return "", err
	}
	return data.Schema[0].Type, nil
}

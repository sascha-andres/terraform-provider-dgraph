package dtype

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"livingit.de/code/tf-dgraph/resources"
	"livingit.de/code/tf-dgraph/resources/predicate"
)

// Read retrieves information about a type from Dgraph
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

	fields, err := adjustForChanges(d, data)
	if err != nil {
		return err
	}
	d.Set("fields", fields)

	return nil
}

// adjustForChanges compares local config with remote state
func adjustForChanges(d *schema.ResourceData, data resources.ResourceTypeData) (map[string]interface{}, error) {
	fields := d.Get("fields").(map[string]interface{})
	for _, v := range data.Types[0].Fields {
		if _, ok := fields[v.Name]; !ok { // expected field not found
			t, err := getPredicateType(v.Name)
			if err != nil {
				return nil, err
			}
			fields[v.Name] = t
		}
	}
	for k := range fields {
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
				return nil, err
			}
			fields[n] = t
		}
	}
	return fields, nil
}

// getPredicateType is a helper to get the type of a predicate
func getPredicateType(predicateName string) (string, error) {
	data, err := predicate.GetPredicate(predicateName)
	if err != nil {
		return "", err
	}
	return data.Schema[0].Type, nil
}

package predicate

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"livingit.de/code/tf-dgraph/resources"
)

// GetPredicate returns data about a single predicate
func GetPredicate(predicateName string, client *dgo.Dgraph) (*resources.ResourcePredicateData, error) {
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
		return nil, err
	}

	var data resources.ResourcePredicateData
	err = json.Unmarshal(resp.Json, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

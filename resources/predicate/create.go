package predicate

import (
	"context"
	"fmt"

	"github.com/dgraph-io/dgo/v2/protos/api"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"livingit.de/code/tf-dgraph/resources"
)

const iDPredicateTemplate = "predicate_%s"

// Create adds a predicate to Dgraph
func Create(d *schema.ResourceData, m interface{}) error {
	client, err := m.(resources.Meta).Client()
	if err != nil {
		return err
	}

	predicateName := d.Get("name").(string)
	predicateType := d.Get("type").(string)
	isArray := d.Get("array").(bool)
	if isArray {
		predicateType = fmt.Sprintf("[%s]", predicateType)
	}
	predicateIndex := d.Get("index").(bool)
	predicateTokenizer := d.Get("tokenizer").(string)
	if predicateIndex && "" == predicateTokenizer {
		return fmt.Errorf("tokenizer must be set for %s when index is true", predicateName)
	}
	predicateCount := ""
	if d.Get("edge_count").(bool) {
		predicateCount = "@count"
	}
	predicateLang := ""
	if d.Get("lang").(bool) {
		if predicateType != "string" || (predicateType == "string" && isArray) {
			return fmt.Errorf("lang for %s may only be set if of dtype string", predicateName)
		}
		predicateLang = "@lang"
	}

	if predicateIndex {
		predicateTokenizer = fmt.Sprintf("@index(%s)", predicateTokenizer)
	} else {
		predicateTokenizer = ""
	}

	err = client.Alter(context.Background(), &api.Operation{
		Schema: fmt.Sprintf("%s: %s %s %s %s .", predicateName, predicateType, predicateTokenizer, predicateLang, predicateCount),
	})

	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf(iDPredicateTemplate, predicateName))

	return Read(d, m)
}

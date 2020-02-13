package predicate

import (
	"github.com/dgraph-io/dgo/v2"
	"livingit.de/code/tf-dgraph/resources"
)

var client *dgo.Dgraph

func init() {
	client = resources.NewClient()
}

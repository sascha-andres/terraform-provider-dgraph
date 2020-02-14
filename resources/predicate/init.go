package predicate

import (
	"github.com/dgraph-io/dgo/v2"
	"livingit.de/code/tf-dgraph/resources"
)

var client *dgo.Dgraph

// init used to create a Dgraph client
func init() {
	client = resources.NewClient()
}

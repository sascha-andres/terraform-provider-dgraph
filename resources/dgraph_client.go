package resources

import (
	"sync"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"google.golang.org/grpc"
)

var (
	client     *dgo.Dgraph
	clientOnce sync.Once
)

type Meta struct {
	Client func() (*dgo.Dgraph, error)
}

// DeferredGetClient returns a function that returns a dgraph client
func DeferredGetClient(d *schema.ResourceData) func() (*dgo.Dgraph, error) {
	return func() (*dgo.Dgraph, error) {
		var err error

		clientOnce.Do(func() {
			// Dial a gRPC connection. The address to dial to can be configured when
			// setting up the dgraph cluster.
			d, err := grpc.Dial(d.Get("server").(string), grpc.WithInsecure())
			if err == nil {
				client = dgo.NewDgraphClient(
					api.NewDgraphClient(d),
				)
			}
		})

		return client, err
	}
}

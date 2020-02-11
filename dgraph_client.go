package main

import (
	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"google.golang.org/grpc"
	"log"
)

func newClient() *dgo.Dgraph {
	// Dial a gRPC connection. The address to dial to can be configured when
	// setting up the dgraph cluster.
	d, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return dgo.NewDgraphClient(
		api.NewDgraphClient(d),
	)
}

type resourceData struct {
	Schema []struct {
		Name      string `json:"predicate"`
		Type      string
		Index     bool
		Reverse   string
		Tokenizer []string
		List      bool
		Count     bool
		Upsert    string
		Lang      bool
	}
}

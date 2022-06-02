package resources

import (
	"crypto/tls"
	"crypto/x509"
	"sync"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	client     *dgo.Dgraph
	clientOnce sync.Once
)

type Meta struct {
	Client func() (*dgo.Dgraph, error)
}

func getClient(d *schema.ResourceData) (*dgo.Dgraph, error) {
	// Dial a gRPC connection. The address to dial to can be configured when
	// setting up the dgraph cluster.
	conn, err := grpc.Dial(d.Get("server").(string), grpc.WithInsecure())
	if err == nil {
		client = dgo.NewDgraphClient(
			api.NewDgraphClient(conn),
		)
	}

	return client, err
}

func getClientTLS(d *schema.ResourceData) (*dgo.Dgraph, error) {
	cert := []byte(d.Get("client_certificate").(string))
	key := []byte(d.Get("client_key").(string))

	keyPair, err := tls.X509KeyPair(cert, key)

	config := &tls.Config{
		InsecureSkipVerify: d.Get("insecure_skip_verify").(bool),
		Certificates:       []tls.Certificate{keyPair},
	}

	if ca, ok := d.GetOk("ca_certificate"); ok {
		rootCAs := x509.NewCertPool()
		rootCAs.AppendCertsFromPEM([]byte(ca.(string)))

		config.RootCAs = rootCAs
	}

	credsClient := credentials.NewTLS(config)

	// Dial a gRPC connection. The address to dial to can be configured when
	// setting up the dgraph cluster.
	conn, err := grpc.Dial(d.Get("server").(string), grpc.WithTransportCredentials(credsClient))
	if err == nil {
		client = dgo.NewDgraphClient(
			api.NewDgraphClient(conn),
		)
	}

	return client, err
}

// DeferredGetClient returns a function that returns a dgraph client
func DeferredGetClient(d *schema.ResourceData) func() (*dgo.Dgraph, error) {
	return func() (*dgo.Dgraph, error) {
		var err error

		clientOnce.Do(func() {
			if _, ok := d.GetOk("client_certificate"); ok {
				client, err = getClientTLS(d)
			} else {
				client, err = getClient(d)
			}
		})

		return client, err
	}
}

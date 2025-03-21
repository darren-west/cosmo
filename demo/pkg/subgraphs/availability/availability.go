package availability

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/wundergraph/graphql-go-tools/v2/pkg/engine/datasource/pubsub_datasource"

	"github.com/wundergraph/cosmo/demo/pkg/subgraphs/availability/subgraph"
	"github.com/wundergraph/cosmo/demo/pkg/subgraphs/availability/subgraph/generated"
)

func NewSchema(pubSubBySourceName map[string]pubsub_datasource.NatsPubSub, pubSubName func(string) string) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{Resolvers: &subgraph.Resolver{
		NatsPubSubByProviderID: pubSubBySourceName,
		GetPubSubName:          pubSubName,
	}})
}

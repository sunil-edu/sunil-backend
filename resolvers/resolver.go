package resolvers

import (
	"myeduate"
	"myeduate/ent"

	"github.com/99designs/gqlgen/graphql"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{ client *ent.Client }

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return myeduate.NewExecutableSchema(myeduate.Config{
		Resolvers: &Resolver{client},
	})
}

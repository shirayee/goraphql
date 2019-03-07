package schema

import (
	"github.com/graphql-go/graphql"
	"goraphql/app/domain/mutation"
	"goraphql/app/domain/query"
)

var GoraphQLSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    query.RootQuery,
	Mutation: mutation.RootMutation,
})

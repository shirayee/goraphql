package mutation

import (
	"github.com/graphql-go/graphql"
	"goraphql/app/domain/fields"
)

var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createUser": fields.CreateUserField,
	},
})

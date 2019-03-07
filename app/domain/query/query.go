package query

import (
	"github.com/graphql-go/graphql"
	"goraphql/app/domain/fields"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"user": fields.UserField,
		"userList": fields.UserListField,
	},
})

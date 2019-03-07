package fields

import (
	"errors"
	"github.com/graphql-go/graphql"
	"goraphql/app/application/usecase"
	"goraphql/app/domain/model"
	"goraphql/app/domain/types"
	"goraphql/app/infrastructure/datastore"
)

var UserField = &graphql.Field{
	Type:        types.UserType,
	Description: "Get user object",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		id, ok := params.Args["id"].(int)
		if ok {
			return usecase.NewUserUseCase(datastore.NewUserRepository()).GetUserByID(id)
		}
		return nil, errors.New("id is invalid")
	},
}

var UserListField = &graphql.Field{
	Type: graphql.NewList(types.UserType),
	Description: "List of users",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		return usecase.NewUserUseCase(datastore.NewUserRepository()).GetUsers()
	},
}

var CreateUserField = &graphql.Field{
	Type:        types.UserType,
	Description: "Create new user",
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		name, ok := params.Args["name"].(string)
		if !ok {
			return nil, errors.New("name is invalid")
		}
		newUser, err := model.NewUser(name)
		if err != nil {
			return nil, err
		}
		newUser, err = usecase.NewUserUseCase(datastore.NewUserRepository()).CreateUser(newUser)
		return newUser, err
	},
}

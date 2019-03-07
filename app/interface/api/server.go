package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"goraphql/app/domain/schema"
	"io/ioutil"
	"net/http"
)

func Start() {
	http.HandleFunc("/graphql", func(writer http.ResponseWriter, request *http.Request) {
		defer request.Body.Close()

		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			fmt.Printf("%#v\n", err)
		}

		result, err := execQuery(fmt.Sprintf("%s", body), schema.GoraphQLSchema)
		if err != nil {
			fmt.Printf("%#v\n", err)
			return
		}
		json.NewEncoder(writer).Encode(result)
	})

	fmt.Println("server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func execQuery(query string, schema graphql.Schema) (*graphql.Result, error) {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		return nil, errors.New(fmt.Sprintf("wrong result, unexpected errors: %v\n", result.Errors))

	}
	return result, nil
}

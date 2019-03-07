package main

import (
	"fmt"
	"goraphql/app/domain/service"
	"goraphql/app/interface/api"
)

func main() {
	err := service.NewDbService().Migrate()
	if err != nil {
		fmt.Printf("%#v\n", err)
		panic(err)
	}
	api.Start()
}

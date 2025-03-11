package main

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {
	fmt.Println("Testing imports")
	_ = graphql.SchemaConfig{}
	_ = handler.Config{}
}

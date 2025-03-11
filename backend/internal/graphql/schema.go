// Package graphql provides GraphQL functionality for the JSON prettifier
package graphql

import (
	"github.com/graphql-go/graphql"
)

// CreateSchema creates and returns the GraphQL schema
func CreateSchema() (graphql.Schema, error) {
	// Define the root query
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"formatJSON": &graphql.Field{
				Type: graphql.String,
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: formatJSONResolver,
			},
		},
	})

	// Create the schema config
	schemaConfig := graphql.SchemaConfig{
		Query: rootQuery,
	}

	// Create and return the schema
	return graphql.NewSchema(schemaConfig)
}

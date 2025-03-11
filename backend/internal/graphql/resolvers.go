// Package graphql provides GraphQL functionality for the JSON prettifier
package graphql

import (
	"encoding/json"

	"github.com/graphql-go/graphql"
)

// formatJSONResolver handles the formatJSON query
func formatJSONResolver(p graphql.ResolveParams) (interface{}, error) {
	// Get the input JSON string from arguments
	inputJSON, ok := p.Args["input"].(string)
	if !ok {
		return nil, nil
	}

	// Parse and format the JSON
	var temp interface{}
	if err := json.Unmarshal([]byte(inputJSON), &temp); err != nil {
		return nil, err
	}

	// Pretty print the JSON
	formattedJSON, err := json.MarshalIndent(temp, "", "  ")
	if err != nil {
		return nil, err
	}

	return string(formattedJSON), nil
}

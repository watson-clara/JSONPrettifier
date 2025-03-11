package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"

	gql "github.com/watson-clara/JSONPrettifier/backend/internal/graphql"
)

func main() {
	r := gin.Default()

	// Enable CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	r.Use(cors.New(config))

	// Set up GraphQL
	schema, err := gql.CreateSchema()
	if err != nil {
		log.Fatalf("Failed to create GraphQL schema: %v", err)
	}

	// Create a GraphQL HTTP handler
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true, // Enable GraphiQL UI
	})

	// Routes
	r.POST("/api/format", formatJSON)

	// GraphQL endpoint
	r.POST("/graphql", func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	})

	// GraphiQL UI
	r.GET("/graphiql", func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	})

	log.Println("Server starting on :8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func formatJSON(c *gin.Context) {
	var input struct {
		JSON string `json:"json"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	formatted, err := prettyJSON([]byte(input.JSON))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON input",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"formatted": string(formatted),
	})
}

func prettyJSON(input []byte) ([]byte, error) {
	var temp interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, err
	}
	return json.MarshalIndent(temp, "", "  ")
}

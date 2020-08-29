package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"

	"github.com/neekomar/go-graph/dynamo"
	"github.com/neekomar/go-graph/mutation"
	"github.com/neekomar/go-graph/query"
)

func main() {
	db := dynamo.InitSession()

	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "Query",
			Fields: query.GetQueries(db),
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name:   "Mutation",
			Fields: mutation.GetMutations(db),
		}),
	}

	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		log.Fatalf("Failed to create new schema, error: %v", err)
	}

	httpHandler := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})

	http.Handle("/", httpHandler)
	log.Print("ready: listening at http://localhost:8000\n")

	http.ListenAndServe(":8000", nil)
}

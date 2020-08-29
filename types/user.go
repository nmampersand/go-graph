package types

import (
	"github.com/graphql-go/graphql"
)

// User type definition.
type User struct {
	ID        string `db:"id" json:"id"`
	Firstname string `db:"firstname" json:"firstname"`
	Lastname  string `db:"lastname" json:"lastname"`
}

// UserType is the GraphQL schema for the user type.
var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id":        &graphql.Field{Type: graphql.String},
		"firstname": &graphql.Field{Type: graphql.String},
		"lastname":  &graphql.Field{Type: graphql.String},
	},
})

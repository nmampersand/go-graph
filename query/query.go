package query

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/graphql-go/graphql"
)

var usersTable = "User"

// GetQueries returns all the available queries.
func GetQueries(db *dynamodb.DynamoDB) graphql.Fields {
	return graphql.Fields{
		"user":  GetUserQuery(db, usersTable),
		"users": GetUsersQuery(db, usersTable),
	}
}

package mutation

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/graphql-go/graphql"
)

var usersTable = "User"

// GetMutations returns all the available mutations.
func GetMutations(db *dynamodb.DynamoDB) graphql.Fields {
	return graphql.Fields{
		"createUser": CreateUserMutation(db, usersTable),
	}
}

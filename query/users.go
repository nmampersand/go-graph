package query

import (
	"github.com/neekomar/go-graph/types"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/graphql-go/graphql"
)

// GetUsersQuery returns the queries available against user type.
func GetUsersQuery(db *dynamodb.DynamoDB, tableName string) *graphql.Field {

	return &graphql.Field{
		Type: graphql.NewList(types.UserType),
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			users := []types.User{}

			result, err := db.Scan(&dynamodb.ScanInput{
				TableName: aws.String(tableName),
			})

			if err != nil {
				return nil, err
			}

			err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &users)
			if err != nil {
				return nil, err
			}

			return users, nil
		},
	}
}

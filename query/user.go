package query

import (
	"github.com/neekomar/go-graph/types"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/graphql-go/graphql"
)

// GetUserQuery returns the queries available against user type.
func GetUserQuery(db *dynamodb.DynamoDB, tableName string) *graphql.Field {

	return &graphql.Field{
		Type: types.UserType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			user := types.User{}

			result, err := db.GetItem(&dynamodb.GetItemInput{
				TableName: aws.String(tableName),
				Key: map[string]*dynamodb.AttributeValue{
					"id": {
						S: aws.String(params.Args["id"].(string)),
					},
				},
			})

			if err != nil {
				return nil, err
			}

			err = dynamodbattribute.UnmarshalMap(result.Item, &user)
			if err != nil {
				return nil, err
			}

			return user, nil
		},
	}
}

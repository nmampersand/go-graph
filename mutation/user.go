package mutation

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
	"github.com/neekomar/go-graph/types"
)

// CreateUserMutation creates a new user and returns it.
func CreateUserMutation(db *dynamodb.DynamoDB, tableName string) *graphql.Field {

	return &graphql.Field{
		Type: types.UserType,
		Args: graphql.FieldConfigArgument{
			"firstname": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"lastname": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			user := &types.User{
				ID:        uuid.New().String(),
				Firstname: params.Args["firstname"].(string),
				Lastname:  params.Args["lastname"].(string),
			}

			newUser, err := dynamodbattribute.MarshalMap(user)
			if err != nil {
				return nil, err
			}

			_, err = db.PutItem(&dynamodb.PutItemInput{
				Item:      newUser,
				TableName: aws.String(tableName),
			})
			if err != nil {
				return nil, err
			}

			return user, nil
		},
	}
}

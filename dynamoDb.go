package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const usersTable = "SlackifyUsers"

func DynamodbClient() *dynamodb.DynamoDB {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{
			Region:      aws.String(os.Getenv("AWS_REGION")),
			Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), ""),
		},
	}))
	// Create and return a DynamoDB client
	client := dynamodb.New(sess)
	return client
}

// This function relies on the result of TableExists(usersTable)
// if the false the table will be created
func CreateUserTable() {

	client := DynamodbClient()

	if !TableExists(usersTable) {
		fmt.Printf("%s dynamoDb table doesn't exist, creating it...", usersTable)
		input := &dynamodb.CreateTableInput{
			AttributeDefinitions: []*dynamodb.AttributeDefinition{
				{
					AttributeName: aws.String("Id"),
					AttributeType: aws.String("S"),
				},
				{
					AttributeName: aws.String("Username"),
					AttributeType: aws.String("S"),
				},
			},
			KeySchema: []*dynamodb.KeySchemaElement{
				{
					AttributeName: aws.String("Id"),
					KeyType:       aws.String("HASH"),
				},
				{
					AttributeName: aws.String("Username"),
					KeyType:       aws.String("RANGE"),
				},
			},
			BillingMode: aws.String("PAY_PER_REQUEST"),
			TableName:   aws.String(usersTable),
		}

		_, err := client.CreateTable(input)

		if err != nil {
			log.Fatal(err)
		}

		// wait for the table to be provisioned:

		wErr := client.WaitUntilTableExists(&dynamodb.DescribeTableInput{
			TableName: aws.String(usersTable),
		})

		if wErr != nil {
			fmt.Println(wErr)
		}

		fmt.Println("Done!")
	} else {
		fmt.Printf("%s dynamoDb table already exists, skipping \n", usersTable)
	}

}

// We can use this function to check if a dynamoDb
// table exists or not.
// This is used to handle first run scnearios
func TableExists(tableName string) bool {

	client := DynamodbClient()

	_, err := client.Scan(&dynamodb.ScanInput{
		TableName: aws.String(tableName),
	})
	if err != nil {
		return false
	} else {
		return true
	}
}

func PutUser(user UserModel) {
	client := DynamodbClient()

	existingUser := GetUser(user)

	if len(existingUser.Id) > 0 {
		fmt.Println("User already exists")
		return
	}

	av, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		log.Fatalf("Error adding user: %s", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(usersTable),
	}

	_, err = client.PutItem(input)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
	}

	fmt.Printf("Successfully added new User %s'", user.Username)
}

func GetUser(user UserModel) UserModel {
	client := DynamodbClient()

	result, err := client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(usersTable),
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				S: aws.String(user.Id),
			},
			"Username": {
				S: aws.String(user.Username),
			},
		},
	})

	if err != nil {
		log.Fatalf("Go error getting user: %s", err)
	}

	if result.Item != nil {
		existingUser := UserModel{}

		err = dynamodbattribute.UnmarshalMap(result.Item, &existingUser)

		if err != nil {
			log.Fatal(err)
		}
		return existingUser
	}
	return UserModel{}
}

package dao

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var svc *dynamodb.DynamoDB

func InitializeDB() {
	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	if err != nil {
		panic(err)
	}

	// Create DynamoDB client
	svc = dynamodb.New(sess)
}

func Put(item interface{}) error {

	av, err := dynamodbattribute.MarshalMap(item)

	// Create item in table Movies
	input := &dynamodb.PutItemInput{
		Item: av,
		TableName: aws.String("TranslationsEngine"),
	}

	_, err = svc.PutItem(input)

	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func Update(item interface{}) error {
	av, err := dynamodbattribute.MarshalMap(item)

	// Create item in table Movies
	input := &dynamodb.UpdateItemInput{
		Key: av,
		TableName: aws.String("TranslationsEngine"),
	}

	_, err = svc.UpdateItem(input)

	if err != nil {
		fmt.Println("Got error calling UpdateItem:")
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func Delete(item interface{}) error {
	av, err := dynamodbattribute.MarshalMap(item)

	// Create item in table Movies
	input := &dynamodb.DeleteItemInput{
		Key: av,
		TableName: aws.String("TranslationsEngine"),
	}

	_, err = svc.DeleteItem(input)

	if err != nil {
		fmt.Println("Got error calling DeleteItem:")
		fmt.Println(err.Error())
		return err
	}

	return nil
}

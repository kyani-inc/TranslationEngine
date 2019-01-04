package translation_key

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/catmullet/TranslationEngine/database"
	"os"
)

func Put(item TranslationKeys) error {

	av, err := dynamodbattribute.MarshalMap(item)

	// Create item in table Movies
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("TranslationsEngine"),
	}

	_, err = database.SVC.PutItem(input)

	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func Get(locale string) TranslationKeys {
	item := TranslationKeys{}

	localefilter := expression.Name("locale").Equal(expression.Value(locale))
	proj := expression.NamesList(expression.Name("locale"), expression.Name("keys"))

	expr, err := expression.NewBuilder().WithFilter(localefilter).WithProjection(proj).Build()

	if err != nil {
		fmt.Println("Got error building expression:")
		fmt.Println(err.Error())
	}

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String("TranslationsEngine"),
	}

	result, err := database.SVC.Scan(params)

	if err != nil {
		fmt.Println("Query API call failed:")
		fmt.Println((err.Error()))
	}

	for _, i := range result.Items {
		err = dynamodbattribute.UnmarshalMap(i, &item)

		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		return item
	}

	return item
}

func GetAll() TranslationKeysList {

	list := TranslationKeysList{}
	item := TranslationKeys{}

	proj := expression.NamesList(expression.Name("locale"), expression.Name("keys"))

	expr, err := expression.NewBuilder().WithProjection(proj).Build()

	if err != nil {
		fmt.Println("Got error building expression:")
		fmt.Println(err.Error())
	}

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String("TranslationsEngine"),
	}

	result, err := database.SVC.Scan(params)

	if err != nil {
		fmt.Println("Query API call failed:")
		fmt.Println((err.Error()))
	}

	for _, i := range result.Items {
		err = dynamodbattribute.UnmarshalMap(i, &item)

		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		list.List = append(list.List, item)
	}

	return list
}

func Delete(item interface{}) error {
	av, err := dynamodbattribute.MarshalMap(item)

	// Create item in table Movies
	input := &dynamodb.DeleteItemInput{
		Key:       av,
		TableName: aws.String("TranslationsEngine"),
	}

	_, err = database.SVC.DeleteItem(input)

	if err != nil {
		fmt.Println("Got error calling DeleteItem:")
		fmt.Println(err.Error())
		return err
	}

	return nil
}

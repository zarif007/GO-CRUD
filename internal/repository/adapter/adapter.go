package adapter

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Database struct {
	connection *dynamodb.DynamoDB
	logMode    bool
}

type Interface interface {
}

func NewAdapter() Interface {

}
 
func (db *Database) Health() bool {
	_, err := db.connection.ListTables(&dynamodb.ListTablesInput{})

	return err == nil
}

func (db *Database) FindAll {

}

func (db *Database) FindOne(condition map[string]interface{}, tableName)(response *dynamodb.GetItemOutput, err error) {

	conditionaParsed, err := dynamodbattribute.MarshalMap(condition)

	if err != nil {
		return nil, err 
	}

	input := &dynamodb.GetItemInput{
		tableName: aws.String(tableName),
		key: conditionaParsed,
	}

	return db.connection.GetItem(input)
}

func (db *Database) CreateOrUpdate(entity interface[], tablename string)(response *dynamodb.PutItemOutput, err error) {

	entityParsed, err := dynamodbattribute.MarshalMap(entity)

	if err != nil {
		return nil, err 
	}

	input := &dynamodb.PutItemInput{
		Iten : entityParsed,
		tableName: aws.String(tableName),
	}

	return db.connection.PutItem(input)
}

func (db *Database) Delete(condition map[string]interface{}, tableName string)(response *dynamodb.DeleteItemOutput, err error) {
	conditionaParsed, err := dynamodbattribute.MarshalMap(condition)

	item := &dynamodb.DeleteItemOutput{
		key: conditionParsed,
		tableName: aws.String(tableName),
	}

	return db.connection.DeleteItem(input)
}




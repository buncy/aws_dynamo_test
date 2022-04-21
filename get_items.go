package main

import(
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func getItems(svc *dynamodb.Client)error{
	params := &dynamodb.BatchGetItemInput{}
	if output,err := svc.ExecuteStatement()

}
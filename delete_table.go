package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func dropTable(svc *dynamodb.Client) error {

	params := &dynamodb.DeleteTableInput{
		TableName: aws.String("Music"),
	}
	if output, err := svc.DeleteTable(context.Background(), params); err != nil {
		return err
	} else {
		fmt.Println(output)
		return err
	}

}

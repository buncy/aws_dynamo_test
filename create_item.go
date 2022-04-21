package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func createItem(svc *dynamodb.Client) error {
	params := &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]types.WriteRequest{
			"Music": {
				{
					// DeleteRequest: &types.DeleteRequest{
					// 	Key: map[string]types.AttributeValue{

					// 	},
					// },
					PutRequest: &types.PutRequest{
						Item: map[string]types.AttributeValue{
							"Artist":     &types.AttributeValueMemberS{Value: "No One You Know"},
							"SongTitle":  &types.AttributeValueMemberS{Value: "Call Me Today"},
							"AlbumTitle": &types.AttributeValueMemberS{Value: "Somewhat Famous"},
							"Awards":     &types.AttributeValueMemberN{Value: "1"},
						},
					},
				},
			},
		},
		ReturnConsumedCapacity:      "",
		ReturnItemCollectionMetrics: "",
	}
	if output, err := svc.BatchWriteItem(context.Background(), params); err != nil {
		return err
	} else {
		fmt.Println(output)
		return err
	}
}

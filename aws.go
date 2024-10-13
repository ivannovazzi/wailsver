package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var dClient *dynamodb.Client

// createDynamoDBClient initializes a DynamoDB client using AWS SSO stored configuration.
func createDynamoDBClient(context context.Context) (*dynamodb.Client, error) {
	// Load AWS configuration using SSO
	cfg, err := config.LoadDefaultConfig(context, config.WithSharedConfigProfile("AdministratorAccess-159675774736"))
	if err != nil {
		return nil, fmt.Errorf("unable to load AWS SDK config: %v", err)
	}
	// Create DynamoDB client
	client := dynamodb.NewFromConfig(cfg)
	return client, nil
}

// GetVersion reads the 'version' from the 'flare-platform-staging' table where key is 'service' and has value 'frontend'.
func (a *App) GetVersion() (string, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(a.config.TableName),
		Key: map[string]types.AttributeValue{
			"Service": &types.AttributeValueMemberS{Value: "Frontend"},
		},
	}

	result, err := dClient.GetItem(a.ctx, input)
	if err != nil {
		return "", fmt.Errorf("failed to get item: %v", err)
	}

	versionAttr, ok := result.Item["Version"].(*types.AttributeValueMemberS)
	if !ok {
		return "", fmt.Errorf("version attribute not found or not a string")
	}

	return versionAttr.Value, nil
}

// UpdateVersion updates the 'version' in the 'flare-platform-staging' table for key 'frontend'.
func (a *App) UpdateVersion(ctx context.Context, newVersion string) error {
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(a.config.TableName),
		Key: map[string]types.AttributeValue{
			"Service": &types.AttributeValueMemberS{Value: "Frontend"},
		},
		UpdateExpression: aws.String("SET Version = :v"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":v": &types.AttributeValueMemberS{Value: newVersion},
		},
	}

	_, err := dClient.UpdateItem(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to update item: %v", err)
	}

	return nil
}

// Example usage
func prepareClient(ctx context.Context) {
	// create a DynamoDB client if it does not exist
	if dClient != nil {
		return
	}
	// create a new client
	client, err := createDynamoDBClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create DynamoDB client: %v", err)
	}
	// save the client in the global variable
	dClient = client
}

package test

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"os"
	"testing"
	"time"
)

const AwsRegionCn = "cn-north-1"
const DynamoDBLocal = "http://localhost:8000"
const ViewTable = "share_view"
const ShareTable = "share_share"

/**
  https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#example_DynamoDB_Query_shared00
*/

type ShareItem struct {
	TargetId   string `json:"target_id"` // hash key
	ID         string `json:"id"`        // range key
	CratedTime int    `json:"crated_time"`
	ParentId   string `json:"parent_id"`
	TargetType string `json:"target_type"`
	RootId     string `json:"root_id"`
	ShareBy    string `json:"share_by"`
	CreatedBy  string `json:"created_by"`
}
type ViewItem struct {
	ViewId            string `json:"view_id"`
	CreatedTime       int64  `json:"created_time"`
	TargetType        string `json:"target_type"`       // BGC\UGC
	ViewType          string `json:"view_type"`         // READ\SHARE
	ViewChannelType   string `json:"view_channel_type"` // WECHAT
	ViewUserType      string `json:"view_user_type"`    // WECHAT, UNKNOW(未知), ALLINONE
	ViewUserId        string `json:"view_user_id"`
	ViewUserNickname  string `json:"view_user_nickname"`
	ViewUserAvatarUrl string `json:"view_user_avatar_url"`
	ViewUserPhone     string `json:"view_user_phone"`
}

func initDynamoDBClient() *dynamodb.DynamoDB {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(AwsRegionCn),
		Endpoint:    aws.String(DynamoDBLocal),
		Credentials: credentials.NewStaticCredentials("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY", ""),
	})

	if err != nil {
		fmt.Errorf(err.Error())
		return nil
	}

	// Create DynamoDB client
	return dynamodb.New(sess)

}

func TestListTable(t *testing.T) {
	svc := initDynamoDBClient()
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()

	// create the input configuration instance
	input := &dynamodb.ListTablesInput{}

	fmt.Printf("Tables:\n")

	listTables(input, timeoutCtx, svc)
}

func TestCreatedTable(t *testing.T) {
	svc := initDynamoDBClient()
	attributeDefinitions := []*dynamodb.AttributeDefinition{
		{ // "target_id"_"share_by"
			AttributeName: aws.String("view_id"),
			AttributeType: aws.String("S"),
		},
		{ // timestamp:Millisecond
			AttributeName: aws.String("created_time"),
			AttributeType: aws.String("N"),
		},
		{ // timestamp:Millisecond
			AttributeName: aws.String("view_user_id"),
			AttributeType: aws.String("S"),
		},
	}

	keySchema := []*dynamodb.KeySchemaElement{
		{
			AttributeName: aws.String("view_id"),
			KeyType:       aws.String("HASH"),
		},
		{
			AttributeName: aws.String("created_time"),
			KeyType:       aws.String("RANGE"),
		},
	}

	globalSecondaryIndexes := []*dynamodb.GlobalSecondaryIndex{
		{
			IndexName: aws.String("gsi_view_user_id"),
			KeySchema: []*dynamodb.KeySchemaElement{
				{
					AttributeName: aws.String("view_user_id"),
					KeyType:       aws.String("HASH"),
				},
			},
			Projection: &dynamodb.Projection{
				ProjectionType: aws.String(dynamodb.ProjectionTypeAll),
			},
			ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
				ReadCapacityUnits:  aws.Int64(5),
				WriteCapacityUnits: aws.Int64(5),
			},
		},
	}

	localSecondaryIndexes := []*dynamodb.LocalSecondaryIndex{
		{
			IndexName: aws.String("lsi_view_id_view_user_id"),
			KeySchema: []*dynamodb.KeySchemaElement{
				{
					AttributeName: aws.String("view_id"),
					KeyType:       aws.String("HASH"),
				},
				{
					AttributeName: aws.String("view_user_id"),
					KeyType:       aws.String("RANGE"),
				},
			},
			Projection: &dynamodb.Projection{
				ProjectionType: aws.String(dynamodb.ProjectionTypeAll),
			},
		},
	}

	cratedTable(svc, "share_view", attributeDefinitions, keySchema, globalSecondaryIndexes, localSecondaryIndexes)
}

func TestDeleteTable(t *testing.T) {
	svc := initDynamoDBClient()
	deleteTable(svc, "cn-north-1-mdm-one-share-share-v1-dev")
}

func TestInsertItem(t *testing.T) {
	svc := initDynamoDBClient()

	item := ViewItem{
		//ViewId:            "037333af-5c3b-43f2-9434-d57b1796735a+1000395666",
		CreatedTime:       time.Now().UnixNano(),
		TargetType:        "BGC",
		ViewType:          "READ",
		ViewChannelType:   "WECHAT",
		ViewUserType:      "ALLINONE",
		ViewUserId:        "1000395688",
		ViewUserNickname:  "Lee",
		ViewUserAvatarUrl: "https://community-s3.marykay.com.cn/PROD/Community_UserManagement/avatar/1000395688/1491349533424.png",
		ViewUserPhone:     "13554539296",
	}

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		fmt.Println("Got error marshalling new movie item:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	createItemInput := &dynamodb.PutItemInput{
		ConditionExpression:         nil,
		ConditionalOperator:         nil,
		Expected:                    nil,
		ExpressionAttributeNames:    nil,
		ExpressionAttributeValues:   nil,
		Item:                        av,
		ReturnConsumedCapacity:      nil,
		ReturnItemCollectionMetrics: nil,
		ReturnValues:                nil,
		TableName:                   aws.String("share_view"),
	}

	insertItem(svc, createItemInput)
}

func TestGetItem(t *testing.T) {
	svc := initDynamoDBClient()
	input := &dynamodb.GetItemInput{
		TableName: aws.String(ViewTable),
		Key: map[string]*dynamodb.AttributeValue{
			"view_id": {
				S: aws.String("037333af-5c3b-43f2-9434-d57b1796735a+1000395688"),
			},
		},
	}

	getItem(svc, input)
}

func TestScanItem(t *testing.T) {
	svc := initDynamoDBClient()

	filter := expression.Name("view_id").Equal(expression.Value("037333af-5c3b-43f2-9434-d57b1796735a+1000395666"))

	expr, err := expression.NewBuilder().WithFilter(filter).Build()

	if err != nil {
		fmt.Println("Got error building expression:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		TableName:                 aws.String(ViewTable),
	}

	// Make the DynamoDB Query API call
	result, err := svc.Scan(params)

	if err != nil {
		fmt.Println("Query API call failed:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _, i := range result.Items {
		item := ViewItem{}

		err = dynamodbattribute.UnmarshalMap(i, &item)

		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println("view item: ", item)
		fmt.Println("view item targetType: ", item.TargetType)
		fmt.Println()
	}
}

func TestQueryItem(t *testing.T) {
	svc := initDynamoDBClient()

	params := &dynamodb.QueryInput{
		KeyConditionExpression: aws.String("#a1=:v1"),
		ExpressionAttributeNames: map[string]*string{
			"#a1": aws.String("target_id"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":v1": {
				S: aws.String("0c7559fb-391b-4184-90d7-766b80c422c3"),
			},
		},
		Limit:            aws.Int64(1),
		ScanIndexForward: aws.Bool(true),
		TableName:        aws.String(ShareTable),
		IndexName:        aws.String("g_index_target_id_created_time"),
	}

	// Make the DynamoDB Query API call
	result, err := svc.Query(params)

	if err != nil {
		fmt.Println("Query API call failed:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _, i := range result.Items {
		item := ShareItem{}

		err = dynamodbattribute.UnmarshalMap(i, &item)

		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println("share item: ", item)
	}
}

func TestUpdateItem(t *testing.T) {
	svc := initDynamoDBClient()

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			//":r": {
			//	N: aws.String(movieRating),
			//},
		},
		TableName: aws.String(ViewTable),
		Key:       map[string]*dynamodb.AttributeValue{
			//"Year": {
			//	N: aws.String(movieYear),
			//},
			//"Title": {
			//	S: aws.String(movieName),
			//},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set Rating = :r"),
	}

	updateItem(svc, input, ViewTable)
}

func TestDeleteItem(t *testing.T) {
	svc := initDynamoDBClient()

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"view_id": {
				S: aws.String("037333af-5c3b-43f2-9434-d57b1796735a+1000395666"),
			},
		},
	}

	_, err := svc.DeleteItem(input)
	if err != nil {
		fmt.Println("Got error calling DeleteItem")
		fmt.Println(err.Error())
		return
	}
}

func cratedTable(svc *dynamodb.DynamoDB, tableName string,
	attributeDefinitions []*dynamodb.AttributeDefinition,
	keySchema []*dynamodb.KeySchemaElement,
	globalSecondaryIndexes []*dynamodb.GlobalSecondaryIndex, localSecondaryIndexes []*dynamodb.LocalSecondaryIndex) {

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: attributeDefinitions,
		KeySchema:            keySchema,
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName:              aws.String(tableName),
		GlobalSecondaryIndexes: globalSecondaryIndexes,
		LocalSecondaryIndexes:  localSecondaryIndexes,
	}

	_, err := svc.CreateTable(input)
	if err != nil {
		fmt.Println("Got error calling CreateTable:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Created the table", tableName)
}

func deleteTable(svc *dynamodb.DynamoDB, tableName string) {
	input := &dynamodb.DeleteTableInput{TableName: aws.String(tableName)}

	_, err := svc.DeleteTable(input)

	if err != nil {
		log.Error(err)
	}
}

func insertItem(svc *dynamodb.DynamoDB, input *dynamodb.PutItemInput) {
	_, err := svc.PutItem(input)
	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func getItem(svc *dynamodb.DynamoDB, input *dynamodb.GetItemInput) {
	result, err := svc.GetItem(input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	item := ViewItem{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	log.Info(item)
}

func updateItem(svc *dynamodb.DynamoDB, input *dynamodb.UpdateItemInput, tableName string) {
	_, err := svc.UpdateItem(input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func listTables(input *dynamodb.ListTablesInput, timeoutCtx context.Context, svc *dynamodb.DynamoDB) {
	for {
		// Get the list of tables
		result, err := svc.ListTablesWithContext(timeoutCtx, input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case dynamodb.ErrCodeInternalServerError:
					fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				fmt.Println(err.Error())
			}
			return
		}

		for _, n := range result.TableNames {
			fmt.Println(*n)
		}

		// assign the last read tablename as the start for our next call to the ListTables function
		// the maximum number of table names returned in a call is 100 (default), which requires us to make
		// multiple calls to the ListTables function to retrieve all table names
		input.ExclusiveStartTableName = result.LastEvaluatedTableName

		if result.LastEvaluatedTableName == nil {
			break
		}
	}
}

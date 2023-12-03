package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	if err != nil {
		log.Fatal("Failed to create AWS session:", err)
	}

	createEc2(sess)
	createS3(sess)
	createDynamoDb(sess)
}

func createEc2(sess *session.Session) {
	ec2Svc := ec2.New(sess)
	ec2Instance, err := ec2Svc.RunInstances(&ec2.RunInstancesInput{
		ImageId:      aws.String("your_ami"),
		InstanceType: aws.String("your_instance_type"),
		MinCount:     aws.Int64(1),
		MaxCount:     aws.Int64(1),
	})
	if err != nil {
		log.Fatal("Failed to create EC2 instance:", err)
	}
	fmt.Println("EC2 Instance ID:", *ec2Instance.Instances[0].InstanceId)

}

func createS3(sess *session.Session) {
	s3Svc := s3.New(sess)
	bucketName := "your_bucket"
	_, err := s3Svc.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		log.Fatal("Failed to create S3 bucket:", err)
	}
	fmt.Println("S3 Bucket created:", bucketName)

}

func createDynamoDb(sess *session.Session) {
	dynamoDBSvc := dynamodb.New(sess)
	tableName := "YourTable"
	_, err := dynamoDBSvc.CreateTable(&dynamodb.CreateTableInput{
		TableName: aws.String(tableName),
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("YourPrimaryKey"),
				KeyType:       aws.String("HASH"),
			},
		},
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("YourPrimaryKey"),
				AttributeType: aws.String("S"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
	})
	if err != nil {
		log.Fatal("Failed to create DynamoDB table:", err)
	}
	fmt.Println("DynamoDB Table created:", tableName)
}

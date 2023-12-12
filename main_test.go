package main

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/s3"
)

func TestCreateEc2(t *testing.T) {
	mockSession := session.Must(session.NewSession(&aws.Config{}))

	mockAMI := "ami-12345678"
	mockInstanceType := "t2.micro"

	ec2Svc := &ec2.EC2{
		Client: ec2.New(mockSession),
	}

	err := createEc2(mockSession, ec2Svc, mockAMI, mockInstanceType)
	if err != nil {
		t.Errorf("Failed to create EC2 instance: %v", err)
	}
}

func TestCreateS3(t *testing.T) {
	mockSession := session.Must(session.NewSession(&aws.Config{}))

	mockBucketName := "test-bucket"

	s3Svc := &s3.S3{
		Client: s3.New(mockSession),
	}

	err := createS3(mockSession, s3Svc, mockBucketName)
	if err != nil {
		t.Errorf("Failed to create S3 bucket: %v", err)
	}
}

func TestCreateDynamoDb(t *testing.T) {
	mockSession := session.Must(session.NewSession(&aws.Config{}))

	mockTableName := "test-table"

	dynamoDBSvc := &dynamodb.DynamoDB{
		Client: dynamodb.New(mockSession),
	}

	err := createDynamoDb(mockSession, dynamoDBSvc, mockTableName)
	if err != nil {
		t.Errorf("Failed to create DynamoDB table: %v", err)
	}
}

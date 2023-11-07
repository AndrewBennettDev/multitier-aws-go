# Multi-tier AWS Starter App

### Overview

This is a quick and easy way to launch and EC2 instance, an S3 bucket, and a DynamoDB table.
As of the current interation this only sets the bare minimum and would need additional
configuration for maximum utility.

### Usage

Before running this you would need to have your `aws configure` setup with your AWS
credentials, and likely you will need to have correct permissions set within the AWS
console to allow the creation of any or all of these resources. Next, clone the repo, cd
into the project folder and run `go get github.com/aws/aws-sdk-go/` in install the AWS SDK
for Go. Finally, you can update anything like instance name and type, table names, number
of instances, etc. Then it's ready for a `go run main.go` and if everything is configured
properly you should see your handy dandy multi-tier AWS multi-tier project!

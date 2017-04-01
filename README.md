# AWS Role Assumer for Go

Some AWS users prefer bit different access management based on AssumeRole API method:
http://docs.aws.amazon.com/STS/latest/APIReference/API_AssumeRole.html

This little example shows proper flow that implements the role assumption access model. 

Dependency:
- `github.com/aws/aws-sdk-go`

package client

import (
	"time"

	"github.com/Bplotka/aws-role-assumer-go"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

const (
	myCustomArn = "arn:aws:iam::<account_number>:role/service/<service_name>"
	expiration  = 999 * time.Second
)

// NewS3Client is an example S3Client constructor making a new S3Client with assume role permission policy.
func NewS3Client(accessKey string, secretKey string, region string) (s3iface.S3API, error) {
	creds := credentials.NewStaticCredentials(accessKey, secretKey, "")

	cfg, err := assumer.AssumeRoleConfig(myCustomArn, creds, region, expiration)
	if err != nil {
		return nil, err
	}

	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}
	return s3.New(sess, cfg), nil
}

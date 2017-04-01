package assumer

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

const (
	maxExpiration = 3600 * time.Second
	minExpiration = 900 * time.Second
)

func validateExpiration(expiration time.Duration) error {
	if expiration > maxExpiration {
		return fmt.Errorf("given expiration %v is too long, setting to max value: %v",
			expiration, maxExpiration)

	}

	if expiration < minExpiration {
		return fmt.Errorf("given expiration %v is too short, setting to min value: %v",
			expiration, minExpiration)

	}

	return nil
}

// AssumeRoleConfig provides config that can be passed to construct any AWS client for temporary credentials model.
// This config specifies special `AssumeRoleProvider` (STS) that will be automatically assuming role when credentials will expire.
// NOTE: See http://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_use-resources.html#using-temp-creds-sdk
// for details.
func AssumeRoleConfig(arn string, creds *credentials.Credentials, region string, expiration time.Duration) (*aws.Config, error) {
	err := validateExpiration(expiration)
	if err != nil {
		return nil, err
	}

	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}

	stsConfig := aws.NewConfig().WithCredentials(creds).WithRegion(region)
	assumeRoleCreds := stscreds.NewCredentialsWithClient(sts.New(sess, stsConfig), arn, func(p *stscreds.AssumeRoleProvider) {
		p.Duration = expiration
	})
	return aws.NewConfig().WithCredentials(assumeRoleCreds).WithRegion(region), nil
}

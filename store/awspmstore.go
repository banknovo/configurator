package store

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/banknovo/configurator/config"
)

// ensure AWSPMStore confirms to Store interface
var _ Store = &AWSPMStore{}

// AWSPMStore is used to fetch keys from AWS Parameters Store
type AWSPMStore struct {
	svc *ssm.SSM
}

// NewAWSPMStore creates a new store which used AWS Parameters Store
func NewAWSPMStore() (*AWSPMStore, error) {
	awsConfig := &aws.Config{Region: aws.String("us-east-1")}
	sess, err := session.NewSessionWithOptions(session.Options{
		Config:            *awsConfig,
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		return nil, err
	}
	awspmstore := &AWSPMStore{
		svc: ssm.New(sess, awsConfig),
	}
	return awspmstore, nil
}

// FetchAll fetches all parameters by path from Parameters Store
func (s *AWSPMStore) FetchAll(path string) ([]*config.Config, error) {
	awsParams, err := s.getParametersByPath(path)
	if err != nil {
		return nil, err
	}
	return convert(awsParams), nil
}

// covert the parameters from SSM data type to native
func convert(awsParams []*ssm.Parameter) []*config.Config {
	params := make([]*config.Config, len(awsParams))
	for i, awsParam := range awsParams {
		param := &config.Config{
			Key:   *awsParam.Name,
			Value: *awsParam.Value,
		}
		params[i] = param
		i++
	}
	return params
}

// getParametersByPath does API pagination and returns all parameters
func (s *AWSPMStore) getParametersByPath(path string) ([]*ssm.Parameter, error) {
	params := make([]*ssm.Parameter, 0)
	var token *string
	for {
		input := &ssm.GetParametersByPathInput{
			Path:           &path,
			WithDecryption: aws.Bool(true),
			Recursive:      aws.Bool(true),
			MaxResults:     aws.Int64(10),
		}
		if token != nil {
			input.NextToken = token
		}
		output, err := s.svc.GetParametersByPath(input)
		if err != nil {
			return nil, err
		}
		params = append(params, output.Parameters...)
		token = output.NextToken
		if token == nil {
			break
		}
	}
	return params, nil
}

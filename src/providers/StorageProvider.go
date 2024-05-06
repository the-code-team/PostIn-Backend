package providers

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"log"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	storageOnce sync.Once
	instance    *s3.Client
)

func GetStorageClient() *s3.Client {
	storageOnce.Do(func() {
		// Load the SDK's configuration from the environment only
		// We need to load also the endpoint
		cfg, err := config.LoadDefaultConfig(
			context.Background(),
			config.WithRegion("eu-west-1"),
			config.WithEndpointResolver(s3.EndpointResolverFunc(
				func(region string, options s3.EndpointResolverOptions) (aws.Endpoint, error) {
					return aws.Endpoint{
						URL: os.Getenv("S3_ENDPOINT"),
					}, nil
				},
			)),
		)

		if err != nil {
			log.Fatal(err)
		}

		instance = s3.NewFromConfig(cfg)
	})

	return instance
}

func InitStorageClient() {
	GetStorageClient()
}

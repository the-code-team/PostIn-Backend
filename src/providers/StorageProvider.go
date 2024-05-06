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
		// Load the SDK's configuration from the environment
		cfg, err := config.LoadDefaultConfig(
			context.Background(),
			config.WithRegion("eu-west-1"),
		)

		if err != nil {
			log.Fatal(err)
		}

		// Create an Amazon S3 service client
		instance = s3.NewFromConfig(cfg, func(options *s3.Options) {
			options.BaseEndpoint = aws.String(os.Getenv("S3_ENDPOINT"))
		})
	})

	// Return the instance
	return instance
}

func InitStorageClient() {
	GetStorageClient()
}

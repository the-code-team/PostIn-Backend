package providers

import (
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	storageOnce     sync.Once
	instance 		*s3.S3
)

func GetStorageClient() *s3.S3 {
	storageOnce.Do(func() {
		// Create a new AWS session
		sess := session.Must(session.NewSession(&aws.Config{
			Region: aws.String("us-west-2"), // Replace with your desired region
		}))

		// Create a new S3 client
		instance = s3.New(sess)
	})

	return instance
}
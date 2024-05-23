package use_cases

import (
	"bufio"
	"context"
	"epsa.upv.es/postin_backend/src/models"
	"epsa.upv.es/postin_backend/src/modules/events_mod/domain/commands"
	"epsa.upv.es/postin_backend/src/providers"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gogolfing/cbus"
	"os"
)

func PublishPhotoHeaderUseCase() {
	bus := providers.GetCommandBus()
	bus.Handle(&commands.PublishPhotoHeaderCommand{}, cbus.HandlerFunc(PublishPhotoHeaderHandler))
}

func PublishPhotoHeaderHandler(ctx context.Context, command cbus.Command) (interface{}, error) {
	// Get the providers
	db := providers.GetDatabase()
	cache := providers.GetQueryCacheProvider()
	s3Client := providers.GetStorageClient()

	// Upload the photo header into S3
	event := &models.Event{}

	// Get the event from the database
	query := cache.Wrap(
		db.Model(&models.Event{}).
			Where("EventId = ?", command.(*commands.PublishPhotoHeaderCommand).EventId),
	)

	// Return the event into the result variable
	query = query.First(&event)

	if query.Error != nil {
		return nil, query.Error
	}

	// Upload the photo header into the storage with Amazon S3 API
	for i, photo := range command.(*commands.PublishPhotoHeaderCommand).Photos {
		_, err := s3Client.PutObject(context.Background(), &s3.PutObjectInput{
			Bucket: aws.String(os.Getenv("S3_BUCKET")),
			Key:    aws.String(`events/` + event.EventId.String() + `/photo/` + string(rune(i))),
			Body:   bufio.NewReader(&photo),
		})

		if err != nil {
			return nil, err
		}
	}

	// The photo header has been uploaded successfully
	return nil, nil
}

package usecases

import (
	"bufio"
	"context"
	"epsa.upv.es/postin_backend/src/models"
	"epsa.upv.es/postin_backend/src/modules/profile_mod/domain/commands"
	"epsa.upv.es/postin_backend/src/providers"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gogolfing/cbus"
	"os"
)

func UpdatePhotosUseCase() {
	bus := providers.GetCommandBus()
	bus.Handle(&commands.UpdatePhotosCommand{}, cbus.HandlerFunc(UpdatePhotoHandler))
}

func UpdatePhotoHandler(ctx context.Context, command cbus.Command) (interface{}, error) {
	// Get the providers
	db := providers.GetDatabase()
	s3Client := providers.GetStorageClient()

	// Get the profile
	profile := &models.Profile{}
	query := db.Model(&models.Profile{})
	query = query.Where("Email = ?", command.(*commands.UpdatePhotosCommand).Email)
	query = query.First(&profile)

	if query.Error != nil {
		return nil, query.Error
	}

	// Upload the photos into the storage with Amazon S3 API
	for i, photo := range command.(*commands.UpdatePhotosCommand).Photos {
		_, err := s3Client.PutObject(context.Background(), &s3.PutObjectInput{
			Bucket: aws.String(os.Getenv("S3_BUCKET")),
			Key:    aws.String(`profiles/` + profile.UserId + `/photo/` + string(rune(i))),
			Body:   bufio.NewReader(&photo),
		})

		if err != nil {
			return nil, err
		}
	}

	// The photos have been uploaded successfully
	return nil, nil
}

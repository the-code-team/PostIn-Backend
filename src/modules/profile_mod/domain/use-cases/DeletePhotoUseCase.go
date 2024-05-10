package usecases

import (
	"context"
	"epsa.upv.es/postin_backend/src/models"
	"epsa.upv.es/postin_backend/src/modules/profile_mod/domain/commands"
	"epsa.upv.es/postin_backend/src/providers"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gogolfing/cbus"
	"os"
)

func DeletePhotoUseCase() {
	bus := providers.GetCommandBus()
	bus.Handle(&commands.DeletePhotoCommand{}, cbus.HandlerFunc(DeletePhotoHandler))
}

func DeletePhotoHandler(ctx context.Context, command cbus.Command) (interface{}, error) {
	// Get the providers
	db := providers.GetDatabase()
	cache := providers.GetQueryCacheProvider()
	s3Client := providers.GetStorageClient()

	// Get the profile
	profile := &models.Profile{}

	// Get the profile from the database
	query := cache.Wrap(
		db.Model(&models.Profile{}).
			Where("Email = ?", command.(*commands.DeletePhotoCommand).Email),
	)

	// Return the profile into the result variable
	query = query.First(&profile)

	if query.Error != nil {
		return nil, query.Error
	}

	// Delete the photo from the storage with Amazon S3 API
	_, err := s3Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET")),
		Key:    aws.String(`profiles/` + profile.UserId + `/photo/` + string(rune(command.(*commands.DeletePhotoCommand).PhotoId))),
	})

	if err != nil {
		return nil, err
	}

	// The photo has been deleted successfully
	return nil, nil
}

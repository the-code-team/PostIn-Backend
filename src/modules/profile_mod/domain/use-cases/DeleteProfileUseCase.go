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

func DeleteProfileUseCase() {
	bus := providers.GetCommandBus()
	bus.Handle(&commands.DeleteProfileCommand{}, cbus.HandlerFunc(DeleteProfileHandler))
}

func DeleteProfileHandler(ctx context.Context, command cbus.Command) (interface{}, error) {
	// Get the providers
	db := providers.GetDatabase()
	s3Client := providers.GetStorageClient()

	// Get the profile
	profile := &models.Profile{}
	query := db.Model(&models.Profile{})
	query = query.Where("Email = ?", command.(*commands.DeleteProfileCommand).Email)
	query = query.First(&profile)

	if query.Error != nil {
		return nil, query.Error
	}

	// Delete the profile folder from S3 storage
	_, err := s3Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET")),
		Key:    aws.String(`profiles/` + profile.UserId + `/`),
	})

	if err != nil {
		return nil, err
	}

	// Delete the profile from the database
	query = db.Model(&models.Profile{}).Delete(&profile)

	if query.Error != nil {
		return nil, query.Error
	}

	// The profile has been deleted successfully
	return nil, nil
}

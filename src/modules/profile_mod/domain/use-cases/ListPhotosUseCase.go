package usecases

import (
	"context"
	"epsa.upv.es/postin_backend/src/models"
	domainmodels "epsa.upv.es/postin_backend/src/modules/profile_mod/domain/models"
	"epsa.upv.es/postin_backend/src/modules/profile_mod/domain/queries"
	"epsa.upv.es/postin_backend/src/providers"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gogolfing/cbus"
	"os"
)

func ListPhotosUseCase() {
	bus := providers.GetCommandBus()
	bus.Handle(&queries.ListPhotosQuery{}, cbus.HandlerFunc(ListPhotosHandler))
}

func ListPhotosHandler(ctx context.Context, command cbus.Command) (interface{}, error) {
	// Get the providers
	db := providers.GetDatabase()
	s3Client := providers.GetStorageClient()
	s3Signer := s3.NewPresignClient(s3Client)

	// Get the profile
	profile := &models.Profile{}
	query := db.Model(&models.Profile{})
	query = query.Where("Email = ?", command.(*queries.ListPhotosQuery).Email)
	query = query.First(&profile)

	if query.Error != nil {
		return nil, query.Error
	}

	// List the photos from the storage with Amazon S3 API
	var photos []*domainmodels.ProfilePhoto

	output, err := s3Client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(os.Getenv("S3_BUCKET")),
		Prefix: aws.String(`profiles/` + profile.UserId + `/photo/`),
	})

	// Sign the photo URLs before returning them
	for _, item := range output.Contents {
		// Sign the photo URL
		photoUrl, err := s3Signer.PresignGetObject(ctx, &s3.GetObjectInput{
			Bucket: aws.String(os.Getenv("S3_BUCKET")),
			Key:    item.Key,
		})

		if err != nil {
			return nil, err
		}

		// Append the photo to the list
		photos = append(photos, &domainmodels.ProfilePhoto{
			PhotoUri:  photoUrl.URL,
			UpdatedAt: item.LastModified,
		})
	}

	if err != nil {
		return nil, err
	}

	// Return the photos
	return photos, nil
}

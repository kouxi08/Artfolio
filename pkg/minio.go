package pkg

import (
	"context"
	"os"

	"github.com/kouxi08/Artfolio/utils"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewMinioClient() (*minio.Client, error) {
	utils.Env()
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKey := os.Getenv("ACCESS_KEY")
	secretAccessKey := os.Getenv("SECRETACCESS_KEY")
	useSSL := true

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}
	return minioClient, nil

}

func MakeBucket(bucketName string) (string, error) {
	minioClient, _ := NewMinioClient()

	err := minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{Region: "us-east-1", ObjectLocking: true})
	if err != nil {
		return "", err
	}
	return "Successfully created ", nil
}

func RemoveBucket(bucketName string) (string, error) {
	minioClient, _ := NewMinioClient()

	err := minioClient.RemoveBucket(context.Background(), bucketName)
	if err != nil {
		return "", err
	}
	return "Successfully deleted", nil
}

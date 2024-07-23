package s3

import (
	"github.com/caitlinelfring/go-env-default"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var Client *minio.Client

func Open() {
	endpoint := env.GetDefault("S3_ENDPOINT", "")
	accessKeyID := env.GetDefault("S3_ACCESS_KEY_ID", "")
	secretAccessKey := env.GetDefault("S3_SECRET_ACCESS_KEY", "")
	secure := env.GetBoolDefault("S3_SECURE", true)

	var err error

	Client, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: secure,
	})
	if err != nil {
		panic(err)
	}
}

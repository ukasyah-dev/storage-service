package controller

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/minio/minio-go/v7"
	"github.com/ukasyah-dev/common/errors"
	"github.com/ukasyah-dev/common/id"
	"github.com/ukasyah-dev/common/validator"
	"github.com/ukasyah-dev/storage-service/model"
	"github.com/ukasyah-dev/storage-service/s3"
)

func CreateFile(ctx context.Context, req *model.CreateFileRequest) (*model.File, error) {
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	file, err := req.File.Open()
	if err != nil {
		log.Errorf("Failed to open file: %s", err)
		return nil, errors.Internal()
	}
	defer file.Close()

	_, err = s3.Client.PutObject(ctx, "public", req.File.Filename, file, int64(req.File.Size), minio.PutObjectOptions{})
	if err != nil {
		log.Errorf("Failed to upload file: %s", err)
		return nil, errors.Internal()
	}

	f := &model.File{
		ID:          id.New(),
		Name:        req.File.Filename,
		ContentType: req.File.Header.Get("Content-Type"),
	}

	return f, nil
}

func GetFiles(ctx context.Context, req *model.GetFilesRequest) (*model.GetFilesResponse, error) {
	return nil, errors.Internal("Not implemented")
}

func GetFile(ctx context.Context, req *model.GetFileRequest) (*model.File, error) {
	return nil, errors.Internal("Not implemented")
}

func UpdateFile(ctx context.Context, req *model.UpdateFileRequest) (*model.File, error) {
	return nil, errors.Internal("Not implemented")
}

func DeleteFile(ctx context.Context, req *model.DeleteFileRequest) (*model.File, error) {
	return nil, errors.Internal("Not implemented")
}

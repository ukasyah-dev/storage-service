package controller

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/minio/minio-go/v7"
	"github.com/samber/lo"
	"github.com/ukasyah-dev/common/constant"
	"github.com/ukasyah-dev/common/errors"
	"github.com/ukasyah-dev/common/id"
	"github.com/ukasyah-dev/common/validator"
	"github.com/ukasyah-dev/storage-service/db"
	"github.com/ukasyah-dev/storage-service/model"
	"github.com/ukasyah-dev/storage-service/s3"
)

func CreateFile(ctx context.Context, req *model.CreateFileRequest) (*model.File, error) {
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	file := &model.File{
		ID:          id.New(),
		Name:        req.File.Filename,
		ContentType: req.File.Header.Get("Content-Type"),
		Type:        req.Type,
	}

	tags := map[string]string{}

	// Attach user ID to tags (optional)
	if userID, ok := ctx.Value(constant.UserID).(string); ok && userID != "" {
		tags["user_id"] = userID
	}

	file.Tags = lo.MapToSlice(tags, func(k, v string) *model.Tag {
		return &model.Tag{Key: k, Value: v}
	})

	// Store file into database
	err := db.DB.WithContext(ctx).Create(file).Error
	if err != nil {
		log.Errorf("Failed to store file into database: %s", err)
		return nil, errors.Internal()
	}

	// Add name to tags
	tags["name"] = file.Name

	f, err := req.File.Open()
	if err != nil {
		log.Errorf("Failed to open file: %s", err)
		return nil, errors.Internal()
	}
	defer f.Close()

	// Upload file to S3
	_, err = s3.Client.PutObject(ctx, file.Type, file.ID, f, int64(req.File.Size), minio.PutObjectOptions{
		ContentType: file.ContentType,
		UserTags:    tags,
	})
	if err != nil {
		log.Errorf("Failed to upload file: %s", err)
		return nil, errors.Internal()
	}

	return file, nil
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

package model

import (
	"mime/multipart"
	"time"

	commonModel "github.com/ukasyah-dev/common/model"
)

type File struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	ContentType string    `json:"contentType"`
	Type        string    `json:"type"`
	Tags        []*Tag    `gorm:"many2many:file_tags;" json:"tags"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type CreateFileRequest struct {
	File *multipart.FileHeader `formData:"file" validate:"required"`
	Type string                `formData:"type" validate:"required,oneof=public private" example:"public"`
}

type GetFilesRequest struct {
	commonModel.PaginationRequest
}

type GetFilesResponse struct {
	commonModel.PaginationResponse
	Data []*File `json:"data"`
}

type GetFileRequest struct {
	ID string `params:"fileId" path:"fileId" validate:"required"`
}

type UpdateFileRequest struct {
	ID string `params:"fileId" path:"fileId" validate:"required"`
}

type DeleteFileRequest struct {
	ID string `params:"fileId" path:"fileId" validate:"required"`
}

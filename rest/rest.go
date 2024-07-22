package rest

import (
	"net/http"
	"os"

	"github.com/swaggest/openapi-go/openapi31"
	"github.com/ukasyah-dev/common/rest/handler"
	"github.com/ukasyah-dev/common/rest/server"
	"github.com/ukasyah-dev/storage-service/controller"
)

var Server *server.Server

func init() {
	description := "File management."
	spec := openapi31.Spec{
		Openapi: "3.1.0",
		Info: openapi31.Info{
			Title:       "Storage Service",
			Version:     "0.0.1",
			Description: &description,
		},
		Servers: []openapi31.Server{
			{URL: os.Getenv("OPENAPI_SERVER_URL")},
		},
	}

	// Create new server
	Server = server.New(server.Config{
		OpenAPI: server.OpenAPI{Spec: &spec},
	})

	handler.AddHealthCheck(Server)

	// File
	handler.Add(Server, http.MethodPost, "/files", controller.CreateFile, handler.Config{
		Summary:     "Create new file",
		Description: "Create new file",
		Tags:        []string{"File"},
	})
	handler.Add(Server, http.MethodGet, "/files", controller.GetFiles, handler.Config{
		Summary:     "Get all files",
		Description: "Get all files",
		Tags:        []string{"File"},
	})
	handler.Add(Server, http.MethodGet, "/files/:fileId", controller.GetFile, handler.Config{
		Summary:     "Get file",
		Description: "Get file",
		Tags:        []string{"File"},
	})
	handler.Add(Server, http.MethodPut, "/files/:fileId", controller.UpdateFile, handler.Config{
		Summary:     "Update file",
		Description: "Update file",
		Tags:        []string{"File"},
	})
	handler.Add(Server, http.MethodDelete, "/files/:fileId", controller.DeleteFile, handler.Config{
		Summary:     "Delete file",
		Description: "Delete file",
		Tags:        []string{"File"},
	})
}

package rest

import (
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/swaggest/openapi-go/openapi31"
	commonAuth "github.com/ukasyah-dev/common/auth"
	"github.com/ukasyah-dev/common/rest/handler"
	"github.com/ukasyah-dev/common/rest/middleware"
	"github.com/ukasyah-dev/common/rest/server"
	"github.com/ukasyah-dev/storage-service/controller"
	"github.com/ukasyah-dev/storage-service/model"
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

	// Parse JWT public key
	jwtPublicKey, err := commonAuth.ParsePublicKeyFromBase64(os.Getenv("BASE64_JWT_PUBLIC_KEY"))
	if err != nil {
		panic(err)
	}

	// Create new server
	Server = server.New(server.Config{
		OpenAPI:      server.OpenAPI{Spec: &spec},
		JWTPublicKey: jwtPublicKey,
	})

	handler.AddHealthCheck(Server)

	// Create file
	Server.FiberApp.Add(http.MethodPost, "/files", middleware.Authenticate(jwtPublicKey), func(c *fiber.Ctx) error {
		file, _ := c.FormFile("file")

		result, err := controller.CreateFile(c.Context(), &model.CreateFileRequest{
			File: file,
			Type: c.FormValue("type"),
		})
		if err != nil {
			return err
		}

		return c.JSON(result)
	})
	op, _ := Server.Config.OpenAPI.Reflector.NewOperationContext(http.MethodPost, "/files")
	op.SetID("createFile")
	op.SetSummary("Create file")
	op.SetDescription("Create file")
	op.SetTags("File")
	op.AddReqStructure(&model.CreateFileRequest{})
	op.AddRespStructure(&model.File{})
	op.AddSecurity("Bearer auth")
	if err := Server.Config.OpenAPI.Reflector.AddOperation(op); err != nil {
		panic(err)
	}

	// File
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

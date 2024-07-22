package rest

import (
	"os"

	"github.com/swaggest/openapi-go/openapi31"
	"github.com/ukasyah-dev/common/rest/handler"
	"github.com/ukasyah-dev/common/rest/server"
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
}

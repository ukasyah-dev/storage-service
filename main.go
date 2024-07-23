package main

import (
	"context"

	"github.com/appleboy/graceful"
	"github.com/caitlinelfring/go-env-default"
	"github.com/ukasyah-dev/storage-service/db"
	"github.com/ukasyah-dev/storage-service/rest"
	"github.com/ukasyah-dev/storage-service/s3"
)

var port = env.GetIntDefault("HTTP_PORT", 3000)

func init() {
	db.Open()
	s3.Open()
}

func main() {
	m := graceful.NewManager()

	m.AddRunningJob(func(ctx context.Context) error {
		return rest.Server.Start(port)
	})

	m.AddShutdownJob(func() error {
		return rest.Server.Shutdown()
	})

	m.AddShutdownJob(func() error {
		return db.Close()
	})

	<-m.Done()
}

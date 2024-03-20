package main

import (
	"context"
	"flag"
	"log"

	"github.com/sarastee/LamodaTest/internal/app"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", ".env", "path to config file")
	flag.Parse()
}

func main() {
	ctx := context.Background()

	application, err := app.NewApp(ctx, configPath)
	if err != nil {
		log.Fatalf("failure while initialize app: %s", err.Error())
	}

	if err := application.Run(); err != nil {
		log.Fatalf("failure while running app: %s", err.Error())
	}
}

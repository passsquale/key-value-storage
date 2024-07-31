package main

import (
	"context"
	"github.com/passsquale/key-value-storage/internal/configuration"
	"github.com/passsquale/key-value-storage/internal/initialization"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	ConfigFileName = os.Getenv("CONFIG_FILE_NAME")
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg := &configuration.Config{}
	if ConfigFileName != "" {
		var err error
		cfg, err = configuration.Load(ConfigFileName)
		if err != nil {
			log.Fatal(err)
		}
	}

	initializer, err := initialization.NewInitializer(cfg)
	if err != nil {
		log.Fatal(err)
	}

	if err = initializer.StartDatabase(ctx); err != nil {
		log.Fatal(err)
	}
}

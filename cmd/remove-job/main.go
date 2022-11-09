package main

import (
	_ "github.com/sfomuseum/go-offline-gocloud"
)

import (
	"context"
	"github.com/sfomuseum/go-offline/app/job/remove"
	"log"
)

func main() {

	ctx := context.Background()
	logger := log.Default()

	err := remove.Run(ctx, logger)

	if err != nil {
		logger.Fatalf("Failed to add job, %v", err)
	}
}

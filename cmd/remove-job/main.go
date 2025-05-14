package main

import (
	"context"
	"log"

	_ "github.com/sfomuseum/go-offline-gocloud"

	"github.com/sfomuseum/go-offline/app/job/remove"
)

func main() {

	ctx := context.Background()
	err := remove.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to remove job, %v", err)
	}
}

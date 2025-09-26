package main

import (
	"context"
	"log"

	_ "github.com/sfomuseum/go-offline-gocloud"

	"github.com/sfomuseum/go-offline/app/job/get"
)

func main() {

	ctx := context.Background()
	err := get.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to get job, %v", err)
	}
}

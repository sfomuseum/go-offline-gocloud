package main

import (
	"context"
	"log"

	_ "github.com/sfomuseum/go-offline-gocloud"

	"github.com/sfomuseum/go-offline/app/server"
)

func main() {

	ctx := context.Background()
	err := server.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to run job server, %v", err)
	}
}

package main

import (
	"context"
	"log"

	_ "github.com/sfomuseum/go-offline-gocloud"

	"github.com/sfomuseum/go-offline/app/job/add"
)

func main() {

	ctx := context.Background()
	logger := log.Default()

	/*
		fs := add.DefaultFlagSet()

		fs.VisitAll(func(f *flag.Flag){

			switch f.Name {
			case "database-uri":
				f.DefValue = "awsdynamodb://offlinejobs?partition_key=Id&local=true"
			default:
				//
			}
		})

		err := add.RunWithFlagSet(ctx, fs, logger)
	*/

	err := add.Run(ctx, logger)

	if err != nil {
		logger.Fatalf("Failed to add job, %v", err)
	}
}

package gocloud

import (
	"context"
	"fmt"
	"io"
	"time"

	_ "gocloud.dev/docstore/memdocstore"

	aa_docstore "github.com/aaronland/gocloud/docstore"
	"github.com/sfomuseum/go-offline"
	"gocloud.dev/docstore"
)

type DocstoreDatabase struct {
	offline.Database
	collection *docstore.Collection
}

func init() {

	ctx := context.Background()

	// See below
	err := offline.RegisterDatabase(ctx, "awsdynamodb", NewDocstoreDatabase)

	if err != nil {
		panic(err)
	}

	for _, scheme := range docstore.DefaultURLMux().CollectionSchemes() {

		err = offline.RegisterDatabase(ctx, scheme, NewDocstoreDatabase)

		if err != nil {
			panic(err)
		}
	}
}

func NewDocstoreDatabase(ctx context.Context, uri string) (offline.Database, error) {

	col, err := aa_docstore.OpenCollection(ctx, uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to create collection, %w", err)
	}

	db := &DocstoreDatabase{
		collection: col,
	}

	return db, nil
}

func (db *DocstoreDatabase) AddJob(ctx context.Context, job *offline.Job) error {

	err := db.collection.Put(ctx, job)

	if err != nil {
		return fmt.Errorf("Failed to store job, %w", err)
	}

	return nil
}

func (db *DocstoreDatabase) GetJob(ctx context.Context, id int64) (*offline.Job, error) {

	q := db.collection.Query()
	q = q.Where("Id", "=", id)

	iter := q.Get(ctx)
	defer iter.Stop()

	var job offline.Job
	err := iter.Next(ctx, &job)

	if err != nil {

		if err == io.EOF {
			return nil, fmt.Errorf("Not found")
		}

		return nil, fmt.Errorf("Failed to retrieve next item in iterator, %w", err)
	}

	return &job, nil
}

func (db *DocstoreDatabase) UpdateJob(ctx context.Context, job *offline.Job) error {

	now := time.Now()
	ts := now.Unix()

	job.LastModified = ts

	return db.collection.Replace(ctx, job)
}

func (db *DocstoreDatabase) RemoveJob(ctx context.Context, job *offline.Job) error {

	return db.collection.Delete(ctx, job)
}

func (db *DocstoreDatabase) ListJobs(ctx context.Context, cb offline.ListJobsCallback) error {

	q := db.collection.Query()

	iter := q.Get(ctx)
	defer iter.Stop()

	for {

		select {
		case <-ctx.Done():
			return nil
		default:
			// pass
		}

		var job offline.Job
		err := iter.Next(ctx, &job)

		if err == io.EOF {
			break
		} else if err != nil {
			return fmt.Errorf("Failed to iterate jobs, %w", err)
		} else {
			//
		}

		err = cb(ctx, &job)

		if err != nil {
			return fmt.Errorf("Callback failed for job %d", job.Id)
		}
	}

	return nil
}

func (db *DocstoreDatabase) PruneJobs(ctx context.Context, status offline.Status, lastmodified int64) error {

	q := db.collection.Query()

	q = q.Where("Status", "=", status)
	q = q.Where("LastModified", "<=", lastmodified)

	iter := q.Get(ctx)
	defer iter.Stop()

	for {

		select {
		case <-ctx.Done():
			return nil
		default:
			// pass
		}

		var job offline.Job
		err := iter.Next(ctx, &job)

		if err == io.EOF {
			break
		} else if err != nil {
			return fmt.Errorf("Failed to iterate jobs, %w", err)
		} else {
			//
		}

		err = db.RemoveJob(ctx, &job)

		if err != nil {
			return fmt.Errorf("Failed to remove job '%d', %v", job.Id, err)
		}
	}

	return nil

}

func (db *DocstoreDatabase) Close(ctx context.Context) error {
	return db.collection.Close()
}

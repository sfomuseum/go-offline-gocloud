package gocloud

import (
	_ "gocloud.dev/docstore/memdocstore"
)

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"time"

	aa_dynamodb "github.com/aaronland/go-aws-dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/sfomuseum/go-offline"
	"gocloud.dev/docstore"
	"gocloud.dev/docstore/awsdynamodb"	
)

type DocstoreDatabase struct {
	offline.Database
	collection *docstore.Collection
}

func init() {

	ctx := context.Background()

	// See below
	offline.RegisterDatabase(ctx, "awsdynamodb", NewDocstoreDatabase)

	for _, scheme := range docstore.DefaultURLMux().CollectionSchemes() {
		offline.RegisterDatabase(ctx, scheme, NewDocstoreDatabase)
	}
}

func NewDocstoreDatabase(ctx context.Context, uri string) (offline.Database, error) {

	// START OF put me in a package or something

	db_u, err := url.Parse(uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse URI, %w", err)
	}

	var col *docstore.Collection

	if db_u.Scheme == "awsdynamodb" {

		sess, err := aa_dynamodb.NewSessionWithURI(ctx, uri)

		if err != nil {
			return nil, fmt.Errorf("Failed to create session, %w", err)
		}

		// Connect local dynamodb using Golang
		// https://gist.github.com/Tamal/02776c3e2db7eec73c001225ff52e827
		// https://gocloud.dev/howto/docstore/#dynamodb-ctor

		table := db_u.Host

		db_q := db_u.Query()
		partition_key := db_q.Get("partition_key")

		opts := &awsdynamodb.Options{
			AllowScans: true,
			// RunQueryFallback: fallback_func,
		}

		// END OF necessary for order by created/lastupdate dates

		db, err := awsdynamodb.OpenCollection(dynamodb.New(sess), table, partition_key, "", opts)

		if err != nil {
			return nil, fmt.Errorf("Failed to open collection, %w", err)
		}

		col = db
	} else {

		db, err := docstore.OpenCollection(ctx, uri)

		if err != nil {
			return nil, fmt.Errorf("Failed to create collection, %w", err)
		}

		col = db
	}

	// END OF put me in a package or something

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

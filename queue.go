package gocloud

import (
	_ "gocloud.dev/pubsub/awssnssqs"
	_ "gocloud.dev/pubsub/mempubsub"
)

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/sfomuseum/go-offline"
	"github.com/sfomuseum/go-pubsub/publisher"
)

type PubSubQueue struct {
	offline.Queue
	publisher publisher.Publisher
}

func init() {

	ctx := context.Background()

	publisher.RegisterGoCloudPublishers(ctx)

	for _, uri := range publisher.PublisherSchemes() {
		scheme := strings.Replace(uri, "://", "", 1)
		offline.RegisterQueue(ctx, scheme, NewPubSubQueue)
	}
}

func NewPubSubQueue(ctx context.Context, uri string) (offline.Queue, error) {

	p, err := publisher.NewPublisher(ctx, uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to create publisher, %w", err)
	}

	q := &PubSubQueue{
		publisher: p,
	}

	return q, nil
}

func (q *PubSubQueue) ScheduleJob(ctx context.Context, job_id int64) error {

	str_id := strconv.FormatInt(job_id, 10)

	err := q.publisher.Publish(ctx, str_id)

	if err != nil {
		return fmt.Errorf("Failed to send message, %w", err)
	}

	return nil
}

func (q *PubSubQueue) Close(ctx context.Context) error {

	err := q.publisher.Close()

	if err != nil {
		return fmt.Errorf("Failed to close publisher, %w", err)
	}

	return nil
}

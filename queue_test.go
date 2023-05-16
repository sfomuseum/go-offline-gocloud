package gocloud

import (
	"context"
	"testing"

	"github.com/sfomuseum/go-offline"	
)

func TestQueueJob(t *testing.T) {

	ctx := context.Background()

	q_uri := "mem://testing"

	q, err := offline.NewQueue(ctx, q_uri)

	if err != nil {
		t.Fatalf("Failed to create new queue, %v", err)
	}

	job_id, err := offline.NewJobId(ctx)

	if err != nil {
		t.Fatalf("Failed to create new job ID, %v", err)
	}

	err = q.ScheduleJob(ctx, job_id)

	if err != nil {
		t.Fatalf("Failed to schedule job, %v", err)
	}

	err = q.Close(ctx)

	if err != nil {
		t.Fatalf("Failed to close queue, %v", err)
	}
}

package queue

import (
	"testing"

	"github.com/tharun-ragu22/job_scheduler/internal/jobs"
)



func Test_Enqueue_WhenEmailJobAddedToScheduler_JobAddedToScheduler(t *testing.T) {
    // Given a user wants to send an email
    scheduler := &Scheduler{}

    // When they submit a request to send an email
    emailJob := &jobs.EmailJob{
        To:      "user@example.com",
        Subject: "Test Email",
        Body:    "Hello World",
    }

    scheduler.Enqueue(emailJob)

	addedJob := scheduler.PeekTop()

    // Then the system accepts the task and schedules it for processing
    if addedJob == nil {
		t.Fatalf("Expected a job to be added to the scheduler, got nil")
	}

    queuedJob, ok := (*addedJob).(*jobs.EmailJob)
    if !ok {
        t.Fatalf("Expected job to be of type *EmailJob")
    }

    if queuedJob.To != "user@example.com" || queuedJob.Subject != "Test Email" {
        t.Errorf("Queued job does not match submitted job")
    }
}


func Test_Enqueue_WhenAPIRequestAddedToScheduler_JobAddedToScheduler(t *testing.T) {
    // Given a user wants to make an API request
    scheduler := &Scheduler{}

    example_url := "www.example.com"
    example_method := "POST"

    // When they submit a request to send the request
    apiJob := &jobs.APIJob{
        URL:      example_url,
        Method: example_method,
    }

    scheduler.Enqueue(apiJob)

	addedJob := scheduler.PeekTop()

    // Then the system accepts the request and schedules it for processing
    if addedJob == nil {
		t.Fatalf("Expected a job to be added to the scheduler, got nil")
	}

    queuedJob, ok := (*addedJob).(*jobs.APIJob)
    if !ok {
        t.Fatalf("Expected job to be of type *APIJob")
    }

    if queuedJob.URL != example_url || queuedJob.Method != example_method {
        t.Errorf("Queued job does not match submitted job")
    }
}

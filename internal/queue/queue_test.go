package queue

import (
	"reflect"
	"testing"

	"github.com/tharun-ragu22/job_scheduler/internal/jobs"
)

type JobTestCase struct {
    name string
    job jobs.Job
    expectedType reflect.Type

}

func Test_Enqueue_WhenJobAddedToScheduler_JobAddedToScheduler(t *testing.T) {
    // Given a user wants to create a task

    emailJob := &jobs.EmailJob{
        To:      "user@example.com",
        Subject: "Test Email",
        Body:    "Hello World",
    }

    apiJob := &jobs.APIJob{
        URL:      "www.example.com",
        Method: "POST",
    }

    testCases := []JobTestCase{
        {
            name: "Email Job",
            job:          emailJob,
            expectedType: reflect.TypeOf(&jobs.EmailJob{}),
        },
        {
            name: "API Job",
            job:          apiJob,
            expectedType: reflect.TypeOf(&jobs.APIJob{}),
        },
    }

    for _, tc := range testCases {
    
        scheduler := &Scheduler{}

        // When they submit the required information for the task
        scheduler.Enqueue(tc.job)

        // Then the system accepts the task and schedules it for processing
        addedJob := scheduler.PeekTop()

        if addedJob == nil {
            t.Fatalf("Expected a job to be added to the scheduler, got nil")
        }
        
        if *addedJob != tc.job {
            t.Errorf("Queued job does not match submitted job %s", tc.name)
        }
    }
}

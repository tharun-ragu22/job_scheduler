package queue

import (
	"github.com/tharun-ragu22/job_scheduler/internal/jobs"
)

// Job represents a task in the scheduler.

type Job = jobs.Job

type Scheduler struct {
	Jobs []Job
}

func (s *Scheduler) Enqueue(job Job) {
	s.Jobs = append(s.Jobs, job)
}

func (s *Scheduler) PeekTop() *Job {
	if len(s.Jobs) == 0 {
		return nil
	}
	topJob := s.Jobs[0]
	return &topJob
}



package jobs

import "fmt"

type CronJob struct {
	ID int
	Schedule string
	Task     string
}

func (c *CronJob) Run() error {
	// Simulate sending an email
	fmt.Printf("Running Cron Job %d", c.ID)
	return nil
}
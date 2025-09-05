package jobs

import "fmt"
type EmailJob struct {
	To      string
	Subject string
	Body    string
}

func (e *EmailJob) Run() error {
	// Simulate sending an email
	fmt.Printf("Sending email to %s with subject %s\n", e.To, e.Subject)
	return nil
}
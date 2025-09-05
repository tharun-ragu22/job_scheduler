package jobs

import "fmt"
type APIJob struct {
	URL string
	Method string
}

func (e *APIJob) Run() error {
	// Simulate sending an email
	fmt.Printf("Sending API %s request to  %s\n", e.Method, e.URL)
	return nil
}
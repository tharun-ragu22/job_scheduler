package jobs

import "testing"

func TestEmailJob(t *testing.T){
	payload := EmailPayload{
		To: "user@example.com",
        Subject: "Hello",
        Body: "Welcome to the job queue!",
    }

	result, err := RunEmailJob(payload)
	if err != nil {
		t.Fatalf("expected no error, got %s", err)
	}

	if result.Status != "sent" {
		t.Fatalf("expected status sent, got %s", result.Status)
	}
}
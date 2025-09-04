package jobs

// Mock EmailPayload struct
type EmailPayload struct {
    To      string
    Subject string
    Body    string
}

// Mock JobResult struct
type JobResult struct {
    Status    string
    MessageID string
}

// Mock RunEmailJob function
func RunEmailJob(payload EmailPayload) (*JobResult, error) {
    // Just return a dummy result immediately
    return &JobResult{
        Status:    "sent",
        MessageID: "mock123",
    }, nil
}

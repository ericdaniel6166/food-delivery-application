package asyncjob

//go:generate go run github.com/dmarkham/enumer -type=JobState -json -sql -transform=snake-upper
type JobState int

const (
	StateInit JobState = iota
	StateRunning
	StateFailed
	StateTimeout
	StateCompleted
	StateRetryFailed
)

package enums

type JobStatus string

const (
	JobStatus_Runing    JobStatus = "running"
	JobStatus_Succeeded JobStatus = "succeeded"
	JobStatus_Failed    JobStatus = "failed"
)

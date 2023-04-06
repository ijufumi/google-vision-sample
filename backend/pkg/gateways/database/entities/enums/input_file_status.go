package enums

type InputFileStatus string

const (
	InputFileStatus_Runing    InputFileStatus = "running"
	InputFileStatus_Succeeded InputFileStatus = "succeeded"
	InputFileStatus_Failed    InputFileStatus = "failed"
)

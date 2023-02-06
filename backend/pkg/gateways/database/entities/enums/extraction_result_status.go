package enums

type ExtractionResultStatus string

const (
	ExtractionResultStatus_Runing    ExtractionResultStatus = "running"
	ExtractionResultStatus_Succeeded ExtractionResultStatus = "succeeded"
	ExtractionResultStatus_Failed    ExtractionResultStatus = "failed"
)

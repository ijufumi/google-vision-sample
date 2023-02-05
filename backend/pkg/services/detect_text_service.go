package services

type DetectTextService interface {
}

func NewDetectTextService() DetectTextService {
	return &detectTextService{}
}

type detectTextService struct {
}

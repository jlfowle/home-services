package canary

import (
	"net/http"
)

type CanaryService interface {
	GetServiceHealth() (int, error)
	GetServiceReadiness() (int, error)
}

func NewCanaryService() CanaryService {
	return &canaryService{}
}

type canaryService struct{}

func (canaryService) GetServiceHealth() (int, error) {
	return http.StatusOK, nil
}

func (canaryService) GetServiceReadiness() (int, error) {
	return http.StatusOK, nil
}

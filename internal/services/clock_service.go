package services

import (
	"time"
)

//go:generate mockgen -destination=../../mocks/services/mock_clock_service.go -package=services_mocks github.com/fromsi/tg_reaction/internal/services ClockService
type ClockService interface {
	Now() time.Time
}

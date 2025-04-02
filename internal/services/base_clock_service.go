package services

import (
	"time"
)

type BaseClockService struct{}

func NewBaseClockService() *BaseClockService {
	return &BaseClockService{}
}

func (receiver *BaseClockService) Now() time.Time {
	return time.Now()
}

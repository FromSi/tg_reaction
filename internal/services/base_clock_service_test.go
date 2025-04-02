package services

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBaseClockService_Now(t *testing.T) {
	service := NewBaseClockService()

	before := time.Now()
	result := service.Now()
	after := time.Now()

	assert.True(t, !result.Before(before), "Now() result should not be before the time before the call")
	assert.True(t, !result.After(after), "Now() result should not be after the time after the call")
}

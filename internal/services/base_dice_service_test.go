package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBaseDiceService_IsWinningValue(t *testing.T) {
	service := NewBaseDiceService()

	assert.True(t, service.IsWinningValue(6), "Value 6 should be a winning value")

	for i := 1; i <= 5; i++ {
		assert.False(t, service.IsWinningValue(i), "Value %d should not be a winning value", i)
	}
}

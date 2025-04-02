package services

import (
	"crypto/rand"
	"io"
	"regexp"
	"testing"
	"time"

	services_mocks "github.com/fromsi/tg_reaction/mocks/services"
	"github.com/fromsi/tg_reaction/pkg/json"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

// mockErrorReader implements io.Reader and always returns an error
type mockErrorReader struct{}

func (m *mockErrorReader) Read(p []byte) (n int, err error) {
	return 0, io.ErrClosedPipe
}

func TestBaseRegexService_GetRandomReaction_CryptoError(t *testing.T) {
	originalReader := rand.Reader
	defer func() {
		rand.Reader = originalReader
	}()

	rand.Reader = &mockErrorReader{}

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockClock := services_mocks.NewMockClockService(mockController)

	config := &json.Config{}
	service := NewBaseRegexService(config, mockClock)

	reactions := []json.Reaction{json.ThumbsUp, json.Heart, json.Fire}

	results := make(map[json.Reaction]bool)
	for i := 0; i < 10; i++ {
		time.Sleep(time.Nanosecond)
		reaction := service.getRandomReaction(reactions)
		results[reaction] = true
	}

	assert.Greater(t, len(results), 1, "Should get different reactions with time-based fallback")

	for reaction := range results {
		found := false
		for _, validReaction := range reactions {
			if reaction == validReaction {
				found = true
				break
			}
		}
		assert.True(t, found, "Returned reaction should be from the input list")
	}
}

func TestBaseRegexService_FindReaction_Everyday(t *testing.T) {
	config := &json.Config{
		Everyday: []json.Pattern{
			{
				Pattern:   regexp.MustCompile("(?i)привет|здравствуй"),
				Reactions: []json.Reaction{json.ThumbsUp, json.Heart},
			},
			{
				Pattern:   regexp.MustCompile("(?i)пока|до свидания"),
				Reactions: []json.Reaction{json.Handshake},
			},
		},
	}

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockClock := services_mocks.NewMockClockService(mockController)
	mockClock.EXPECT().Now().Return(time.Date(2023, 6, 1, 12, 0, 0, 0, time.UTC)).AnyTimes()

	service := NewBaseRegexService(config, mockClock)

	reaction := service.FindReaction("Привет, как дела?")
	assert.Contains(t, []json.Reaction{json.ThumbsUp, json.Heart}, reaction)

	reaction = service.FindReaction("До свидания!")
	assert.Equal(t, json.Handshake, reaction)

	reaction = service.FindReaction("Хорошая погода сегодня.")
	assert.Equal(t, json.Reaction(""), reaction, "Should return empty for non-matching text")
}

func TestBaseRegexService_FindReaction_Holiday(t *testing.T) {
	config := &json.Config{
		Everyday: []json.Pattern{
			{
				Pattern:   regexp.MustCompile("(?i)привет|здравствуй"),
				Reactions: []json.Reaction{json.ThumbsUp, json.Heart},
			},
		},
		Holidays: map[string]json.Holiday{
			"new_year": {
				StartDay:   25,
				StartMonth: 12,
				EndDay:     5,
				EndMonth:   1,
				Reactions:  []json.Reaction{json.ChristmasTree, json.Santa},
				Patterns: []json.Pattern{
					{
						Pattern:   regexp.MustCompile("(?i)с новым годом|рождество"),
						Reactions: []json.Reaction{json.ChristmasTree, json.Santa, json.Snowman},
					},
				},
			},
		},
	}

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockClock := services_mocks.NewMockClockService(mockController)

	service := NewBaseRegexService(config, mockClock)

	// Testing during holiday period
	holidayDate := time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)
	mockClock.EXPECT().Now().Return(holidayDate).Times(3)

	holiday, holidayName := service.getCurrentHoliday()
	assert.NotNil(t, holiday, "Should be in a holiday period")
	assert.Equal(t, "new_year", holidayName, "Holiday should be new_year")

	reaction := service.FindReaction("С Новым Годом!")
	assert.Contains(t, []json.Reaction{json.ChristmasTree, json.Santa, json.Snowman}, reaction)

	reaction = service.FindReaction("Привет, как дела?")
	assert.Contains(t, []json.Reaction{json.ChristmasTree, json.Santa}, reaction,
		"Should use holiday reactions for everyday patterns during holiday period")

	// Testing outside holiday period (using the same service and mock)
	
	nonHolidayDate := time.Date(2023, 6, 1, 12, 0, 0, 0, time.UTC)
	mockClock.EXPECT().Now().Return(nonHolidayDate).Times(2)

	holiday, holidayName = service.getCurrentHoliday()
	assert.Nil(t, holiday, "Should not be in a holiday period")
	assert.Equal(t, "", holidayName, "No holiday should be active")

	reaction = service.FindReaction("С Новым Годом!")
	assert.Equal(t, json.Reaction(""), reaction)
}

func TestBaseRegexService_isDateInHolidayPeriod(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockClock := services_mocks.NewMockClockService(mockController)

	service := NewBaseRegexService(&json.Config{}, mockClock)

	holiday := json.Holiday{
		StartDay:   10,
		StartMonth: 5,
		EndDay:     20,
		EndMonth:   5,
	}

	assert.False(t, service.isDateInHolidayPeriod(9, 5, holiday))
	assert.True(t, service.isDateInHolidayPeriod(10, 5, holiday))
	assert.True(t, service.isDateInHolidayPeriod(15, 5, holiday))
	assert.True(t, service.isDateInHolidayPeriod(20, 5, holiday))
	assert.False(t, service.isDateInHolidayPeriod(21, 5, holiday))

	holiday = json.Holiday{
		StartDay:   25,
		StartMonth: 5,
		EndDay:     5,
		EndMonth:   6,
	}

	assert.False(t, service.isDateInHolidayPeriod(24, 5, holiday))
	assert.True(t, service.isDateInHolidayPeriod(25, 5, holiday))
	assert.True(t, service.isDateInHolidayPeriod(31, 5, holiday))
	assert.True(t, service.isDateInHolidayPeriod(1, 6, holiday))
	assert.True(t, service.isDateInHolidayPeriod(5, 6, holiday))
	assert.False(t, service.isDateInHolidayPeriod(6, 6, holiday))

	holiday = json.Holiday{
		StartDay:   25,
		StartMonth: 12,
		EndDay:     5,
		EndMonth:   1,
	}

	assert.False(t, service.isDateInHolidayPeriod(24, 12, holiday))
	assert.True(t, service.isDateInHolidayPeriod(25, 12, holiday))
	assert.True(t, service.isDateInHolidayPeriod(31, 12, holiday))
	assert.True(t, service.isDateInHolidayPeriod(1, 1, holiday))
	assert.True(t, service.isDateInHolidayPeriod(5, 1, holiday))
	assert.False(t, service.isDateInHolidayPeriod(6, 1, holiday))
}

// TestBaseRegexService_FindReaction_PatternWithoutReactions tests the case when a pattern has no reactions but the holiday has reactions
func TestBaseRegexService_FindReaction_PatternWithoutReactions(t *testing.T) {
	config := &json.Config{
		Holidays: map[string]json.Holiday{
			"new_year": {
				StartDay:   25,
				StartMonth: 12,
				EndDay:     5,
				EndMonth:   1,
				Reactions:  []json.Reaction{json.ChristmasTree, json.Santa},
				Patterns: []json.Pattern{
					{
						Pattern:   regexp.MustCompile("(?i)с новым годом|рождество"),
						Reactions: []json.Reaction{}, // Empty reactions list
					},
				},
			},
		},
	}

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockClock := services_mocks.NewMockClockService(mockController)
	mockClock.EXPECT().Now().Return(time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)).AnyTimes()

	service := NewBaseRegexService(config, mockClock)

	reaction := service.FindReaction("С Новым Годом!")
	t.Logf("Reaction for 'Happy New Year!' with empty pattern reactions: %s", reaction)
	assert.Contains(t, []json.Reaction{json.ChristmasTree, json.Santa}, reaction)
}

// TestBaseRegexService_MatchPattern_InvalidRegex tests the case when a regex pattern is invalid
func TestBaseRegexService_MatchPattern_InvalidRegex(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockClock := services_mocks.NewMockClockService(mockController)

	service := NewBaseRegexService(&json.Config{}, mockClock)

	// Create a pattern with nil regexp
	pattern := json.Pattern{
		Pattern:   nil, // Nil regexp
		Reactions: []json.Reaction{},
	}

	// This should return false due to nil regexp
	result := service.matchPattern("test text", pattern)
	assert.False(t, result, "Should return false for nil regexp")
}

// TestBaseRegexService_GetRandomReaction_EmptyList tests the case when the reactions list is empty
func TestBaseRegexService_GetRandomReaction_EmptyList(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockClock := services_mocks.NewMockClockService(mockController)

	service := NewBaseRegexService(&json.Config{}, mockClock)

	// Test with empty reactions list
	reaction := service.getRandomReaction([]json.Reaction{})
	assert.Equal(t, json.Reaction(""), reaction, "Should return empty reaction for empty list")
}

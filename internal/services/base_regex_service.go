package services

import (
	"crypto/rand"
	"math/big"
	"time"

	"github.com/fromsi/tg_reaction/pkg/json"
)

type BaseRegexService struct {
	config       *json.Config
	clockService ClockService
}

func NewBaseRegexService(config *json.Config, clockService ClockService) *BaseRegexService {
	return &BaseRegexService{
		config:       config,
		clockService: clockService,
	}
}

// FindReaction finds an appropriate reaction for the text based on configuration
func (receiver *BaseRegexService) FindReaction(text string) json.Reaction {
	// First, check if we are in a holiday period
	holiday, _ := receiver.getCurrentHoliday()

	if holiday != nil {
		// Check holiday patterns
		for _, pattern := range holiday.Patterns {
			if receiver.matchPattern(text, pattern) {
				// If there's a match, return a random reaction from the list
				if len(pattern.Reactions) > 0 {
					return receiver.getRandomReaction(pattern.Reactions)
				}
				// If the pattern has no reactions, use holiday reactions
				if len(holiday.Reactions) > 0 {
					return receiver.getRandomReaction(holiday.Reactions)
				}
			}
		}

		// If no matches found in holiday patterns, check everyday patterns
		for _, pattern := range receiver.config.Everyday {
			if receiver.matchPattern(text, pattern) {
				// During holiday period, use holiday reactions instead of everyday reactions
				if len(holiday.Reactions) > 0 {
					return receiver.getRandomReaction(holiday.Reactions)
				}
			}
		}
	} else {
		// If not in a holiday period, check everyday patterns
		for _, pattern := range receiver.config.Everyday {
			if receiver.matchPattern(text, pattern) {
				if len(pattern.Reactions) > 0 {
					return receiver.getRandomReaction(pattern.Reactions)
				}
			}
		}
	}

	// If nothing found, return an empty reaction
	return ""
}

// matchPattern checks if the text matches the pattern
func (receiver *BaseRegexService) matchPattern(text string, pattern json.Pattern) bool {
	// Simply use the pre-compiled regexp
	if pattern.Pattern == nil {
		return false
	}

	return pattern.Pattern.MatchString(text)
}

// getCurrentHoliday returns the current holiday if we are in a holiday period
func (receiver *BaseRegexService) getCurrentHoliday() (*json.Holiday, string) {
	now := receiver.clockService.Now()
	currentDay := now.Day()
	currentMonth := int(now.Month())

	for holidayName, holiday := range receiver.config.Holidays {
		// Check if we are in a holiday period
		if receiver.isDateInHolidayPeriod(currentDay, currentMonth, holiday) {
			return &holiday, holidayName
		}
	}

	return nil, ""
}

// isDateInHolidayPeriod checks if the date is within a holiday period
func (receiver *BaseRegexService) isDateInHolidayPeriod(day, month int, holiday json.Holiday) bool {
	startDay, startMonth := holiday.StartDay, holiday.StartMonth
	endDay, endMonth := holiday.EndDay, holiday.EndMonth

	// If the holiday is within a single month
	if startMonth == endMonth {
		return month == startMonth && day >= startDay && day <= endDay
	}

	// If the holiday spans to the next month
	if startMonth < endMonth {
		return (month > startMonth && month < endMonth) ||
			(month == startMonth && day >= startDay) ||
			(month == endMonth && day <= endDay)
	}

	// If the holiday spans to the next year (e.g., from December to January)
	return (month > startMonth || month < endMonth) ||
		(month == startMonth && day >= startDay) ||
		(month == endMonth && day <= endDay)
}

// getRandomReaction returns a random reaction from the list
func (receiver *BaseRegexService) getRandomReaction(reactions []json.Reaction) json.Reaction {
	if len(reactions) == 0 {
		return ""
	}

	// Generate cryptographically secure random index
	max := big.NewInt(int64(len(reactions)))
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		// Fallback to time-based random in case of error
		currentTime := time.Now()
		return reactions[int(currentTime.UnixNano())%len(reactions)]
	}

	return reactions[n.Int64()]
}

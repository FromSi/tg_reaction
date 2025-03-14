package json

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
)

// JSONPattern represents the JSON structure of a pattern
type JSONPattern struct {
	Pattern   string   `json:"pattern"`
	Reactions []string `json:"reactions"`
	Prefix    string   `json:"prefix,omitempty"`
	Suffix    string   `json:"suffix,omitempty"`
}

// JSONHoliday represents the JSON structure of a holiday
type JSONHoliday struct {
	StartDay   int           `json:"start_day"`
	StartMonth int           `json:"start_month"`
	EndDay     int           `json:"end_day"`
	EndMonth   int           `json:"end_month"`
	Reactions  []string      `json:"reactions"`
	Patterns   []JSONPattern `json:"patterns,omitempty"`
}

// JSONConfig represents the JSON structure of the configuration
type JSONConfig struct {
	Common struct {
		Prefix string `json:"prefix"`
		Suffix string `json:"suffix"`
	} `json:"common"`
	Everyday []JSONPattern          `json:"everyday"`
	Holidays map[string]JSONHoliday `json:"holidays"`
}

// Load loads configuration from a JSON file
func Load(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var jsonConfig JSONConfig
	if err := json.Unmarshal(data, &jsonConfig); err != nil {
		return nil, err
	}

	return convertConfig(&jsonConfig)
}

// convertPattern converts JSONPattern to Pattern considering prefix and suffix inheritance
func convertPattern(pattern JSONPattern, commonPrefix, commonSuffix string, patternName, holidayName string) (Pattern, error) {
	reactions := make([]Reaction, len(pattern.Reactions))
	for i, r := range pattern.Reactions {
		reaction := Reaction(r)
		if !reaction.IsValid() {
			if holidayName != "" {
				if patternName != "" {
					return Pattern{}, fmt.Errorf("invalid reaction %q in holiday %q pattern %q", r, holidayName, patternName)
				}
				return Pattern{}, fmt.Errorf("invalid reaction %q in holiday %q", r, holidayName)
			}
			return Pattern{}, fmt.Errorf("invalid reaction %q in everyday pattern %q", r, patternName)
		}
		reactions[i] = reaction
	}

	// Apply inheritance rules for prefix and suffix
	prefix := pattern.Prefix
	if prefix == "" {
		prefix = commonPrefix
	}
	suffix := pattern.Suffix
	if suffix == "" {
		suffix = commonSuffix
	}

	// Compile the regular expression
	fullPattern := prefix + pattern.Pattern + suffix
	re, err := regexp.Compile(fullPattern)
	if err != nil {
		if holidayName != "" {
			if patternName != "" {
				return Pattern{}, fmt.Errorf("invalid regex pattern %q in holiday %q pattern %q: %v", fullPattern, holidayName, patternName, err)
			}
			return Pattern{}, fmt.Errorf("invalid regex pattern %q in holiday %q: %v", fullPattern, holidayName, err)
		}
		return Pattern{}, fmt.Errorf("invalid regex pattern %q in everyday pattern %q: %v", fullPattern, patternName, err)
	}

	return Pattern{
		Pattern:   re,
		Reactions: reactions,
	}, nil
}

// convertConfig converts JSON configuration to Config
func convertConfig(jsonConfig *JSONConfig) (*Config, error) {
	config := &Config{
		Everyday: make([]Pattern, 0, len(jsonConfig.Everyday)),
		Holidays: make(map[string]Holiday),
	}

	// Convert everyday patterns
	for _, pattern := range jsonConfig.Everyday {
		convertedPattern, err := convertPattern(pattern, jsonConfig.Common.Prefix, jsonConfig.Common.Suffix, "", "")
		if err != nil {
			return nil, err
		}
		config.Everyday = append(config.Everyday, convertedPattern)
	}

	// Convert holidays
	for name, holiday := range jsonConfig.Holidays {
		// Check holiday's main reactions
		holidayPattern := JSONPattern{
			Reactions: holiday.Reactions,
		}
		convertedHolidayPattern, err := convertPattern(holidayPattern, jsonConfig.Common.Prefix, jsonConfig.Common.Suffix, "", name)
		if err != nil {
			return nil, err
		}
		reactions := convertedHolidayPattern.Reactions

		patterns := make([]Pattern, 0, len(holiday.Patterns))
		for _, pattern := range holiday.Patterns {
			convertedPattern, err := convertPattern(pattern, jsonConfig.Common.Prefix, jsonConfig.Common.Suffix, "", name)
			if err != nil {
				return nil, err
			}
			patterns = append(patterns, convertedPattern)
		}

		config.Holidays[name] = Holiday{
			StartDay:   holiday.StartDay,
			StartMonth: holiday.StartMonth,
			EndDay:     holiday.EndDay,
			EndMonth:   holiday.EndMonth,
			Reactions:  reactions,
			Patterns:   patterns,
		}
	}

	return config, nil
}

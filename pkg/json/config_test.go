package json

import (
	"os"
	"strings"
	"testing"
)

func TestLoad(t *testing.T) {
	// Create a temporary JSON file for testing
	jsonContent := `{
		"common": {
			"prefix": "(?i)(\\A|\\s)(",
			"suffix": ")(\\s|\\z)"
		},
		"everyday": [
			{
				"pattern": "всем|утро|привет|здравствуйте",
				"reactions": ["🤝", "👌"],
				"prefix": "(?i)(",
				"suffix": ")"
			},
			{
				"pattern": "обед|приятного|аппетита",
				"reactions": ["🌭", "🍌"]
			},
			{
				"pattern": "встреча|митинг|созвон",
				"reactions": ["🙉", "👌", "✍️"]
			}
		],
		"holidays": {
			"new_year": {
				"start_day": 31,
				"start_month": 12,
				"end_day": 2,
				"end_month": 1,
				"reactions": ["🎅", "🎄", "☃️"],
				"patterns": [
					{
						"pattern": "поздравляю|с новым годом",
						"reactions": ["🎄", "🎅"]
					}
				]
			},
			"halloween": {
				"start_day": 31,
				"start_month": 10,
				"end_day": 31,
				"end_month": 10,
				"reactions": ["🎃", "👻"],
				"patterns": [
					{
						"pattern": "страшно|ужас|тыква",
						"reactions": ["👻"],
						"prefix": "(?i)(",
						"suffix": ")"
					},
					{
						"pattern": "конфеты|сладости",
						"reactions": ["🍌", "🎉"]
					}
				]
			},
			"valentine": {
				"start_day": 14,
				"start_month": 2,
				"end_day": 14,
				"end_month": 2,
				"reactions": ["❤️", "💘", "💋"]
			},
			"womens_day": {
				"start_day": 8,
				"start_month": 3,
				"end_day": 8,
				"end_month": 3,
				"reactions": ["💋", "🤗", "🎉"],
				"patterns": [
					{
						"pattern": "с праздником|с 8 марта",
						"reactions": ["💋", "🎉"]
					}
				]
			}
		}
	}`

	tmpfile, err := os.CreateTemp("", "reactions*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(jsonContent)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Test configuration loading
	config, err := Load(tmpfile.Name())
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	// Check everyday patterns
	if len(config.Everyday) != 3 {
		t.Errorf("Expected 3 everyday patterns, got %d", len(config.Everyday))
	}

	// Check first pattern with overridden prefix/suffix
	pattern := config.Everyday[0]
	if pattern.Pattern == nil {
		t.Fatal("Expected compiled regexp in first pattern")
	}
	// Test that the pattern matches expected strings
	if !pattern.Pattern.MatchString("привет") {
		t.Errorf("First pattern should match 'привет'")
	}
	if !pattern.Pattern.MatchString("здравствуйте") {
		t.Errorf("First pattern should match 'здравствуйте'")
	}
	if pattern.Pattern.MatchString("до свидания") {
		t.Errorf("First pattern should not match 'до свидания'")
	}

	// Check second pattern without prefix/suffix override
	pattern = config.Everyday[1]
	if pattern.Pattern == nil {
		t.Fatal("Expected compiled regexp in second pattern")
	}
	// Test that the pattern matches expected strings
	if !pattern.Pattern.MatchString("обед") {
		t.Errorf("Second pattern should match 'обед'")
	}
	if !pattern.Pattern.MatchString("приятного аппетита") {
		t.Errorf("Second pattern should match 'приятного аппетита'")
	}
	if pattern.Pattern.MatchString("привет") {
		t.Errorf("Second pattern should not match 'привет'")
	}

	// Check holidays
	if len(config.Holidays) != 4 {
		t.Errorf("Expected 4 holidays, got %d", len(config.Holidays))
	}

	// Check New Year
	newYear, ok := config.Holidays["new_year"]
	if !ok {
		t.Fatal("Expected new_year holiday")
	}
	if len(newYear.Patterns) != 1 {
		t.Errorf("Expected 1 pattern in new_year, got %d", len(newYear.Patterns))
	}
	pattern = newYear.Patterns[0]
	if pattern.Pattern == nil {
		t.Fatal("Expected compiled regexp in new_year pattern")
	}
	if !pattern.Pattern.MatchString("поздравляю") {
		t.Errorf("New year pattern should match 'поздравляю'")
	}

	// Check Halloween (with prefix/suffix override in pattern)
	halloween, ok := config.Holidays["halloween"]
	if !ok {
		t.Fatal("Expected halloween holiday")
	}
	if len(halloween.Patterns) != 2 {
		t.Errorf("Expected 2 patterns in halloween, got %d", len(halloween.Patterns))
	}
	pattern = halloween.Patterns[0]
	if pattern.Pattern == nil {
		t.Fatal("Expected compiled regexp in halloween pattern")
	}
	// Test that the pattern matches expected strings
	if !pattern.Pattern.MatchString("страшно") {
		t.Errorf("Halloween pattern should match 'страшно'")
	}
	if !pattern.Pattern.MatchString("ужас") {
		t.Errorf("Halloween pattern should match 'ужас'")
	}

	// Check Valentine's Day
	valentine, ok := config.Holidays["valentine"]
	if !ok {
		t.Fatal("Expected valentine holiday")
	}
	if len(valentine.Patterns) != 0 {
		t.Errorf("Expected no patterns in valentine, got %d", len(valentine.Patterns))
	}

	// Check Women's Day
	womensDay, ok := config.Holidays["womens_day"]
	if !ok {
		t.Fatal("Expected womens_day holiday")
	}
	if len(womensDay.Patterns) != 1 {
		t.Errorf("Expected 1 pattern in womens_day, got %d", len(womensDay.Patterns))
	}
}

func TestLoadErrors(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		wantErr  bool
		errMatch string
	}{
		{
			name:    "non_existent_file",
			wantErr: true,
		},
		{
			name:     "invalid_json",
			content:  "invalid json",
			wantErr:  true,
			errMatch: "invalid",
		},
		{
			name: "invalid_reaction_in_everyday",
			content: `{
				"everyday": [
					{
						"pattern": "привет",
						"reactions": ["invalid_reaction"]
					}
				]
			}`,
			wantErr:  true,
			errMatch: "invalid reaction \"invalid_reaction\" in everyday pattern \"\"",
		},
		{
			name: "invalid_reaction_in_holiday",
			content: `{
				"holidays": {
					"new_year": {
						"start_day": 31,
						"start_month": 12,
						"end_day": 1,
						"end_month": 1,
						"reactions": ["invalid_reaction"]
					}
				}
			}`,
			wantErr:  true,
			errMatch: "invalid reaction \"invalid_reaction\" in holiday \"new_year\"",
		},
		{
			name: "invalid_reaction_in_holiday_pattern",
			content: `{
				"holidays": {
					"new_year": {
						"start_day": 31,
						"start_month": 12,
						"end_day": 1,
						"end_month": 1,
						"reactions": ["🎅"],
						"patterns": [
							{
								"pattern": "поздравляю",
								"reactions": ["invalid_reaction"]
							}
						]
					}
				}
			}`,
			wantErr:  true,
			errMatch: "invalid reaction \"invalid_reaction\" in holiday \"new_year\"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var tmpfile *os.File
			var err error

			if tt.name == "non_existent_file" {
				_, err = Load("nonexistent.json")
			} else {
				tmpfile, err = os.CreateTemp("", "reactions*.json")
				if err != nil {
					t.Fatal(err)
				}
				defer os.Remove(tmpfile.Name())

				if _, err := tmpfile.Write([]byte(tt.content)); err != nil {
					t.Fatal(err)
				}
				if err := tmpfile.Close(); err != nil {
					t.Fatal(err)
				}

				_, err = Load(tmpfile.Name())
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.errMatch != "" && (err == nil || !strings.Contains(err.Error(), tt.errMatch)) {
				t.Errorf("Load() error = %v, should contain %q", err, tt.errMatch)
			}
		})
	}
}

func TestConvertPattern(t *testing.T) {
	tests := []struct {
		name         string
		pattern      JSONPattern
		commonPrefix string
		commonSuffix string
		patternName  string
		holidayName  string
		wantErr      bool
		errMatch     string
	}{
		{
			name: "everyday pattern with custom prefix and suffix",
			pattern: JSONPattern{
				Pattern:   "привет|здравствуй",
				Reactions: []string{"👍", "❤️"},
				Prefix:    "(?i)",
				Suffix:    "",
			},
			commonPrefix: "(?i)",
			commonSuffix: "",
			patternName:  "greeting",
			wantErr:      false,
		},
		{
			name: "pattern inheriting common prefix and suffix",
			pattern: JSONPattern{
				Pattern:   "обед|приятного",
				Reactions: []string{"🌭", "🍌"},
			},
			commonPrefix: "(?i)",
			commonSuffix: "",
			patternName:  "lunch",
			wantErr:      false,
		},
		{
			name: "holiday pattern",
			pattern: JSONPattern{
				Pattern:   "с новым годом|рождество",
				Reactions: []string{"🎄", "🎅", "☃️"},
				Prefix:    "(?i)",
				Suffix:    "",
			},
			commonPrefix: "(?i)",
			commonSuffix: "",
			patternName:  "holiday_greeting",
			holidayName:  "new_year",
			wantErr:      false,
		},
		{
			name: "invalid reaction in everyday pattern",
			pattern: JSONPattern{
				Pattern:   "привет",
				Reactions: []string{"invalid_reaction"},
			},
			commonPrefix: "(?i)",
			commonSuffix: "",
			patternName:  "greeting",
			wantErr:      true,
			errMatch:     "invalid reaction \"invalid_reaction\" in everyday pattern \"greeting\"",
		},
		{
			name: "invalid reaction in holiday pattern",
			pattern: JSONPattern{
				Pattern:   "с новым годом",
				Reactions: []string{"invalid_reaction"},
			},
			commonPrefix: "(?i)",
			commonSuffix: "",
			patternName:  "holiday_greeting",
			holidayName:  "new_year",
			wantErr:      true,
			errMatch:     "invalid reaction \"invalid_reaction\" in holiday \"new_year\" pattern \"holiday_greeting\"",
		},
		{
			name: "invalid regex pattern",
			pattern: JSONPattern{
				Pattern:   "[", // Invalid regex pattern (unclosed character class)
				Reactions: []string{"👍"},
			},
			commonPrefix: "(?i)",
			commonSuffix: "",
			patternName:  "invalid_regex",
			wantErr:      true,
			errMatch:     "invalid regex pattern",
		},
		{
			name: "invalid regex pattern in holiday",
			pattern: JSONPattern{
				Pattern:   "[", // Invalid regex pattern (unclosed character class)
				Reactions: []string{"🎄"},
			},
			commonPrefix: "(?i)",
			commonSuffix: "",
			patternName:  "",
			holidayName:  "new_year",
			wantErr:      true,
			errMatch:     "invalid regex pattern \"(?i)[\" in holiday \"new_year\"",
		},
		{
			name: "invalid regex pattern in holiday pattern",
			pattern: JSONPattern{
				Pattern:   "[", // Invalid regex pattern (unclosed character class)
				Reactions: []string{"🎄"},
			},
			commonPrefix: "(?i)",
			commonSuffix: "",
			patternName:  "congrats",
			holidayName:  "new_year",
			wantErr:      true,
			errMatch:     "invalid regex pattern \"(?i)[\" in holiday \"new_year\" pattern \"congrats\"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pattern, err := convertPattern(tt.pattern, tt.commonPrefix, tt.commonSuffix, tt.patternName, tt.holidayName)

			if tt.wantErr {
				if err == nil {
					t.Errorf("convertPattern() error = nil, wantErr %v", tt.wantErr)
					return
				}
				if tt.errMatch != "" && !strings.Contains(err.Error(), tt.errMatch) {
					t.Errorf("convertPattern() error = %v, want error containing %q", err, tt.errMatch)
				}
				return
			}

			if err != nil {
				t.Errorf("convertPattern() unexpected error = %v", err)
				return
			}

			// Check that the pattern was compiled correctly
			if pattern.Pattern == nil {
				t.Errorf("convertPattern() pattern.Pattern is nil")
				return
			}

			// Test that the pattern matches expected strings
			if tt.name == "everyday pattern with custom prefix and suffix" {
				if !pattern.Pattern.MatchString("Привет") {
					t.Errorf("Pattern %q should match 'Привет'", pattern.Pattern.String())
				}
				if !pattern.Pattern.MatchString("здравствуй") {
					t.Errorf("Pattern %q should match 'здравствуй'", pattern.Pattern.String())
				}
				if pattern.Pattern.MatchString("до свидания") {
					t.Errorf("Pattern %q should not match 'до свидания'", pattern.Pattern.String())
				}
			}

			// Check that reactions were converted correctly
			expectedReactions := make([]Reaction, len(tt.pattern.Reactions))
			for i, r := range tt.pattern.Reactions {
				expectedReactions[i] = Reaction(r)
			}

			if len(pattern.Reactions) != len(expectedReactions) {
				t.Errorf("convertPattern() got %d reactions, want %d", len(pattern.Reactions), len(expectedReactions))
				return
			}

			for i, r := range pattern.Reactions {
				if r != expectedReactions[i] {
					t.Errorf("convertPattern() reaction[%d] = %v, want %v", i, r, expectedReactions[i])
				}
			}
		})
	}
}

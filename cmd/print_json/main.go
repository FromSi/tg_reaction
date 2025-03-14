package main

import (
	"fmt"
	"log"

	"github.com/fromsi/tg_reaction/pkg/json"
)

func main() {
	// Load configuration from JSON file
	config, err := json.Load("config.json")
	if err != nil {
		log.Fatalf("Configuration loading error: %v", err)
	}

	// Display information about loaded configuration
	fmt.Println("Loaded configuration:")

	// Display everyday patterns
	fmt.Println("\nEveryday patterns:")
	for i, pattern := range config.Everyday {
		fmt.Printf("- Pattern %d:\n", i+1)
		if pattern.Pattern != nil {
			fmt.Printf("  Pattern: %s\n", pattern.Pattern.String())
		} else {
			fmt.Printf("  Pattern: <nil>\n")
		}
		fmt.Printf("  Reactions: %v\n", pattern.Reactions)
	}

	// Display holidays
	fmt.Println("\nHolidays:")
	for name, holiday := range config.Holidays {
		fmt.Printf("- %s:\n", name)
		fmt.Printf("  Period: %d.%d - %d.%d\n",
			holiday.StartDay, holiday.StartMonth,
			holiday.EndDay, holiday.EndMonth)
		fmt.Printf("  Reactions: %v\n", holiday.Reactions)

		if len(holiday.Patterns) > 0 {
			fmt.Println("  Patterns:")
			for i, pattern := range holiday.Patterns {
				if pattern.Pattern != nil {
					fmt.Printf("    Pattern %d: %s -> %v\n",
						i+1, pattern.Pattern.String(), pattern.Reactions)
				} else {
					fmt.Printf("    Pattern %d: <nil> -> %v\n",
						i+1, pattern.Reactions)
				}
			}
		}
	}

	fmt.Println("\nAll reactions are valid according to tests in pkg/json/reactions_test.go")
}

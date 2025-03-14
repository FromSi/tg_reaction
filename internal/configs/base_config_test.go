package configs

import (
	"os"
	"testing"
)

func TestBaseConfig_GetTelegramToken(t *testing.T) {
	key := "APP_TELEGRAM_TOKEN"

	os.Unsetenv(key)
	defer os.Unsetenv(key)

	// Test with default token
	config := NewBaseConfig()
	if got := config.GetTelegramToken(); got != BaseConfigDefaultAppTelegramToken {
		t.Errorf("GetTelegramToken() = %v, want %v", got, BaseConfigDefaultAppTelegramToken)
	}

	// Test with environment variable set
	expectedToken := "new_secret"
	os.Setenv(key, expectedToken)
	config = NewBaseConfig()
	if got := config.GetTelegramToken(); got != expectedToken {
		t.Errorf("GetTelegramToken() = %v, want %v", got, expectedToken)
	}
}

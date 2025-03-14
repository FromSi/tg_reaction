package env

import (
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	key := "TEST_ENV_VAR"
	defaultValue := "default"

	os.Unsetenv(key)
	defer os.Unsetenv(key)

	// Test when the environment variable is not set
	if got := GetEnv(key, defaultValue); got != defaultValue {
		t.Errorf("GetEnv() = %v, want %v", got, defaultValue)
	}

	// Test when the environment variable is set
	expectedValue := "value"
	os.Setenv(key, expectedValue)
	if got := GetEnv(key, defaultValue); got != expectedValue {
		t.Errorf("GetEnv() = %v, want %v", got, expectedValue)
	}
}

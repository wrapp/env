package env

// Convenience functions for working with environment variables

import (
	"os"
	"strings"
)

// Default gets the value of an environment variable or a default if not set
func Default(key string, def string) string {
	env := os.Getenv(key)
	if env == "" {
		return def
	}
	return env
}

// Bool returns the boolean value of a environment variable
// "true", "1",  "yes", "on" in any casing are considered true
func Bool(key string) bool {
	env := strings.ToLower(os.Getenv(key))
	if env == "1" || env == "true" || env == "yes" || env == "on" {
		return true
	}
	return false
}

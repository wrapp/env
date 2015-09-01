package env

// Convenience functions for working with environment variables

import (
	"os"
	"strconv"
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

// DefaultInt gets the parsed int value of an environment variable or a default if not set or parsing failed
func DefaultInt(key string, def int) int {
	env := os.Getenv(key)
	if env == "" {
		return def
	}
	i, err := strconv.ParseInt(env, 10, 32)
	if err != nil {
		return def
	}
	return int(i)
}

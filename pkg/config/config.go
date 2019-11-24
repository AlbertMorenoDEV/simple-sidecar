package config

import (
	"os"
	"strconv"
	"strings"
)

const defaultParametersPort = "7983"
const defaultParametersWriteTimeout = 15
const defaultParametersReadTimeout = 15
const defaultParametersIdleTimeout = 60
const defaultParametersGracefulTimeout = 15

// ParametersConfig contains the API server configuration
type ParametersConfig struct {
	Port            string
	WriteTimeout    int
	ReadTimeout     int
	IdleTimeout     int
	GracefulTimeout int
}

// Config contains full Simple Sidecar configuration
type Config struct {
	Parameters           ParametersConfig
	DebugMode            bool
	AuthenticationTokens []string
}

// New returns a new Config struct
func New() *Config {
	return &Config{
		Parameters: ParametersConfig{
			Port:            getEnv("SS_PARAMETERS_PORT", defaultParametersPort),
			WriteTimeout:    getEnvAsInt("SS_PARAMETERS_WRITE_TIMEOUT", defaultParametersWriteTimeout),
			ReadTimeout:     getEnvAsInt("SS_PARAMETERS_READ_TIMEOUT", defaultParametersReadTimeout),
			IdleTimeout:     getEnvAsInt("SS_PARAMETERS_IDLE_TIMEOUT", defaultParametersIdleTimeout),
			GracefulTimeout: getEnvAsInt("SS_PARAMETERS_GRACEFUL_TIMEOUT", defaultParametersGracefulTimeout),
		},
		DebugMode:            getEnvAsBool("SS_DEBUG_MODE", false),
		AuthenticationTokens: getEnvAsSlice("SS_AUTH_TOKENS", []string{""}, ","),
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// Simple helper function to read an environment variable into integer or return a default value
func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

// Helper to read an environment variable into a bool or return default value
func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

// Helper to read an environment variable into a string slice or return default value
func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := getEnv(name, "")

	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)

	return val
}

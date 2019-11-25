package config

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

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

// Setup prepares config
func Setup(com *cobra.Command) {
	viper.SetEnvPrefix("ss_parameters")

	// persistent flags
	com.PersistentFlags().StringP("port", "p", "7983", "api server port")
	com.PersistentFlags().Int("write-timeout", 15, "api server write timeout in seconds")
	com.PersistentFlags().Int("read-timeout", 15, "api server read timeout in seconds")
	com.PersistentFlags().Int("idle-timeout", 60, "api server idle timeout in seconds")
	com.PersistentFlags().Int("graceful-timeout", 15, "duration for which the server gracefully wait for existing connections to finish in seconds")
	com.PersistentFlags().Bool("debug-mode", false, "debug mode status")
	com.PersistentFlags().StringSlice("auth-tokens", []string{""}, "authetication tokens")

	// binded flags
	viper.BindPFlag("port", com.PersistentFlags().Lookup("port"))
	viper.BindPFlag("write-timeout", com.PersistentFlags().Lookup("write-timeout"))
	viper.BindPFlag("read-timeout", com.PersistentFlags().Lookup("read-timeout"))
	viper.BindPFlag("idle-timeout", com.PersistentFlags().Lookup("idle-timeout"))
	viper.BindPFlag("graceful-timeout", com.PersistentFlags().Lookup("graceful-timeout"))
	viper.BindPFlag("debug-mode", com.PersistentFlags().Lookup("debug-mode"))
	viper.BindPFlag("auth-tokens", com.PersistentFlags().Lookup("auth-tokens"))

	viper.BindEnv("port")
	viper.BindEnv("write-timeout")
	viper.BindEnv("read-timeout")
	viper.BindEnv("idle-timeout")
	viper.BindEnv("graceful-timeout")
	viper.BindEnv("debug-mode")
	viper.BindEnv("auth-tokens")
}

// New returns a new Config struct
func New() *Config {
	return &Config{
		Parameters: ParametersConfig{
			Port:            viper.GetString("port"),
			WriteTimeout:    viper.GetInt("write-timeout"),
			ReadTimeout:     viper.GetInt("read-timeout"),
			IdleTimeout:     viper.GetInt("idle-timeout"),
			GracefulTimeout: viper.GetInt("graceful-timeout"),
		},
		DebugMode:            viper.GetBool("debug-mode"),
		AuthenticationTokens: viper.GetStringSlice("auth-tokens"),
	}
}

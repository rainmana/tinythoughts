package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	FrameworksDir string
	SessionFile   string
	OutputFormat  string
}

func Load() (*Config, error) {
	viper.SetEnvPrefix("TINYTHOUGHTS")
	viper.AutomaticEnv()
	
	viper.SetDefault("frameworks_dir", "$HOME/.config/tinythoughts/frameworks")
	viper.SetDefault("session_file", "$HOME/.config/tinythoughts/session.json")
	viper.SetDefault("output_format", "json")
	
	viper.SetConfigName("tinythoughts")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.config/tinythoughts")
	viper.AddConfigPath(".")
	
	_ = viper.ReadInConfig()
	
	return &Config{
		FrameworksDir: viper.GetString("frameworks_dir"),
		SessionFile:   viper.GetString("session_file"),
		OutputFormat:  viper.GetString("output_format"),
	}, nil
}

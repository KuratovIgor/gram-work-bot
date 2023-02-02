package config

import (
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	TelegramToken string
	Messages      Messages
}

type Messages struct {
	Errors
	Responses
}

type Errors struct {
	Default string `mapstructure:"default"`
}

type Responses struct {
	Start          string `mapstructure:"start"`
	UnknownCommand string `mapstructure:"unknown_command"`
}

func Init() (*Config, error) {
	if err := setUoViper(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := parseEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func setUoViper() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")

	return viper.ReadInConfig()
}

func unmarshal(cfg *Config) error {
	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("messages.responses", &cfg.Messages.Responses); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("messages.errors", &cfg.Messages.Errors); err != nil {
		return err
	}

	return nil
}

func parseEnv(cfg *Config) error {
	os.Setenv("TELEGRAM_TOKEN", "5625272170:AAGQVFOEIh_aoRMUfB3vBXx6QrDBM5sLYro")

	if err := viper.BindEnv("telegram_token"); err != nil {
		return err
	}

	cfg.TelegramToken = viper.GetString("telegram_token")

	return nil
}

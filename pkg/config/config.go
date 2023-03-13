package config

import (
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	TelegramToken string
	RedirectURI   string
	ClientID      string
	ClientSecret  string
	Messages      Messages
	LkUrl         string
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
	os.Setenv("telegram_token", "5625272170:AAGQVFOEIh_aoRMUfB3vBXx6QrDBM5sLYro")
	os.Setenv("redirect_uri", "http://localhost/")
	os.Setenv("client_id", "P3PD344RTN89887V9SQJB24QNA0U06SPP3JG6TSR59SKMQ1191C8VHRJLC17RO0D")
	os.Setenv("client_secret", "SMIR5S7AK9P7FMOE4MRLJ291N4939TTM5BO9RGM7NOFSEC9RSMF6CCD50TVFBQUS")
	if err := viper.BindEnv("telegram_token"); err != nil {
		return err
	}

	if err := viper.BindEnv("redirect_uri"); err != nil {
		return err
	}

	if err := viper.BindEnv("client_id"); err != nil {
		return err
	}

	if err := viper.BindEnv("client_secret"); err != nil {
		return err
	}

	if err := viper.BindEnv("lk_url"); err != nil {
		return err
	}

	cfg.TelegramToken = viper.GetString("telegram_token")
	cfg.RedirectURI = viper.GetString("redirect_uri")
	cfg.ClientID = viper.GetString("client_id")
	cfg.ClientSecret = viper.GetString("client_secret")
	cfg.LkUrl = viper.GetString("lk_url")

	return nil
}

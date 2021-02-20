package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Telegram struct {
		BotToken string `mapstructure:"bot_token"`
	} `mapstructure:"telegram"`
}

var cfg Config

func Get() Config {
	return cfg
}

func Setup() {
	viper.SetConfigFile("config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Panic("Config dosyası okunamadı.")
	}

	err := viper.Unmarshal(&cfg)
	if err != nil {
		log.Panic("Config unmarshall edilemedi.")
	}
}

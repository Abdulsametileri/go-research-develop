package config

import (
	"github.com/spf13/viper"
	"log"
)

type TelegramConfig struct {
	BotToken             string `mapstructure:"bot_token"`
	AbdulsametTelegramId int64  `mapstructure:"abdulsamet_telegram_id"`
}

type MongoConfig struct {
	User     string
	Password string
	Name     string
	Host     string
	Port     int
}

type RedisConfig struct {
	Host     string
	Username string
	Password string
	DB       int
}

type Config struct {
	TelegramConfig `mapstructure:"telegram"`
	Database       struct {
		MongoConfig `mapstructure:"mongo"`
	}
	RedisConfig `mapstructure:"redis"`
}

var cfg Config

func Get() Config {
	return cfg
}

func Setup() {
	if DEBUG {
		viper.SetConfigFile("config/config_dev.yaml")
	} else {
		viper.SetConfigFile("config/config_prod.yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Panic("Config dosyası okunamadı. " + err.Error())
	}

	err := viper.Unmarshal(&cfg)
	if err != nil {
		log.Panic("Config unmarshall edilemedi.")
	}
}

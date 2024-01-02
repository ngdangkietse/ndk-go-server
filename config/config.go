package config

import (
	"github.com/spf13/viper"
	"log"
)

var EnvConfig *Config

func InitConfig() {
	EnvConfig = LoadConfig()
}

type Config struct {
	PostgresDriver  string `mapstructure:"DATABASE_DRIVER"`
	PostgresHost    string `mapstructure:"POSTGRES_HOST"`
	PostgresPort    string `mapstructure:"POSTGRES_PORT"`
	PostgresUser    string `mapstructure:"POSTGRES_USER"`
	PostgresPass    string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDB      string `mapstructure:"POSTGRES_DB"`
	PostgresSslMode string `mapstructure:"POSTGRES_SSL_MODE"`
	TimeZone        string `mapstructure:"TIME_ZONE"`
}

func LoadConfig() (config *Config) {
	viper.AddConfigPath("./config")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	return
}

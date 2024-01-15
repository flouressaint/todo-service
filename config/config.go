package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBHost       string `mapstructure:"POSTGRES_HOST"`
	DBUsername   string `mapstructure:"POSTGRES_USER"`
	DBPassword   string `mapstructure:"POSTGRES_PASSWORD"`
	DBName       string `mapstructure:"POSTGRES_DB"`
	DBPort       string `mapstructure:"POSTGRES_PORT"`
	ServerPort   string `mapstructure:"PORT"`
	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`
	JwtSignKey   string `mapstructure:"JWT_SIGN_KEY"`
}

func New(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

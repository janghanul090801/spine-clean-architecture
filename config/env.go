package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var E Env

type Env struct {
	AppEnv                 string `mapstructure:"APP_ENV"`
	ServerAddress          string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout         int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost                 string `mapstructure:"DB_HOST"`
	DBPort                 string `mapstructure:"DB_PORT"`
	DBUser                 string `mapstructure:"DB_USER"`
	DBPass                 string `mapstructure:"DB_PASS"`
	DBName                 string `mapstructure:"DB_NAME"`
	DBTlsSkipVerify        bool   `mapstructure:"DB_TLS_SKIP_VERIFY"`
	DBApplicationName      string `mapstructure:"DB_APPLICATION_NAME"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
	Debug                  bool   `mapstructure:"DEBUG"`
	LogFile                string `mapstructure:"LOG_FILE"`
}

func NewEnv() {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("can't find the file .env: %w", err))
	}

	err = viper.Unmarshal(&E)
	if err != nil {
		panic(fmt.Errorf("environment can't be loaded: %w", err))
	}

	if E.AppEnv == "development" {
		log.Println("The App is running in development env")
	}
}

package lib

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

type Env struct {
	ServerPort   string `mapstructure:"SERVER_PORT"`
	Environment  string `mapstructure:"ENV"`
	LogOutput    string `mapstructure:"LOG_OUTPUT"`
	LogLevel     string `mapstructure:"LOG_LEVEL"`
	ClientId     string `mapstructure:"CLIENTID"`
	ClientSecret string `mapstructure:"CLIENT_SECRET_KEY"`
	SwitCode     string `mapstructure:"SWIT_CODE"`
	ServerUrl    string `mapstructure:"SERVER_URL"`
	AppId        string `mapstructure:"APPID"`
}

var (
	once sync.Once
	env  Env
)

// LoadEnv loads environment variables from .env file
func LoadEnv() (Env, error) {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("☠️ cannot read configuration")
		return env, err
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("☠️ environment can't be loaded: ", err)
		return env, err
	}

	return env, nil
}

// NewEnv provides a single instance of Env, ensuring it's only loaded once
func NewEnv() Env {

	once.Do(func() {
		env, _ = LoadEnv()

	})
	return env
}

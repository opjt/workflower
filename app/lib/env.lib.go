package lib

import (
	"log"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type Env struct {
	Server ServerConfig
	Log    LogConfig
	Swit   SwitConfig
}

type ServerConfig struct {
	Port        string `mapstructure:"PORT"`
	Environment string `mapstructure:"ENV"`
	Url         string `mapstructure:"URL"`
}

type LogConfig struct {
	Output string `mapstructure:"OUTPUT"`
	Level  string `mapstructure:"LEVEL"`
}

type SwitConfig struct {
	ClientId     string `mapstructure:"CLIENT_ID"`
	ClientSecret string `mapstructure:"CLIENT_SECRET_KEY"`
	ChannelId    string `mapstructure:"CHANNEL_ID"`
	AccessToken  string `mapstructure:"ACCESS_TOKEN"`
	RefreshToken string `mapstructure:"REFRESH_TOKEN"`
	AppId        string `mapstructure:"APP_ID"`
}

var (
	once sync.Once
	env  Env
)

// LoadEnv loads environment variables from .env file
func LoadEnv() (Env, error) {
	viper.SetConfigFile(".env")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

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

func NewEnv() Env {

	once.Do(func() {
		env, _ = LoadEnv()

	})
	return env
}

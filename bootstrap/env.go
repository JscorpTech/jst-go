package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	Port        string `mapstructure:"PORT"`
	REDIS_URL   string `mapstructure:"REDIS_URL"`
	DB_USER     string `mapstructure:"DB_USER"`
	DB_PASSWORD string `mapstructure:"DB_PASSWORD"`
	DB_HOST     string `mapstructure:"DB_HOST"`
	DB_PORT     int    `mapstructure:"DB_PORT"`
	DB_DATABASE string `mapstructure:"DB_DATABASE"`
	TIMEZONE    string `mapstructure:"TIMEZONE"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")
	viper.SetDefault("TIMEZONE", "Asia/Tashkent")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	return &env
}

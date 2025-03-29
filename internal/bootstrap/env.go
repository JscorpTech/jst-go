package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	Port       string `mapstructure:"PORT"`
	RedisUrl   string `mapstructure:"REDIS_URL"`
	DbUser     string `mapstructure:"DB_USER"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbHost     string `mapstructure:"DB_HOST"`
	DbPort     int    `mapstructure:"DB_PORT"`
	DbDatabase string `mapstructure:"DB_DATABASE"`
	Timezone   string `mapstructure:"TIMEZONE"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")
	viper.SetDefault("Timezone", "Asia/Tashkent")

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

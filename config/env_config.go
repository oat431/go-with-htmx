package config

import (
	"log"

	"github.com/spf13/viper"
)

func init() {
	log.Println("Loading .env")
}

func GetEnv(key string) string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal("Error while reading .env")
	}

	value, ok := viper.Get(key).(string)

	if !ok {
		log.Fatal("Invalid type")
	}

	return value
}

package config

import (
	"log"
	"sync"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AppPort   string `envconfig:"APP_PORT"`
	SecretKey string `envconfig:"SECRET_KEY"`

	ResourceUrl             string `envconfig:"RESOURCE_URL"`
	CurrencyConverterApiKey string `envconfig:"CURRENCY_CONVERTER_API_KEY"`
	CurrencyConverterUrl    string `envconfig:"CURRENCY_CONVERTER_URL"`
}

var c Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		_ = godotenv.Load()
		err := envconfig.Process("", &c)
		if err != nil {
			log.Fatal(err.Error())
		}
	})

	return &c
}

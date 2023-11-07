package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

func LoadViper() {

	viper.AddConfigPath("/")
	viper.SetConfigFile("env.yaml")
	viper.AutomaticEnv()

	if ViperException := viper.ReadInConfig(); ViperException != nil {
		panic(ViperException)
	}

}

func LoadValidator() *validator.Validate {

	validate := validator.New()
	return validate

}

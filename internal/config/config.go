package config

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
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

func LoadTranslator(validate *validator.Validate) ut.Translator {

	en := en.New()
	english := ut.New(en, en)

	translator, _ := english.GetTranslator("en")

	en_translations.RegisterDefaultTranslations(validate, translator)

	return translator

}

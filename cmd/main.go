package main

import (
	"woman-center-be/internal/config"
	v1 "woman-center-be/internal/delivery/v1"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func main() {
	config.LoadViper()
	validator := config.LoadValidator()

	db := config.InitDatabase()

	e := echo.New()

	v1.InitRoutes(e, db, validator)

	e.Logger.Fatal(e.Start(":" + viper.GetString("APP_PORT")))

}

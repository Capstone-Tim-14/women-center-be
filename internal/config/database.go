package config

import (
	"fmt"
	"log"
	"woman-center-be/internal/app/v1/models/schema"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	dbUser := viper.GetString("DATABASE.USER")
	dbPass := viper.GetString("DATABASE.PASS")
	dbHost := viper.GetString("DATABASE.HOST")
	dbPort := viper.GetString("DATABASE.PORT")
	dbName := viper.GetString("DATABASE.NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	var errDB error
	DB, errDB := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errDB != nil {
		log.Fatal("Failed to Connect Database")
	}

	migrations(DB)

	return DB
}

func migrations(db *gorm.DB) {
	db.AutoMigrate(schema.Users{})
}

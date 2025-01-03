package helpers

import (
	"car-rental-framework/internal/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)

func SetupDB() {
	var err error
	dialect := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		GetEnv("DB_USER"),
		GetEnv("DB_PASSWORD"),
		GetEnv("DB_HOST"),
		GetEnv("DB_PORT"),
		GetEnv("DB_NAME"),
	)
	DB, err = gorm.Open(mysql.Open(dialect), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		Logger.Fatal("failed to connect database: ", err)
	}

	err = DB.AutoMigrate(&models.User{}, &models.UserSession{})
	if err != nil {
		Logger.Fatal("failed to migrate database: ", err)
	}

	Logger.Info("database connected")
}

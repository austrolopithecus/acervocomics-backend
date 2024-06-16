package util

import (
	"acervocomics/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("acervo.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	err = DB.AutoMigrate(models.User{})
	if err != nil {
		return err
	}
	return nil
}

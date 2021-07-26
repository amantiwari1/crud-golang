package config

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db = d

}

func GetDB() *gorm.DB {
	return db
}

package database

import (
	"github.com/api/src/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// Init initializes database
func Init(isReset bool, models ...interface{}) {
	c := config.GetConfig()
	var err error

	db, err := gorm.Open(c.GetString("database.provider"), c.GetString("database.url"))
	if err != nil {
		panic(err)
	}
	if isReset {
		db.DropTableIfExists(models)
	}
	db.AutoMigrate(models...)
}

// GetDB returns database connection
func GetDB() *gorm.DB {
	return db
}

// Close closes database
func Close() {
	db.Close()
}

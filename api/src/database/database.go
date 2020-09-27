package database

import (
	"github.com/api/src/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var d *gorm.DB

// Init initializes database
func Init(isReset bool, models ...interface{}) {
	c := config.GetConfig()
	var err error
	// d, err = gorm.Open(c.GetString("db.provider"), c.GetString("db.url"))

	USER := c.GetString("database.user")
	PASS := c.GetString("database.password")
	PROTOCOL := c.GetString("database.port")
	DBNAME := c.GetString("database.dbname")
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(mysql.Open(CONNECT), &gorm.Config{
		SkipDefaultTransaction: true,
	})
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
	return d
}

// Close closes database
func Close() {
	d.Close()
}

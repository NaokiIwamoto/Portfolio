package main

import (
	"flag"

	"github.com/api/src/config"
	"github.com/api/src/database"
	"github.com/api/src/migrate"
	"github.com/api/src/server"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// 設定ファイル名の指定
	env := flag.String("e", "development", "")
	flag.Parse()

	// 設定の適応
	config.Init(*env)

	// DBコネクション
	database.Init(true)
	migrate.MigrateAll()
	server.Init()
}

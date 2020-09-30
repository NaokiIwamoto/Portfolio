package server

import (
	"time"

	"github.com/api/src/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Init initialize server
func Init() error {
	conf := config.GetConfig()
	router := gin.Default()
	// r.Use() で使用したいミドルウェアなどの指定ができる。
	// https://godoc.org/github.com/gin-gonic/gin#RouterGroup.Use
	router.Use(cors.New(cors.Config{
		// 許可したいHTTPメソッドの一覧
		AllowMethods: conf.GetStringSlice("server.allowMethods"),
		// 許可したいHTTPリクエストヘッダの一覧
		AllowHeaders: conf.GetStringSlice("server.allowHeaders"),
		// 許可したいアクセス元の一覧
		AllowOrigins: conf.GetStringSlice("server.accessAllow"),
		MaxAge:       24 * time.Hour,
	}))
	NewRouter(router, conf)
	return nil
}

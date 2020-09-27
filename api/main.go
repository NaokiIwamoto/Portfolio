package main

import (
	"flag"

	"github.com/api/src/config"
	"github.com/api/src/database"
	"github.com/api/src/server"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {

	env := flag.String("e", "development", "")
	flag.Parse()

	config.Init(*env)
	database.Init(false)
	defer database.Close()
	if err := server.Init(); err != nil {
		panic(err)
	}
}

// package main

// import (
// 	"time"

// 	"github.com/api/src/Article"
// 	"github.com/gin-contrib/cors"
// 	"github.com/gin-gonic/gin"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// func main() {
// 	dbMigrate(dbConnect())
// 	r := gin.Default()
// 	// r.Use() で使用したいミドルウェアなどの指定ができる。
// 	// https://godoc.org/github.com/gin-gonic/gin#RouterGroup.Use
// 	r.Use(cors.New(cors.Config{
// 		// 許可したいHTTPメソッドの一覧
// 		AllowMethods: []string{
// 			"POST",
// 			"GET",
// 			"OPTIONS",
// 			"PUT",
// 			"DELETE",
// 		},
// 		// 許可したいHTTPリクエストヘッダの一覧
// 		AllowHeaders: []string{
// 			"Access-Control-Allow-Headers",
// 			"Content-Type",
// 			"Content-Length",
// 			"Accept-Encoding",
// 			"X-CSRF-Token",
// 			"Authorization",
// 		},
// 		// 許可したいアクセス元の一覧
// 		AllowOrigins: []string{
// 			"http://localhost:3000",
// 		},
// 		// 自分で許可するしないの処理を書きたい場合は、以下のように書くこともできる
// 		// AllowOriginFunc: func(origin string) bool {
// 		//  return origin == "https://www.example.com:8080"
// 		// },
// 		// preflight requestで許可した後の接続可能時間
// 		// https://godoc.org/github.com/gin-contrib/cors#Config の中のコメントに詳細あり
// 		MaxAge: 24 * time.Hour,
// 	}))
// 	r.GET("/ping", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"message": "test",
// 		})
// 	})
// 	r.POST("/article", Article.Article)
// 	r.Run(":8080")
// }

// // User is a TODO user
// type User struct {
// 	gorm.Model
// 	NickName string `json:"nickName"`
// }

// func dbMigrate(db *gorm.DB) *gorm.DB {
// 	db.AutoMigrate(&User{})
// 	return db
// }

// func dbConnect() *gorm.DB {
// 	USER := "root"
// 	PASS := "password"
// 	PROTOCOL := "tcp(mysql:3306)"
// 	DBNAME := "sample"
// 	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
// 	db, err := gorm.Open(mysql.Open(CONNECT), &gorm.Config{
// 		SkipDefaultTransaction: true,
// 	})
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return db
// }

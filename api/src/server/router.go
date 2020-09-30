package server

import (
	"log"

	usercontroller "github.com/api/src/controllers/userController"
	"github.com/api/src/middleware"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// NewRouter is constructor for router
func NewRouter(router *gin.Engine, conf *viper.Viper) {
	// ルーティング
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	authMiddleware := middleware.JWTStart()

	// 認証系のルーティング
	authRouter := router.Group("/auth")
	{
		authRouter.POST("/signup", usercontroller.Signup)
		// authRouter.POST("/login", usercontroller.Login)
		authRouter.POST("/login", authMiddleware.LoginHandler)
	}

	// 存在しないURLにアクセスした場合
	router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	// enable only permitted user
	// 認証済みユーザーのみ利用可能
	permittedRouter := router.Group("/name")
	permittedRouter.Use(authMiddleware.MiddlewareFunc())
	{
		permittedRouter.GET("/hello", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "hellow our user.",
			})
		})
		permittedRouter.GET("/userInfo", usercontroller.GetLoginUserInfo)
	}

	// サーバー起動
	router.Run(conf.GetString("server.port"))
}

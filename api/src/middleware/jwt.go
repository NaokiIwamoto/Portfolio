package middleware

import (
	"log"
	"time"

	"github.com/api/src/models/user"

	usercontroller "github.com/api/src/controllers/userController"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var identityKey = "id"

// JWTStart is able to create JWT Token
func JWTStart() *jwt.GinJWTMiddleware {
	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		// ログイン処理 No1
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			log.Println(loginVals.Username)
			userID := loginVals.Username
			password := loginVals.Password

			// if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
			if v := usercontroller.Login(c); v {
				return &user.User{
					Username: userID,
					Password: password,
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		// LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
		// 	c.JSON(http.StatusOK, gin.H{
		// 		"code":   http.StatusOK,
		// 		"token":  token,
		// 		"expire": expire.Format(time.RFC3339),
		// 	})
		// },
		// ログイン処理 No2 Optional
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			// 型アサーションでdataをUser型に変更
			if v, ok := data.(*user.User); ok {
				return jwt.MapClaims{
					identityKey: v.Username,
				}
			}
			return jwt.MapClaims{}
		},

		// MiddlewareFuncの後に実行されるもの Optional
		// IdentityHandler: func(c *gin.Context) interface{} {
		// 	claims := jwt.ExtractClaims(c)
		// 	log.Println(claims)
		// 	return &user.User{
		// 		Username: claims[identityKey].(string),
		// 	}
		// },
		// IdentityHandlerの後に実行される Optional
		// Authorizator: func(data interface{}, c *gin.Context) bool {
		// 	if _, ok := data.(*user.User); ok {
		// 		return true
		// 	}

		// 	return false
		// },
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	return authMiddleware
}

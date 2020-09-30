package usercontroller

import (
	"github.com/api/src/models/user"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Signup is constructer for HealthController
func Signup(c *gin.Context) {
	username := c.PostForm("username")
	password := userPassHash(c.PostForm("password"))
	// 登録ユーザーが重複していた場合にはじく処理
	if err := user.CreateUser(username, password); err != nil {
		c.JSON(400, "")
	}
	c.JSON(200, "")
}

// Login is constructer for HealthController
func Login(c *gin.Context) bool {
	username := c.PostForm("username")
	password := c.PostForm("password")
	// 登録ユーザーが重複していた場合にはじく処理
	dbpass := user.GetUser(username).Password
	return userPassMach(dbpass, password)
}

func GetLoginUserInfo(c *gin.Context) {
	username := c.Query("username")
	loginUser := user.GetUser(username)
	c.JSON(200, gin.H{
		"userID":   loginUser.UID,
		"userName": loginUser.Username,
	})
}

// パスワードのハッシュ化
func userPassHash(pass string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hash)
}

// 入力値とDB値（hash）を比較し、同じならtrue
func userPassMach(hash, pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw)) == nil
}

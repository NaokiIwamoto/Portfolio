package user

import (
	"github.com/api/src/database"
	"github.com/api/src/models"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// User is struct of user
type User struct {
	UID      uuid.UUID `gorm:"type:char(36);primary_key"`
	Username string    `form:"username"`
	Password string    `form:"password"`
	models.BaseModel
}

func (user *User) BeforeCreate(scope *gorm.DB) error {
	user.UID = uuid.NewV4()
	return nil
}

func Migrate() {
	db := database.GetDB()
	db.DropTableIfExists(&User{})
	db.AutoMigrate(&User{})
}

// CreateUser creates a user
func CreateUser(username string, password string) error {
	db := database.GetDB()
	// Insert処理
	db.Create(&User{Username: username, Password: password})
	return nil
}

func GetUser(username string) *User {
	db := database.GetDB()
	var user User
	db.First(&user, "username = ?", username)
	return &user
}

package user

import (
	"github.com/api/src/database"
	"github.com/api/src/models"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// User is struct of user
type User struct {
	UID         uuid.UUID `gorm:"type:char(36);primary_key"`
	Username    string    `form:"username"`
	MailAddress string    `form:"mail_address"`
	Password    string    `form:"password"`
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
func CreateUser(username, mailAddress, password string) error {
	db := database.GetDB()
	// Insert処理
	return db.Create(&User{Username: username, MailAddress: mailAddress, Password: password}).Error
}

func GetUser(mailAddress string) *User {
	db := database.GetDB()
	var user User
	db.First(&user, "mail_address = ?", mailAddress)
	return &user
}

package models

// User is struct of user
type User struct {
	gorm.Model
	UID          string        `json:"uid" gorm:"unique;not null"`
	Name         string        `json:"name" gorm:"unique;not null"`
	Applications []Application `json:"applications"`
	Description  string        `json:"description"`
}

// Create creates a user
func (u *User) Create() (err error) {
	db := database.GetDB()
	return db.Create(u).Error
}

// FindByID finds a user by id
func (u *User) FindByID(id uint) (err error) {
	db := database.GetDB()
	return db.Where("id = ?", id).First(u).Error
}
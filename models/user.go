package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id        int    `form:"id" json:"id" validate:"required"`
	Name      string `form:"name" json:"name" validate:"required"`
	Gmail     string `form:"gmail" json:"gmail" validate:"required"`
	Username  string `form:"username" json:"username" validate:"required"`
	Password  string `form:"password" json:"password" validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// CRUD
func CreateUser(db *gorm.DB, newUser *User) (err error) {
	err = db.Create(newUser).Error
	if err != nil {
		return err
	}
	return nil
}
func ReadUser(db *gorm.DB, users *[]User) (err error) {
	err = db.Find(users).Error
	if err != nil {
		return err
	}
	return nil
}

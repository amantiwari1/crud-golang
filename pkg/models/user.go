package models

import (
	"github.com/amantiwari1/crud-golang/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Name   string `json:"name"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (b *User) CreateUser() *User {
	db.Create(&b)

	return b
}

func GetAllUser() []User {
	var Users []User

	db.Find(&Users)

	return Users
}

func GetUserById(Id int64) (*User, *gorm.DB) {
	var getUser User

	db := db.Where("ID=?", Id).Find(&getUser)
	return &getUser, db
}

func DeleteUser(ID int64) User {
	var user User
	db.Where("ID=?", ID).Delete(&user)
	return user
}

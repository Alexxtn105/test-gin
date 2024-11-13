package models

import (
	"test-gin/helpers"
	"time"
)

type User struct {
	// TODO
	ID        uint64 `gorm:"primaryKey"`
	Username  string `gorm:"size:64"`
	Password  string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func UserCheckAvailability(email string) bool {
	var user User
	DB.Where("username = ?", email).First(&user)

	return user.ID == 0 // если ИД == 0, значит email свободен
}

func UserCreate(email string, password string) *User {
	hshPasswd, _ := helpers.HashPassword(password)
	entry := User{Username: email, Password: hshPasswd}
	DB.Create(&entry)
	return &entry
}

func UserFind(userId uint64) *User {
	var user User
	DB.Where("id = ?", userId).First(&user)
	return &user
}

// UserCheck Проверка корректности ввода пароля
func UserCheck(email string, password string) *User {
	//TODO
	return &User{}
}

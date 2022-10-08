package models

import (
	_ "gorm.io/gorm"
	"time"
)

type User struct {
	ID       uint64 `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
}

type UserErrors struct {
	Err      bool   `json:"error"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Session struct {
	User   uint64    `json:"user"`
	Expire time.Time `json:"expire"`
}

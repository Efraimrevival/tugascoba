package model

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	Id        int    `json:"id"`
	Nama      string `json:"nama" `
	Email     string `json:"email"`
	Password  string `json:"password" `
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UsersResponse struct {
	Id    int    `json:"id"`
	Nama  string `json:"nama" `
	Email string `json:"email"`
	Token string `json:"token"`
}

func (Users) TableName() string {
	return "user"
}

func (UsersResponse) TableName() string {
	return "usersresponse"
}

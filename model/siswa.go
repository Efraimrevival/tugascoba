package model

import (
	"time"

	"gorm.io/gorm"
)

type Siswa struct {
	Id        int    `json:"id" form:"id"`
	Nama      string `json:"nama" form:"nama"`
	Ttl       string `json:"ttl" form:"ttl"`
	Alamat    string `json:"alamat" form:"alamat"`
	No_wali   string `json:"noWali"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Siswa) TableName() string {
	return "siswa"
}

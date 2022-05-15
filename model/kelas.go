package model

import (
	"time"

	"gorm.io/gorm"
)

type Kelas struct {
	Id           int            `json:"id" form:"id"`
	Wali         string         `json:"wali" form:"wali"`
	Nama         string         `json:"nama" form:"nama"`
	Tahun_ajaran string         `json:"tahunAjaran" form:"tahunAjaran"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Kelas) TableName() string {
	return "kelas"
}

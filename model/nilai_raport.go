package model

import (
	"time"

	"gorm.io/gorm"
)

type Nilai_raport struct {
	Id        int `json:"id" from:"id"`
	Id_kelas  int `json:"id_kelas" from:"id_kelas"`
	Id_siswa  int `json:"id_siswa" from:"id_siswa"`
	IPA       int `gorm:"column:ipa"`
	IPS       int `gorm:"column:ips"`
	MTK       int `json:"mtk" from:"mtk"`
	PKN       int `json:"pkn" from:"pkn"`
	B_INDO    int `json:"b_indo" from:"b_indo"`
	B_INGGRIS int `json:"b_inggris" from:"b_inggris"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Nilai_raport) TableName() string {
	return "nilai_raport"
}

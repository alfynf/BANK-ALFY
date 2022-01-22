package models

import (
	"time"

	"gorm.io/gorm"
)

type Customers struct {
	gorm.Model
	NoKTP        string    `gorm:"type:varchar(16);unique;no null" json:"no_ktp" form:"no_ktp"`
	Nama         string    `gorm:"type:varchar(255);not null" json:"name" form:"name"`
	TempatLahir  string    `gorm:"type:varchar(255);not null" json:"tempat_lahir" form:"tempat_lahir"`
	TanggalLahir time.Time `gorm:"type:datetime;not null" json:"tanggal_lahir" form:"tanggal_lahir"`
	NoTelp       string    `gorm:"type:varchar(15);not null" json:"no_telp" form:"no_telp"`
}

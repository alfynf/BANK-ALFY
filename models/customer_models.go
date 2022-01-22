package models

import (
	"gorm.io/gorm"
)

type Nasabah struct {
	gorm.Model
	NoKTP        string `gorm:"type:varchar(16);unique;not null" json:"no_ktp" form:"no_ktp"`
	Nama         string `gorm:"type:varchar(255);not null" json:"nama" form:"nama"`
	Alamat       string `gorm:"type:varchar(255);not null" json:"alamat" form:"alamat"`
	Telp         string `gorm:"type:varchar(50);unique;not null" json:"telp" form:"telp"`
	TempatLahir  string `gorm:"type:varchar(255);not null" json:"tempat_lahir" form:"tempat_lahir"`
	TanggalLahir string `gorm:"type:datetime;not null" json:"tanggal_lahir" form:"tanggal_lahir"`
}

type Body struct {
	NoKTP string `json:"no_ktp" form:"no_ktp"`
}

type GetResponse struct {
	NoKTP        string `json:"no_ktp" form:"no_ktp"`
	Nama         string `json:"nama" form:"nama"`
	Alamat       string `json:"alamat" form:"alamat"`
	Telp         string `json:"telp" form:"telp"`
	TempatLahir  string `json:"tempat_lahir" form:"tempat_lahir"`
	TanggalLahir string `json:"tanggal_lahir" form:"tanggal_lahir"`
}

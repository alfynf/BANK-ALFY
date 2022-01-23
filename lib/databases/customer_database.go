package databases

import (
	"it-bni/config"
	"it-bni/models"
)

// fungsi untuk mendaftarkan nasabah baru
func CreateNasabah(nasabah *models.Nasabah) (*models.Nasabah, error) {
	tx := config.DB.Create(&nasabah)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return nasabah, nil
}

// fungsi untuk mendapatkan data nasabah dengan nomor ktp
func GetNasabahByKTP(ktp string) (interface{}, error) {
	var result models.GetResponse
	tx := config.DB.Model(&models.Nasabah{}).Where("no_ktp = ? AND deleted_at IS NULL", ktp).Find(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected < 1 {
		return nil, nil
	}
	return result, nil
}

// fungsi untuk melihat daftar semua nasabah yang sudah pernah terdaftar
func GetNasabah() (interface{}, error) {
	type DaftarNasabah struct {
		NoKTP string `json:"no_ktp" form:"no_ktp"`
		Nama  string `json:"nama" form:"nama"`
	}
	var result []DaftarNasabah
	tx := config.DB.Model(&models.Nasabah{}).Where("deleted_at IS NULL").Find(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected < 1 {
		return nil, nil
	}
	return result, nil
}

// fungsi untuk melakukan pengkinian data nasabah
func UpdateNasabah(ktp string, updateNasabah *models.Nasabah) (interface{}, error) {
	var nasabah models.Nasabah
	tx := config.DB.Model(&nasabah).Where("no_ktp = ?", ktp).Updates(&updateNasabah)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected < 1 {
		return nil, nil
	}
	var updatedNasabah models.GetResponse
	config.DB.Model(nasabah).Where("no_ktp = ?", ktp).First(&updatedNasabah)
	return updatedNasabah, nil
}

// fungsi untuk menghapus data nasabah yang sudha terdaftar di sistem
func DeleteNasabah(ktp string) (interface{}, error) {
	var nasabah models.Nasabah
	tx := config.DB.Where("no_ktp = ?", ktp).Delete(&nasabah)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected < 1 {
		return nil, nil
	}
	return nasabah, nil
}

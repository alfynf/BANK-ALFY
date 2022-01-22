package controllers

import (
	"it-bni/lib/databases"
	"it-bni/lib/responses"
	"it-bni/models"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// controller untuk mendaftarkan nasabah baru
func CreateNasabahController(c echo.Context) error {
	var body models.Nasabah
	c.Bind(&body)
	body.Nama = strings.TrimSpace(body.Nama)
	if body.Nama == "" || body.NoKTP == "" || body.Telp == "" || body.TempatLahir == "" || body.TanggalLahir == "" {
		return c.JSON(http.StatusBadRequest, responses.FailedReponse("Seluruh data harus diisi"))
	}
	if len(body.NoKTP) != 16 {
		return c.JSON(http.StatusBadRequest, responses.FailedReponse("Nomor KTP berjumlah 16 digit"))
	}
	if len(body.Telp) > 13 {
		return c.JSON(http.StatusBadRequest, responses.FailedReponse("Nomor HP tidak lebih dari 13 digit"))
	}
	if body.Telp[:2] != "08" {
		return c.JSON(http.StatusBadRequest, responses.FailedReponse("Nomor HP tidak memenuhi format nomor HP"))
	}
	// error buat yang duplicate
	duplicate, _ := databases.GetNasabahByKTP(body.NoKTP)
	if duplicate != nil {
		return c.JSON(http.StatusBadRequest, responses.FailedReponse("Nomor KTP sudah digunakan"))
	}
	body.Nama = strings.ToUpper(body.Nama)
	body.TempatLahir = strings.ToUpper(body.TempatLahir)
	_, err := databases.CreateNasabah(&body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.FailedReponse("Gagal membuat nasabah baru"))
	}
	return c.JSON(http.StatusOK, responses.SuccessResponse())
}

// fungsi untuk mendapatkan data nasabah dengan nomor ktp
func GetNasabahByKTPController(c echo.Context) error {
	var body models.Body
	c.Bind(&body)
	nasabah, err := databases.GetNasabahByKTP(body.NoKTP)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.FailedReponse("Gagal memuat informasi nasabah"))
	}
	if nasabah == nil {
		return c.JSON(http.StatusBadRequest, responses.FailedReponse("Nasabah tidak ditemukan"))
	}
	return c.JSON(http.StatusOK, responses.SuccessResponseWithData(nasabah))
}

// fungsi untuk melihat daftar semua nasabah yang sudah pernah terdaftar
func GetNasabahController(c echo.Context) error {
	nasabah, err := databases.GetNasabah()
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.FailedReponse("Gagal memuat informasi nasabah"))
	}
	if nasabah == nil {
		return c.JSON(http.StatusBadRequest, responses.FailedReponse("Nasabah tidak ditemukan"))
	}
	return c.JSON(http.StatusOK, responses.SuccessResponseWithData(nasabah))
}

// fungsi untuk melakukan pengkinian data nasabah
func UpdateNasabahController(c echo.Context) error {
	type BodyUpdate struct {
		NoKTP      string         `json:"no_ktp" form:"no_ktp"`
		UpdateData models.Nasabah `json:"update_data" form:"update_data"`
	}
	var body BodyUpdate
	c.Bind(&body)
	nasabah, err := databases.UpdateNasabah(body.NoKTP, &body.UpdateData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.FailedReponse("Gagal melakukan pengikinian informasi nasabah"))
	}
	return c.JSON(http.StatusOK, responses.SuccessResponseWithData(nasabah))
}

// fungsi untuk menghapus data nasabah yang sudha terdaftar di sistem
func DeleteNasabahController(c echo.Context) error {
	var body models.Body
	c.Bind(&body)
	_, err := databases.DeleteNasabah(body.NoKTP)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.FailedReponse("Gagal menghapus nasabah"))
	}
	return c.JSON(http.StatusOK, responses.SuccessResponse())
}

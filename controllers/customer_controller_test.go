package controllers

import (
	"bytes"
	"encoding/json"
	"it-bni/config"
	"it-bni/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// Struct yang digunakan ketika test request success, dapat menampung banyak data
type Response struct {
	Code   string
	Status string
}

// Struct untuk menampung data test case
type TestCase struct {
	Name       string
	Path       string
	ExpectCode int
}

// data dummy
var (
	mock_data_nasabah = models.Nasabah{
		NoKTP:        "3213035603990004",
		Nama:         "Alfy",
		Alamat:       "Bumi Abdi Praja",
		Telp:         "085956779388",
		TempatLahir:  "Subang",
		TanggalLahir: "1999-04-16",
	}
	mock_data_ktp = models.Body{
		NoKTP: "3213035603990004",
	}
	mock_data_update = models.BodyUpdate{
		NoKTP: "3213035603990004",
		UpdateData: models.Nasabah{
			Alamat: "Sukamelang",
		},
	}
)

// Fungsi untuk memasukkan data user test ke dalam database
func InsertMockDataToDB() error {
	query := config.DB.Save(&mock_data_nasabah)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

// inisialisasi echo
func InitEcho() *echo.Echo {
	config.InitDBTest()
	e := echo.New()

	return e
}

// Fungsi untuk melakukan testing fungsi GetNasabahController kondisi request success
func TestGetNasabahControllerSuccess(t *testing.T) {
	var testCases = []struct {
		Name       string
		Path       string
		ExpectCode int
		ExpectSize int
	}{
		{
			Name:       "success to get all nasabah",
			Path:       "/nasabah",
			ExpectCode: http.StatusOK,
			ExpectSize: 1,
		},
	}

	e := InitEcho()

	InsertMockDataToDB()

	req := httptest.NewRequest(http.MethodGet, "/nasabah", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)

	for _, testCase := range testCases {
		context.SetPath(testCase.Path)

		if assert.NoError(t, GetNasabahController(context)) {
			res_body := res.Body.String()
			var response Response
			er := json.Unmarshal([]byte(res_body), &response)
			if er != nil {
				assert.Error(t, er, "error")
			}
			assert.Equal(t, testCase.ExpectCode, res.Code)
			assert.Equal(t, "Successful Operation", response.Status)
		}
	}
}

// Fungsi untuk melakukan testing fungsi GetNasabahController kondisi request failed
func TestGetNasabahControllerFailed(t *testing.T) {
	var testCase = TestCase{
		Name:       "failed to get nasabah",
		Path:       "/nasabah",
		ExpectCode: http.StatusBadRequest,
	}

	e := InitEcho()

	// Melakukan penghapusan tabel untuk membuat request failed
	config.DB.Migrator().DropTable(&models.Nasabah{})

	req := httptest.NewRequest(http.MethodGet, "/nasabah", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath(testCase.Path)

	if assert.NoError(t, GetNasabahController(context)) {
		res_body := res.Body.String()
		var response Response
		er := json.Unmarshal([]byte(res_body), &response)
		if er != nil {
			assert.Error(t, er, "error")
		}
		assert.Equal(t, testCase.ExpectCode, res.Code)
		assert.Equal(t, "Operation Failed", response.Status)
	}
}

// Fungsi untuk melakukan testing fungsi GetNasabahController kondisi nasabah not found
func TestGetNasabahControllerNotFound(t *testing.T) {
	var testCase = TestCase{
		Name:       "nasabah not found",
		Path:       "/nasabah",
		ExpectCode: http.StatusBadRequest,
	}

	e := InitEcho()

	req := httptest.NewRequest(http.MethodGet, "/nasabah", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath(testCase.Path)

	if assert.NoError(t, GetNasabahController(context)) {
		res_body := res.Body.String()
		var response Response
		er := json.Unmarshal([]byte(res_body), &response)
		if er != nil {
			assert.Error(t, er, "error")
		}
		assert.Equal(t, testCase.ExpectCode, res.Code)
		assert.Equal(t, "Operation Failed", response.Status)
	}
}

// / Fungsi untuk melakukan testing fungsi GetNasabahByKTPController kondisi request success
func TestGetNasabahByKTPControllerSuccess(t *testing.T) {
	var testCase = TestCase{
		Name:       "success to get nasabah by no ktp",
		Path:       "/nasabah/ktp",
		ExpectCode: http.StatusOK,
	}

	e := InitEcho()

	InsertMockDataToDB()

	body, err := json.Marshal(mock_data_ktp)
	if err != nil {
		t.Error(t, err, "error")
	}

	req := httptest.NewRequest(http.MethodGet, "/nasabah/ktp", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath(testCase.Path)

	if assert.NoError(t, GetNasabahByKTPController(context)) {
		res_body := res.Body.String()
		var response Response
		er := json.Unmarshal([]byte(res_body), &response)
		if er != nil {
			assert.Error(t, er, "error")
		}
		assert.Equal(t, testCase.ExpectCode, res.Code)
		assert.Equal(t, "Successful Operation", response.Status)
	}

}

// Fungsi untuk melakukan testing fungsi GetNasabahByKTPController kondisi request failed
func TestGetNasabahByKTPControllerFailed(t *testing.T) {
	var testCase = TestCase{
		Name:       "failed to get nasabah by no ktp",
		Path:       "/nasabah/ktp",
		ExpectCode: http.StatusBadRequest,
	}

	e := InitEcho()

	// Melakukan penghapusan tabel untuk membuat request failed
	config.DB.Migrator().DropTable(&models.Nasabah{})

	body, err := json.Marshal(mock_data_ktp)
	if err != nil {
		t.Error(t, err, "error")
	}

	req := httptest.NewRequest(http.MethodGet, "/nasabah/ktp", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath(testCase.Path)

	if assert.NoError(t, GetNasabahByKTPController(context)) {
		res_body := res.Body.String()
		var response Response
		er := json.Unmarshal([]byte(res_body), &response)
		if er != nil {
			assert.Error(t, er, "error")
		}
		assert.Equal(t, testCase.ExpectCode, res.Code)
		assert.Equal(t, "Operation Failed", response.Status)
	}

}

// Fungsi untuk melakukan testing fungsi GetNasabahByKTPController kondisi nasabah not found
func TestGetNasabahByKTPControllerNotFound(t *testing.T) {
	var testCase = TestCase{
		Name:       "nasabah not found",
		Path:       "/nasabah/ktp",
		ExpectCode: http.StatusBadRequest,
	}

	e := InitEcho()

	body, err := json.Marshal(mock_data_ktp)
	if err != nil {
		t.Error(t, err, "error")
	}

	req := httptest.NewRequest(http.MethodGet, "/nasabah/ktp", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath(testCase.Path)

	if assert.NoError(t, GetNasabahByKTPController(context)) {
		res_body := res.Body.String()
		var response Response
		er := json.Unmarshal([]byte(res_body), &response)
		if er != nil {
			assert.Error(t, er, "error")
		}
		assert.Equal(t, testCase.ExpectCode, res.Code)
		assert.Equal(t, "Operation Failed", response.Status)
	}

}

// / Fungsi untuk melakukan testing fungsi DeleteNasabahController kondisi request success
func TestDeleteNasabahControllerSuccess(t *testing.T) {
	var testCase = TestCase{
		Name:       "success to delete",
		Path:       "/nasabah",
		ExpectCode: http.StatusOK,
	}

	e := InitEcho()

	InsertMockDataToDB()

	body, err := json.Marshal(mock_data_ktp)
	if err != nil {
		t.Error(t, err, "error")
	}

	req := httptest.NewRequest(http.MethodDelete, "/nasabah", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath(testCase.Path)

	if assert.NoError(t, DeleteNasabahController(context)) {
		res_body := res.Body.String()
		var response Response
		er := json.Unmarshal([]byte(res_body), &response)
		if er != nil {
			assert.Error(t, er, "error")
		}
		assert.Equal(t, testCase.ExpectCode, res.Code)
		assert.Equal(t, "Successful Operation", response.Status)
	}

}

// / Fungsi untuk melakukan testing fungsi DeleteNasabahController kondisi request failed
func TestDeleteNasabahControllerFailed(t *testing.T) {
	var testCase = TestCase{
		Name:       "failed to delete nasabah",
		Path:       "/nasabah",
		ExpectCode: http.StatusBadRequest,
	}

	e := InitEcho()

	// Melakukan penghapusan tabel untuk membuat request failed
	config.DB.Migrator().DropTable(&models.Nasabah{})

	body, err := json.Marshal(mock_data_ktp)
	if err != nil {
		t.Error(t, err, "error")
	}

	req := httptest.NewRequest(http.MethodDelete, "/nasabah", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath(testCase.Path)

	if assert.NoError(t, DeleteNasabahController(context)) {
		res_body := res.Body.String()
		var response Response
		er := json.Unmarshal([]byte(res_body), &response)
		if er != nil {
			assert.Error(t, er, "error")
		}
		assert.Equal(t, testCase.ExpectCode, res.Code)
		assert.Equal(t, "Operation Failed", response.Status)
	}

}

// / Fungsi untuk melakukan testing fungsi DeleteNasabahController kondisi request failed
func TestDeleteNasabahControllerNotFound(t *testing.T) {
	var testCase = TestCase{
		Name:       "nasabah not found",
		Path:       "/nasabah",
		ExpectCode: http.StatusBadRequest,
	}

	e := InitEcho()

	body, err := json.Marshal(mock_data_ktp)
	if err != nil {
		t.Error(t, err, "error")
	}

	req := httptest.NewRequest(http.MethodDelete, "/nasabah", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath(testCase.Path)

	if assert.NoError(t, DeleteNasabahController(context)) {
		res_body := res.Body.String()
		var response Response
		er := json.Unmarshal([]byte(res_body), &response)
		if er != nil {
			assert.Error(t, er, "error")
		}
		assert.Equal(t, testCase.ExpectCode, res.Code)
		assert.Equal(t, "Operation Failed", response.Status)
	}

}

// Fungsi untuk melakukan testing fungsi UpdateNasabahController kondisi request success
func TestUpdateNasabahControllerSuccess(t *testing.T) {
	var testCase = TestCase{
		Name:       "success to update nasabah data",
		Path:       "/nasabah",
		ExpectCode: http.StatusOK,
	}

	e := InitEcho()

	InsertMockDataToDB()

	body, err := json.Marshal(mock_data_update)
	if err != nil {
		t.Error(t, err, "error")
	}

	req := httptest.NewRequest(http.MethodPut, "/nasabah", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath(testCase.Path)

	if assert.NoError(t, UpdateNasabahController(context)) {
		res_body := res.Body.String()
		var response Response
		er := json.Unmarshal([]byte(res_body), &response)
		if er != nil {
			assert.Error(t, er, "error")
		}
		assert.Equal(t, testCase.ExpectCode, res.Code)
		assert.Equal(t, "Successful Operation", response.Status)
	}

}

// / Fungsi untuk melakukan testing fungsi UpdateNasabahController kondisi request failed
func TestUpdateNasabahControllerFailed(t *testing.T) {
	var testCase = TestCase{
		Name:       "faileds to update nasabah data",
		Path:       "/nasabah",
		ExpectCode: http.StatusBadRequest,
	}

	e := InitEcho()

	// Melakukan penghapusan tabel untuk membuat request failed
	config.DB.Migrator().DropTable(&models.Nasabah{})

	body, err := json.Marshal(mock_data_update)
	if err != nil {
		t.Error(t, err, "error")
	}

	req := httptest.NewRequest(http.MethodPut, "/nasabah", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath(testCase.Path)

	if assert.NoError(t, UpdateNasabahController(context)) {
		res_body := res.Body.String()
		var response Response
		er := json.Unmarshal([]byte(res_body), &response)
		if er != nil {
			assert.Error(t, er, "error")
		}
		assert.Equal(t, testCase.ExpectCode, res.Code)
		assert.Equal(t, "Operation Failed", response.Status)
	}

}

// / Fungsi untuk melakukan testing fungsi UpdateNasabahController kondisi request failed
func TestUpdateNasabahControllerNotFound(t *testing.T) {
	var testCase = TestCase{
		Name:       "nasabah not found",
		Path:       "/nasabah",
		ExpectCode: http.StatusBadRequest,
	}

	e := InitEcho()

	body, err := json.Marshal(mock_data_update)
	if err != nil {
		t.Error(t, err, "error")
	}

	req := httptest.NewRequest(http.MethodPut, "/nasabah", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath(testCase.Path)

	if assert.NoError(t, UpdateNasabahController(context)) {
		res_body := res.Body.String()
		var response Response
		er := json.Unmarshal([]byte(res_body), &response)
		if er != nil {
			assert.Error(t, er, "error")
		}
		assert.Equal(t, testCase.ExpectCode, res.Code)
		assert.Equal(t, "Operation Failed", response.Status)
	}

}

// Fungsi untuk melakukan testing fungsi CreateNasabahController kondisi request success
func TestCreateNasabahControllerSuccess(t *testing.T) {
	var testCase = TestCase{
		Name:       "success to create nasabah",
		Path:       "/nasabah",
		ExpectCode: http.StatusOK,
	}

	e := InitEcho()

	body, err := json.Marshal(mock_data_nasabah)
	if err != nil {
		t.Error(t, err, "error")
	}

	req := httptest.NewRequest(http.MethodPost, "/nasabah", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath(testCase.Path)

	if assert.NoError(t, CreateNasabahController(context)) {
		res_body := res.Body.String()
		var response Response
		er := json.Unmarshal([]byte(res_body), &response)
		if er != nil {
			assert.Error(t, er, "error")
		}
		assert.Equal(t, testCase.ExpectCode, res.Code)
		assert.Equal(t, "Successful Operation", response.Status)
	}

}

// Fungsi untuk melakukan testing fungsi CreateNasabahController kondisi request failed
func TestCreateNasabahControllerFailed(t *testing.T) {
	var testCase = TestCase{
		Name:       "failed to create nasabah",
		Path:       "/nasabah",
		ExpectCode: http.StatusBadRequest,
	}

	e := InitEcho()

	// Melakukan penghapusan tabel untuk membuat request failed
	config.DB.Migrator().DropTable(&models.Nasabah{})

	body, err := json.Marshal(mock_data_nasabah)
	if err != nil {
		t.Error(t, err, "error")
	}

	req := httptest.NewRequest(http.MethodPost, "/nasabah", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath(testCase.Path)

	if assert.NoError(t, CreateNasabahController(context)) {
		res_body := res.Body.String()
		var response Response
		er := json.Unmarshal([]byte(res_body), &response)
		if er != nil {
			assert.Error(t, er, "error")
		}
		assert.Equal(t, testCase.ExpectCode, res.Code)
		assert.Equal(t, "Operation Failed", response.Status)
	}

}

// Fungsi untuk melakukan testing fungsi CreateNasabahController kondisi nama tidak seluruhnya
func TestCreateNasabahControllerNullData(t *testing.T) {
	var testCase = TestCase{
		Name:       "there is null data",
		Path:       "/nasabah",
		ExpectCode: http.StatusBadRequest,
	}

	e := InitEcho()
	mock_data := models.Nasabah{
		NoKTP:        "3213035603990004",
		Alamat:       "Bumi Abdi Praja",
		Telp:         "085956779388",
		TempatLahir:  "Subang",
		TanggalLahir: "1999-04-16",
	}

	body, err := json.Marshal(mock_data)
	if err != nil {
		t.Error(t, err, "error")
	}

	req := httptest.NewRequest(http.MethodPost, "/nasabah", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath(testCase.Path)

	if assert.NoError(t, CreateNasabahController(context)) {
		res_body := res.Body.String()
		var response Response
		er := json.Unmarshal([]byte(res_body), &response)
		if er != nil {
			assert.Error(t, er, "error")
		}
		assert.Equal(t, testCase.ExpectCode, res.Code)
		assert.Equal(t, "Operation Failed", response.Status)
	}

}

// Fungsi untuk melakukan testing fungsi CreateNasabahController kondisi KTP tidak 16 digit
func TestCreateNasabahControllerWrongKTP(t *testing.T) {
	var testCase = TestCase{
		Name:       "KTP ID must be 16 digit",
		Path:       "/nasabah",
		ExpectCode: http.StatusBadRequest,
	}

	e := InitEcho()
	mock_data := models.Nasabah{
		NoKTP:        "321303560399000454",
		Nama:         "Alfy",
		Alamat:       "Bumi Abdi Praja",
		Telp:         "085956779388",
		TempatLahir:  "Subang",
		TanggalLahir: "1999-04-16",
	}

	body, err := json.Marshal(mock_data)
	if err != nil {
		t.Error(t, err, "error")
	}

	req := httptest.NewRequest(http.MethodPost, "/nasabah", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath(testCase.Path)

	if assert.NoError(t, CreateNasabahController(context)) {
		res_body := res.Body.String()
		var response Response
		er := json.Unmarshal([]byte(res_body), &response)
		if er != nil {
			assert.Error(t, er, "error")
		}
		assert.Equal(t, testCase.ExpectCode, res.Code)
		assert.Equal(t, "Operation Failed", response.Status)
	}

}

// Fungsi untuk melakukan testing fungsi CreateNasabahController kondisi No HP melebihi 13 digit
func TestCreateNasabahControllerWrongPhoneDigit(t *testing.T) {
	var testCase = TestCase{
		Name:       "Phone number must not more than 13 digit",
		Path:       "/nasabah",
		ExpectCode: http.StatusBadRequest,
	}

	e := InitEcho()
	mock_data := models.Nasabah{
		NoKTP:        "3213035603990004",
		Nama:         "Alfy",
		Alamat:       "Bumi Abdi Praja",
		Telp:         "085956779388888",
		TempatLahir:  "Subang",
		TanggalLahir: "1999-04-16",
	}

	body, err := json.Marshal(mock_data)
	if err != nil {
		t.Error(t, err, "error")
	}

	req := httptest.NewRequest(http.MethodPost, "/nasabah", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath(testCase.Path)

	if assert.NoError(t, CreateNasabahController(context)) {
		res_body := res.Body.String()
		var response Response
		er := json.Unmarshal([]byte(res_body), &response)
		if er != nil {
			assert.Error(t, er, "error")
		}
		assert.Equal(t, testCase.ExpectCode, res.Code)
		assert.Equal(t, "Operation Failed", response.Status)
	}

}

// Fungsi untuk melakukan testing fungsi CreateNasabahController kondisi No HP tidak sesuai format
func TestCreateNasabahControllerWrongPhoneFormat(t *testing.T) {
	var testCase = TestCase{
		Name:       "wrong phone number format",
		Path:       "/nasabah",
		ExpectCode: http.StatusBadRequest,
	}

	e := InitEcho()
	mock_data := models.Nasabah{
		NoKTP:        "3213035603990004",
		Nama:         "Alfy",
		Alamat:       "Bumi Abdi Praja",
		Telp:         "185956779388",
		TempatLahir:  "Subang",
		TanggalLahir: "1999-04-16",
	}

	body, err := json.Marshal(mock_data)
	if err != nil {
		t.Error(t, err, "error")
	}

	req := httptest.NewRequest(http.MethodPost, "/nasabah", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath(testCase.Path)

	if assert.NoError(t, CreateNasabahController(context)) {
		res_body := res.Body.String()
		var response Response
		er := json.Unmarshal([]byte(res_body), &response)
		if er != nil {
			assert.Error(t, er, "error")
		}
		assert.Equal(t, testCase.ExpectCode, res.Code)
		assert.Equal(t, "Operation Failed", response.Status)
	}

}

// Fungsi untuk melakukan testing fungsi CreateNasabahController kondisi No KTP sudah digunakan
func TestCreateNasabahControllerDuplicateKTP(t *testing.T) {
	var testCase = TestCase{
		Name:       "duplicate KTP id",
		Path:       "/nasabah",
		ExpectCode: http.StatusBadRequest,
	}

	e := InitEcho()
	InsertMockDataToDB()

	mock_data := models.Nasabah{
		NoKTP:        "3213035603990004",
		Nama:         "Alfy",
		Alamat:       "Bumi Abdi Praja",
		Telp:         "085956779888",
		TempatLahir:  "Subang",
		TanggalLahir: "1999-04-16",
	}

	body, err := json.Marshal(mock_data)
	if err != nil {
		t.Error(t, err, "error")
	}

	req := httptest.NewRequest(http.MethodPost, "/nasabah", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath(testCase.Path)

	if assert.NoError(t, CreateNasabahController(context)) {
		res_body := res.Body.String()
		var response Response
		er := json.Unmarshal([]byte(res_body), &response)
		if er != nil {
			assert.Error(t, er, "error")
		}
		assert.Equal(t, testCase.ExpectCode, res.Code)
		assert.Equal(t, "Operation Failed", response.Status)
	}

}

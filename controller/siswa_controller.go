package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"eraport/config"
	"eraport/model"

	"github.com/labstack/echo/v4"
)

// request GET 'http://127.0.0.1:8080/siswa/:1'

type SiswaResponse struct {
	Name string

	Raport []RaportResponse
}
type RaportResponse struct {
	NamaKelas   string
	Wali        string
	TahunAjaran string
	IPA         int
	IPS         int
	PKN         int
	BIndo       int
	BInggris    int
}

func GetSiswaControllercode(c echo.Context) error {
	// id := c.Param("code")
	id, err := strconv.Atoi(c.Param("code"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid id")
	}
	var siswa *model.Siswa
	config.DB.Where("Id=?", id).Find(&siswa)
	response := SiswaResponse{
		Name:   siswa.Nama,
		Raport: []RaportResponse{},
	}

	var raports []*model.Nilai_raport
	config.DB.Where("Id_siswa= ?", id).Find(&raports)
	for _, raport := range raports {
		var kelas *model.Kelas
		config.DB.Where("Id=?", raport.Id_kelas).Find(&kelas)
		response.Raport = append(response.Raport, RaportResponse{
			NamaKelas:   kelas.Nama,
			Wali:        kelas.Wali,
			TahunAjaran: kelas.Tahun_ajaran,
			IPA:         raport.IPA,
			IPS:         raport.IPS,
			PKN:         raport.PKN,
			BIndo:       raport.B_INDO,
			BInggris:    raport.B_INGGRIS,
		})
	}
	return c.JSON(http.StatusOK, response)
}

// request GET 'http://127.0.0.1:8080/siswa/
func GetAllSiswaController(c echo.Context) error {
	var siswa []model.Siswa
	if err := config.DB.Find(&siswa).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusOK, siswa)
}

// request POST 'http://127.0.0.1:8080/siswa/'
func CreateSiswacontroller(c echo.Context) error {
	siswa := model.Siswa{}
	// fmt.Printf("siswa sebelum bind %#v\n", siswa)
	if err := c.Bind(&siswa); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	// fmt.Printf("siswa setelah bind %#v\n", siswa)
	fmt.Printf("Before insert: %#v\n", siswa)
	if err := config.DB.Save(&siswa).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusOK, siswa)
}

// request PUT 'http://127.0.0.1:8080/siswa/:1'
func UpdateSiswaController(c echo.Context) error {
	siswaId, err := strconv.Atoi(c.Param("code"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid id")
	}
	fmt.Println("Isi siswaId ", siswaId)
	var siswa model.Siswa
	fmt.Printf("Isi siswa sebelum select %#v\n", siswa)
	if err := config.DB.First(&siswa, siswaId).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	if siswa.Id == 0 {
		return c.String(http.StatusNotFound, "siswa tidak di temukan")
	}
	fmt.Printf("Isi siswa setelah select %#v\n", siswa)
	if err := c.Bind(&siswa); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	fmt.Printf("Isi siswa setelah bind %#v\n", siswa)
	fmt.Printf("Before update: %#v\n", siswa)
	if err := config.DB.Save(&siswa).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusOK, siswa)
}

// request DELETE 'http://127.0.0.1:8080/siswa/4'
func DeleteSiswaController(c echo.Context) error {
	siswaId, err := strconv.Atoi(c.Param("code"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid id")
	}
	var raports []*model.Nilai_raport
	config.DB.Where("Id_siswa= ?", siswaId).Find(&raports)
	for _, raport := range raports {
		config.DB.Delete(raport)
	}
	var siswa model.Siswa
	if err := config.DB.First(&siswa, siswaId).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	if siswa.Id == 0 {
		return c.String(http.StatusNotFound, "siswa not found")
	}
	if err := config.DB.Delete(&siswa).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusOK, siswa)
}

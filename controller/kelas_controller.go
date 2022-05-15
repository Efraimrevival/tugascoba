package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"eraport/config"
	"eraport/model"
	"github.com/labstack/echo/v4"
)

// request GET 'http://127.0.0.1:8080/kelas/:1'
func GetKelasControllerCode(c echo.Context) error {
	kelasId, err := strconv.Atoi(c.Param("code"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid id")
	}
	var kelas model.Kelas
	if err := config.DB.First(&kelas, kelasId).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	if kelas.Id == 0 {
		return c.String(http.StatusNotFound, "kelas not found")
	}
	return c.JSON(http.StatusOK, kelas)
}

// request GET 'http://127.0.0.1:8080/kelas/'
func GetAllKelasController(c echo.Context) error {
	var kelas []model.Kelas
	if err := config.DB.Find(&kelas).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusOK, kelas)
}

// request POST 'http://127.0.0.1:8080/kelas/'
func CreateKelasController(c echo.Context) error {
	kelas := model.Kelas{}
	// fmt.Printf("kelas sebelum bind %#v\n", kelas)
	if err := c.Bind(&kelas); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	// fmt.Printf("kelas setelah bind %#v\n", kelas)
	fmt.Printf("Before insert: %#v\n", kelas)
	if err := config.DB.Save(&kelas).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusOK, kelas)
}

// request PUT 'http://127.0.0.1:8080/kelas/:1'
func UpdateKelasController(c echo.Context) error {
	kelasId, err := strconv.Atoi(c.Param("code"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid id")
	}
	fmt.Println("Isi kelaskId ", kelasId)
	var kelas model.Kelas
	fmt.Printf("Isi kelask sebelum select %#v\n", kelas)
	if err := config.DB.First(&kelas, kelasId).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	if kelas.Id == 0 {
		return c.String(http.StatusNotFound, "kelask not found")
	}
	fmt.Printf("Isi kelask setelah select %#v\n", kelas)
	if err := c.Bind(&kelas); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	fmt.Printf("Isi kelas setelah bind %#v\n", kelas)
	fmt.Printf("Before update: %#v\n", kelas)
	if err := config.DB.Save(&kelas).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusOK, kelas)
}

// request DELETE 'http://127.0.0.1:8080/kelas/4'
func DeleteKelasController(c echo.Context) error {
	kelasId, err := strconv.Atoi(c.Param("code"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid id")
	}
	var raports []*model.Nilai_raport
	config.DB.Where("Id_kelas= ?", kelasId).Find(&raports)
	for _, raport := range raports {
		config.DB.Delete(raport)
	}
	var kelas model.Kelas
	if err := config.DB.First(&kelas, kelasId).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	if kelas.Id == 0 {
		return c.String(http.StatusNotFound, "kelas not found")
	}
	if err := config.DB.Delete(&kelas).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusOK, kelas)
}

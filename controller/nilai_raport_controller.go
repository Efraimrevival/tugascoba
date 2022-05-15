package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"eraport/config"
	"eraport/model"

	"github.com/labstack/echo/v4"
)

// request GET 'http://127.0.0.1:8080/nilai_raport/:1'

func GetNilai_raportControllerCode(c echo.Context) error {
	nilai_raportId, err := strconv.Atoi(c.Param("code"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid id")
	}
	var nilai_raport model.Nilai_raport
	if err := config.DB.First(&nilai_raport, nilai_raportId).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	if nilai_raport.Id == 0 {
		return c.String(http.StatusNotFound, "Nilai_raport not found")
	}
	return c.JSON(http.StatusOK, nilai_raport)
}

// request GET 'http://127.0.0.1:8080/nilai_raport/'
func GetAllNilai_raportController(c echo.Context) error {
	var nilai_raport []model.Nilai_raport
	if err := config.DB.Find(&nilai_raport).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusOK, nilai_raport)
}

 
// request POST 'http://127.0.0.1:8080/nilai_raport/'
func CreateNilai_raportController(c echo.Context) error {
	nilai_raport := model.Nilai_raport{}
	// fmt.Printf("Nilai_raport sebelum bind %#v\n", Nilai_raport)
	if err := c.Bind(&nilai_raport); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	// fmt.Printf("Nilai_raport setelah bind %#v\n", Nilai_raport)
	fmt.Printf("Before insert: %#v\n", nilai_raport)
	if err := config.DB.Save(&nilai_raport).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusOK, nilai_raport)
}

// request PUT 'http://127.0.0.1:8080/nilai_raport/:1'
func UpdateNilai_raportController(c echo.Context) error {
	nilai_raportId, err := strconv.Atoi(c.Param("code"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid id")
	}
	fmt.Println("Isi nilai_raportkId ", nilai_raportId)
	var nilai_raport model.Nilai_raport
	fmt.Printf("Isi nilai_raportk sebelum select %#v\n", nilai_raport)
	if err := config.DB.First(&nilai_raport, nilai_raportId).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	if nilai_raport.Id == 0 {
		return c.String(http.StatusNotFound, "nilai_raportk not found")
	}
	fmt.Printf("Isi nilai_raportk setelah select %#v\n", nilai_raport)
	if err := c.Bind(&nilai_raport); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	fmt.Printf("Isi nilai_raport setelah bind %#v\n", nilai_raport)
	fmt.Printf("Before update: %#v\n", nilai_raport)
	if err := config.DB.Save(&nilai_raport).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusOK, nilai_raport)
}

// request DELETE 'http://127.0.0.1:8080/nilai_raport/4'
func DeleteNilai_raportController(c echo.Context) error {
	nilai_raportId, err := strconv.Atoi(c.Param("code"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid id")
	}
	var nilai_raport model.Nilai_raport
	if err := config.DB.First(&nilai_raport, nilai_raportId).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	if nilai_raport.Id == 0 {
		return c.String(http.StatusNotFound, "nilai_raport not found")
	}
	if err := config.DB.Delete(&nilai_raport).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusOK, nilai_raport)
}

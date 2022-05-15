package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"eraport/config"
	"eraport/model"
	"eraport/middleware"
	
	"github.com/labstack/echo/v4"
)

// request GET 'http://127.0.0.1:8080/user/:1'
func GetUserControllerCode(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("code"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid id")
	}
	var user model.Users
	if err := config.DB.First(&user, userId).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	if user.Id == 0 {
		return c.String(http.StatusNotFound, "user not found")
	}
	return c.JSON(http.StatusOK, user)
}

// request GET 'http://127.0.0.1:8080/user/'
func GetAllUserController(c echo.Context) error {
	var user []model.Users
	if err := config.DB.Find(&user).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusOK, user)
}

// request POST 'http://127.0.0.1:8080/user/'
func CreateUserController(c echo.Context) error {
	user := model.Users{}
	// fmt.Printf("user sebelum bind %#v\n", user)
	if err := c.Bind(&user); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	// fmt.Printf("user setelah bind %#v\n", user)
	fmt.Printf("Before insert: %#v\n", user)
	if err := config.DB.Save(&user).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusOK, user)
}

// request PUT 'http://127.0.0.1:8080/user/:1'
func UpdateUserController(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("code"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid id")
	}
	fmt.Println("Isi userkId ", userId)
	var user model.Users
	fmt.Printf("Isi userk sebelum select %#v\n", user)
	if err := config.DB.First(&user, userId).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	if user.Id == 0 {
		return c.String(http.StatusNotFound, "userk not found")
	}
	fmt.Printf("Isi userk setelah select %#v\n", user)
	if err := c.Bind(&user); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	fmt.Printf("Isi user setelah bind %#v\n", user)
	fmt.Printf("Before update: %#v\n", user)
	if err := config.DB.Save(&user).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusOK, user)
}

// request DELETE 'http://127.0.0.1:8080/user/4'
func DeleteUserController(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("code"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid id")
	}
	var raports []*model.Nilai_raport
	config.DB.Where("Id_user= ?", userId).Find(&raports)
	for _, raport := range raports {
		config.DB.Delete(raport)
	}
	var user model.Users
	if err := config.DB.First(&user, userId).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	if user.Id == 0 {
		return c.String(http.StatusNotFound, "user not found")
	}
	if err := config.DB.Delete(&user).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusOK, user)
}

// request POST 'http://127.0.0.1:8080/user/'
func LoginUserController(c echo.Context) error {
	user := model.Users{}
	// fmt.Printf("user sebelum bind %#v\n", user)
	if err := c.Bind(&user); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	// fmt.Printf("user setelah bind %#v\n", user)
	fmt.Printf("Before insert: %#v\n", user)
	if err := config.DB.Where("email=? AND password=?", user.Email, user.Password).First(&user).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "gagal login")
	}

	token,err := middleware.CreateToken(user.Id, user.Nama)

	if (err != nil){
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "gagal login")
	}
	userResponse := model.UsersResponse{user.Id, user.Nama,user.Email, token }

	return c.JSON(http.StatusOK, userResponse,)
}


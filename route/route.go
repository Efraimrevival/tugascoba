package route

import (
	"eraport/constant"
	"eraport/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	//users
	e.POST("/login/", controller.LoginUserController)
	//jwt_user
	ejwt := e.Group("")
	ejwt.Use(middleware.JWT([]byte(constant.SECRET_JWT)))
	ejwt.GET("/user/", controller.GetAllUserController)
	ejwt.GET("/user/:code", controller.GetUserControllerCode)
	ejwt.POST("/user/", controller.CreateUserController)
	ejwt.PUT("/user/:code", controller.UpdateUserController)
	ejwt.DELETE("/user/:code", controller.DeleteUserController)


	//Siswa
	ejwt.GET("/siswa/:code", controller.GetSiswaControllercode)
	e.GET("/siswa/", controller.GetAllSiswaController)
	e.POST("/siswa/", controller.CreateSiswacontroller)
	e.PUT("/siswa/:code", controller.UpdateSiswaController)
	e.DELETE("/siswa/:code", controller.DeleteSiswaController)
	//jwt_Siswa


	//kelas
	e.GET("/kelas/:code", controller.GetKelasControllerCode)
	e.GET("/kelas/", controller.GetAllKelasController)
	e.POST("/kelas/", controller.CreateKelasController)
	e.PUT("/kelas/:code", controller.UpdateKelasController)
	e.DELETE("/kelas/:code", controller.DeleteKelasController)
	//Nilai
	e.GET("/nilai_raport/:code", controller.GetNilai_raportControllerCode)
	e.GET("/nilai_raport/", controller.GetAllNilai_raportController)
	e.POST("/nilai_raport/", controller.CreateNilai_raportController)
	e.PUT("/nilai_raport/:code", controller.UpdateNilai_raportController)
	e.DELETE("/nilai_raport/:code", controller.DeleteNilai_raportController)
	return e
}

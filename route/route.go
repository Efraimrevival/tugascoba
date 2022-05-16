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
	e.POST("/singup/", controller.CreateUserController)
	e.POST("/login/", controller.LoginUserController)
	//jwt_user
	ejwt := e.Group("")
	ejwt.Use(middleware.JWT([]byte(constant.SECRET_JWT)))
	ejwt.GET("/user/", controller.GetAllUserController)
	ejwt.GET("/user/:code", controller.GetUserControllerCode)
	
	ejwt.PUT("/user/:code", controller.UpdateUserController)
	ejwt.DELETE("/user/:code", controller.DeleteUserController)


	//Siswa
	ejwt.GET("/siswa/:code", controller.GetSiswaControllercode)
	ejwt.GET("/siswa/", controller.GetAllSiswaController)
	ejwt.POST("/siswa/", controller.CreateSiswacontroller)
	ejwt.PUT("/siswa/:code", controller.UpdateSiswaController)
	ejwt.DELETE("/siswa/:code", controller.DeleteSiswaController)
	//jwt_Siswa


	//kelas
	ejwt.GET("/kelas/:code", controller.GetKelasControllerCode)
	ejwt.GET("/kelas/", controller.GetAllKelasController)
	ejwt.POST("/kelas/", controller.CreateKelasController)
	ejwt.PUT("/kelas/:code", controller.UpdateKelasController)
	ejwt.DELETE("/kelas/:code", controller.DeleteKelasController)
	//Nilai
	ejwt.GET("/nilai_raport/:code", controller.GetNilai_raportControllerCode)
	ejwt.GET("/nilai_raport/", controller.GetAllNilai_raportController)
	ejwt.POST("/nilai_raport/", controller.CreateNilai_raportController)
	ejwt.PUT("/nilai_raport/:code", controller.UpdateNilai_raportController)
	ejwt.DELETE("/nilai_raport/:code", controller.DeleteNilai_raportController)
	return e
}

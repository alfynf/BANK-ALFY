package routes

import (
	"it-bni/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {

	e := echo.New()

	e.POST("/register", controllers.CreateNasabahController)
	e.GET("/nasabah/ktp", controllers.GetNasabahByKTPController)
	e.GET("/nasabah", controllers.GetNasabahController)
	e.PUT("/nasabah", controllers.UpdateNasabahController)
	e.DELETE("/nasabah", controllers.DeleteNasabahController)

	// User Sign Up & Sign In

	// JWT Group
	// r := e.Group("/jwt")
	// r.Use(m.JWT([]byte(constant.SECRET_JWT)))

	// Users JWT
	// r.GET("/users", controllers.GetUserControllers)

	return e
}

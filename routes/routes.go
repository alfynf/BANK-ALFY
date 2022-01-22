package routes

import (
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {

	e := echo.New()

	// User Sign Up & Sign In
	// e.POST("/signup", controllers.CreateUserControllers)

	// JWT Group
	r := e.Group("/jwt")
	// r.Use(m.JWT([]byte(constant.SECRET_JWT)))

	// Users JWT
	// r.GET("/users", controllers.GetUserControllers)

	return e
}

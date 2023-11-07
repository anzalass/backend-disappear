package routes

import (
	"github.com/capstone-kelompok-7/backend-disappear/middlewares"
	"github.com/capstone-kelompok-7/backend-disappear/module/auth"
	"github.com/capstone-kelompok-7/backend-disappear/module/product"
	"github.com/capstone-kelompok-7/backend-disappear/module/users"
	"github.com/capstone-kelompok-7/backend-disappear/utils"
	"github.com/labstack/echo/v4"
)

func RouteAuth(e *echo.Echo, h auth.HandlerAuthInterface) {
	e.POST("api/v1/auth/register", h.Register())
	e.POST("api/v1/auth/login", h.Login())
}

func RouteUser(e *echo.Echo, h users.HandlerUserInterface, jwtService utils.JWTInterface, userService users.ServiceUserInterface) {
	users := e.Group("api/v1/users")
	users.GET("/list", h.GetAllUsers())
	users.GET("/by-email", h.GetUsersByEmail(), middlewares.AuthMiddleware(jwtService, userService))
	users.GET("/:id", h.GetUsersById(), middlewares.AuthMiddleware(jwtService, userService))
}

func RouteProduct(e *echo.Echo, h product.HandlerProductInterface) {
	products := e.Group("api/v1")
	products.GET("/products", h.GetAllProducts())
}

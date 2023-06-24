package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/observation-system/observation-backend/log"
	"github.com/observation-system/observation-backend/infrastucture"
	"github.com/observation-system/observation-backend/presentation/controller"
	customMiddleware "github.com/observation-system/observation-backend/presentation/middleware"
	_ "github.com/observation-system/observation-backend/docs/swagger"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Init() *echo.Echo {
	userController := controller.NewUserController(infrastucture.NewSqlHandler())

	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{ Output: log.GeneratApiLog() }))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// auth
	a := e.Group("/auth")
	a.POST("/user_register", func(c echo.Context) error { return userController.Register(c) }) // auth/user_register
	a.POST("/user_login", func(c echo.Context) error { return userController.Login(c) }) // auth/user_login

	// user
	u := e.Group("/user")
	u.Use(customMiddleware.UserMiddleware)
	u.GET("/user_check", func(c echo.Context) error { return userController.Check(c) }) // user/user_login
	u.DELETE("/user_delete/:userKey", func(c echo.Context) error { return userController.Delete(c) }) // user/user_delete

	e.Logger.Fatal(e.Start(":8000"))

	return e
}

package app

import (
	"framework/app/structure"
	"framework/framework"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

var (
	utils     framework.Utils
	webserver framework.WebServer
)

func ConfigRoute() *echo.Echo {
	echo.NotFoundHandler = func(c echo.Context) error {
		return webserver.ResultAPI(c, http.StatusNotFound, "Route not Valid", "")
	}
	route := echo.New()

	if structure.SystemConf.ServiceLog {
		webserver.SetupLogFile("log", "data.log", "error.log")
		route.Use(webserver.Logger())
	}
	route.Use(middleware.Recover())

	route.GET("/", appTest)
	initializeModule(route)

	return route
}

func appTest(ctx echo.Context) error {
	return webserver.ResultAPI(ctx, 200, "OK", "")
}

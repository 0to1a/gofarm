package app

import (
	"framework/app/controller"
	"framework/app/structure"
	"framework/framework/webserver"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strings"
)

func ConfigRoute() *echo.Echo {
	echo.NotFoundHandler = func(c echo.Context) error {
		return webserver.ResultAPI(c, http.StatusNotFound, "Route not Valid", "")
	}
	route := echo.New()

	if structure.SystemConf.ServiceLog {
		webserver.SetLogFile("log", "data.log", "error.log")
		route.Use(webserver.Logger())
	}
	route.Use(middleware.Recover())

	route.GET("/", controller.AppTest)
	v2 := route.Group("/v1", AuthUserAPI)
	{
		v2.POST("/", controller.AppTest)
	}

	return route
}

func AuthUserAPI(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := c.Request().Header.Get("Authorization")
		bearerToken := strings.Split(header, " ")
		if len(bearerToken) != 2 {
			return webserver.ResultAPI(c, http.StatusBadRequest, "No Authorization", "")
		}
		if bearerToken[0] != "Bearer" {
			return webserver.ResultAPI(c, http.StatusBadRequest, "Error Authorization", "")
		}

		token, ok := webserver.JWTCheckToken(bearerToken[1])
		if !ok {
			return webserver.ResultAPI(c, http.StatusUnauthorized, "Unauthorized", "")
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			username, ok := claims["username"].(string)
			if !ok {
				return webserver.ResultAPI(c, http.StatusUnauthorized, "Unauthorized", "")
			}

			// TODO: Check "username" exist

			c.Set("username", username)
		}
		return next(c)
	}
}

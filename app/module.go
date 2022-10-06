package app

import (
	"encoding/json"
	"framework/app/calculateModule"
	"framework/app/exampleModule"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strings"
)

func initializeModule(route *echo.Echo) {
	route.Use(setCORS)

	utils.UseModule(calculateModule.InitializeModule(route, authUserAPI))
	utils.UseModule(exampleModule.InitializeModule(route, authUserAPI))
}

func setCORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return next(c)
	}
}

func authUserAPI(next echo.HandlerFunc) echo.HandlerFunc {
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
			userTmp, ok := claims["user"]
			if !ok {
				return webserver.ResultAPI(c, http.StatusUnauthorized, "Unauthorized", "2")
			}

			jsonTmp, _ := json.Marshal(userTmp)
			userResult := new(exampleModule.UserAccess)
			if err := json.Unmarshal(jsonTmp, &userResult); err != nil {
				log.Println("JSON:", err)
				return webserver.ResultAPI(c, http.StatusUnauthorized, "Unauthorized", "3")
			}

			c.Set("userid", userResult.UserId)
			c.Set("user", userResult)
		}
		return next(c)
	}
}

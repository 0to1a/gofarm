package controller

import (
	"framework/framework/webserver"
	"github.com/labstack/echo/v4"
)

func AppTest(context echo.Context) error {
	return webserver.ResultAPI(context, 200, "OK", "")
}

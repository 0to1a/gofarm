package controller

import (
	"fmt"
	"framework/app/model"
	"framework/framework/webserver"
	"github.com/labstack/echo/v4"
	"net/http"
)

func AppTestExample(ctx echo.Context) error {
	return webserver.ResultAPI(ctx, 200, "OK", "")
}

func AppLoginExample(ctx echo.Context) error {
	type inputJSON struct {
		Apikey string `json:"api_key"`
	}

	inJSON := new(inputJSON)
	if err := ctx.Bind(inJSON); err != nil {
		_ = webserver.ResultAPI(ctx, http.StatusBadRequest, "Input not valid", "json")
		return fmt.Errorf("400:nv")
	}

	id, username := model.ExampleGetApikey(inJSON.Apikey)
	if id == nil {
		_ = webserver.ResultAPI(ctx, http.StatusUnauthorized, "Not Authorized", "api_key")
		return fmt.Errorf("401:na")
	}

	token, err := webserver.JWTCreateToken(username, 15)
	if err != nil {
		_ = webserver.ResultAPI(ctx, http.StatusInternalServerError, "Server Error", "")
		return fmt.Errorf("500:se")
	}
	return webserver.ResultAPI(ctx, 200, "OK", token)
}

func AppTestHello(ctx echo.Context) error {
	return webserver.ResultAPI(ctx, 200, "OK", ctx.Get("username"))
}

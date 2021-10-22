package controller

import (
	"framework/framework/database"
	"framework/framework/webserver"
	"github.com/labstack/echo/v4"
	"time"
)

func RedisExampleTime(ctx echo.Context) error {
	randomNumber := time.Now()

	if isCache, result := database.RedisCacheJsonRead(ctx); isCache {
		return webserver.ResultAPIFromJson(ctx, result)
	}

	result := webserver.ResponseAPI(200, "OK", randomNumber)
	database.RedisCacheJson(ctx, 1, result)
	return webserver.ResultAPIFromJson(ctx, result)
}

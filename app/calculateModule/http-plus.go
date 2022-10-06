package calculateModule

import (
	"github.com/labstack/echo/v4"
)

func (w *httpStructure) plus(ctx echo.Context) error {
	numb1 := ctx.Get("number1").(int64)
	numb2 := ctx.Get("number2").(int64)

	total := service.plus(numb1, numb2)

	return webserver.ResultAPI(ctx, 200, "OK", total)
}

package calculateModule

import (
	"github.com/labstack/echo/v4"
)

func (w *httpStructure) multiple(ctx echo.Context) error {
	numb1 := ctx.Get("number1").(int64)
	numb2 := ctx.Get("number2").(int64)

	total := service.multiple(numb1, numb2)

	return webserver.ResultAPI(ctx, 200, "OK", total)
}

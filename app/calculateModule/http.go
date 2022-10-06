package calculateModule

import (
	"github.com/labstack/echo/v4"
)

func httpRoute(route *echo.Echo) {
	root := route.Group("/calculate")
	{
		root.GET("/calculatePlus/:number1/:number2", http.plus)
		root.GET("/calculateMultiple/:number1/:number2", http.multiple)
	}
}

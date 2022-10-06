package calculateModule

import (
	"github.com/labstack/echo/v4"
	"log"
)

const (
	nameModule    = "Calculate Module"
	versionModule = "1.00"
)

func InitializeModule(listModule []string, route *echo.Echo, authMiddleware echo.MiddlewareFunc) (name string) {
	if route != nil {
		httpRoute(route)
	}

	log.Println(">> Attach:", nameModule, versionModule)
	return nameModule
}

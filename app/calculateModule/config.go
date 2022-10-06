package calculateModule

import (
	"framework/app/structure"
	"github.com/labstack/echo/v4"
	"log"
)

const (
	nameModule    = "Calculate Module"
	versionModule = 100
)

func InitializeModule(route *echo.Echo, authMiddleware echo.MiddlewareFunc) structure.ModularStruct {
	config := structure.ModularStruct{
		Name:      nameModule,
		Version:   versionModule,
		Depending: nil,
	}
	if route != nil {
		httpRoute(route)
	}

	log.Println(">> Attach:", nameModule, versionModule)
	return config
}

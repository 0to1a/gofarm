package calculateModule

import (
	"embed"
	"framework/app/structure"
	"github.com/labstack/echo/v4"
	"log"
)

///go:embed migration/*.sql
var fsMigrate embed.FS

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

	utils.MigrateTools(fsMigrate)
	initializeCron()

	log.Println(">> Attach:", nameModule, versionModule)
	return config
}

func initializeCron() {
	cron.AddEverySecond(15, func() {
		log.Println("Hello per 15 seconds")
	})
	cron.AddEveryDay("10:00", func() {
		log.Println("Run every 10:00")
	})
}

package exampleModule

import (
	"embed"
	"framework/app/structure"
	"github.com/labstack/echo/v4"
	"log"
)

//go:embed migration/*.sql
var fsMigrate embed.FS

const (
	nameModule    = "Example CRUD Module"
	versionModule = 100
)

func InitializeModule(route *echo.Echo, authMiddleware echo.MiddlewareFunc) structure.ModularStruct {
	config := structure.ModularStruct{
		Name:      nameModule,
		Version:   versionModule,
		Depending: nil,
	}
	config.Depending = append(config.Depending, structure.ModularStruct{
		Name:       "Calculate Module",
		MinVersion: 100,
		MaxVersion: 0,
	})
	if route != nil {
		httpRoute(route)
	}

	utils.MigrateTools(fsMigrate)
	initializeCron()

	log.Println(">> Attach:", nameModule, versionModule)
	return config
}

func initializeCron() {
	cron.AddEverySecond(60, service.CronHelloWorld)
}

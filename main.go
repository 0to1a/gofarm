package main

import (
	"framework/app"
	"framework/app/migration"
	"framework/app/structure"
	"framework/framework/utils"
	"framework/framework/webserver"
	"log"
)

const (
	nameService = "ProjectName"
)

func main() {
	log.Println(nameService)
	utils.ReloadSystem()

	if structure.SystemConf.UseMigration && structure.SystemConf.Database != "" {
		migration.SeedDatabase()
	}

	webserver.CreateService(structure.SystemConf.ServicePort, app.ConfigRoute())

	select {}
}

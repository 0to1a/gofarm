package main

import (
	"framework/app"
	"framework/app/migration"
	"framework/app/structure"
	"framework/framework/database"
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
	if structure.SystemConf.ServiceMonitor {
		utils.UseMonitor()
	}
	if structure.SystemConf.ServiceCronJob {
		app.CronJobMaker()
	}
	if structure.SystemConf.ServiceRedis {
		database.RedisConnect()
	}

	webserver.CreateService(structure.SystemConf.ServicePort, app.ConfigRoute())

	select {}
}

package main

import (
	"framework/app"
	"framework/app/structure"
	"framework/framework"
	"log"
)

var utils framework.Utils

const (
	goFarmVersion = "1.1.0-alpha"
	nameService   = "ProjectName"
)

func main() {
	log.Println(nameService)

	utils.ReloadSystem()
	Redis := framework.RedisDatabase{
		Prefix:   structure.SystemConf.RedisPrefix,
		Host:     structure.SystemConf.RedisHost,
		Password: structure.SystemConf.RedisPassword,
		Database: structure.SystemConf.RedisDatabase,
	}
	WebService := framework.WebServer{}

	// TODO: create migration module
	//if structure.SystemConf.UseMigration && structure.SystemConf.Database != "" {
	//	migration.SeedDatabase()
	//}
	// TODO: log system
	//if structure.SystemConf.ServiceMonitor {
	//	utils.UseMonitor()
	//}
	if structure.SystemConf.ServiceCronJob {
		app.CronJobMaker()
	}
	if structure.SystemConf.ServiceRedis {
		Redis.Connect()
	}

	WebService.CreateService(structure.SystemConf.ServicePort, app.ConfigRoute())

	select {}
}

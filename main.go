package main

import (
	"framework/app"
	"framework/app/structure"
	"framework/framework"
	"log"
)

var utils framework.Utils

const (
	goFarmVersion = "1.1.1-alpha"
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

	// TODO: create migration module
	//if structure.SystemConf.UseMigration && structure.SystemConf.DatabaseMysql != "" {
	//	migration.SeedDatabase()
	//}
	if structure.SystemConf.ServiceRedis {
		Redis.Connect()
	}

	WebService := framework.WebServer{}
	WebService.CreateService(structure.SystemConf.ServicePort, app.ConfigRoute())

	select {}
}

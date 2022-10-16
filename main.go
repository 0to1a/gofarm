package main

import (
	"framework/app"
	"framework/app/structure"
	"framework/framework"
	"log"
)

var utils framework.Utils

const (
	goFarmVersion = "1.1.3-alpha"
	nameService   = "ProjectName"
)

func main() {
	log.Println(nameService)

	utils.ReloadSystem()

	if structure.SystemConf.ServiceRedis {
		Redis := framework.RedisDatabase{
			Prefix:   structure.SystemConf.RedisPrefix,
			Host:     structure.SystemConf.RedisHost,
			Password: structure.SystemConf.RedisPassword,
			Database: structure.SystemConf.RedisDatabase,
		}
		Redis.Connect()
	}

	WebService := framework.WebServer{}
	WebService.CreateService(structure.SystemConf.ServicePort, app.ConfigRoute())

	select {}
}

package main

import (
	"framework/app"
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

	webserver.CreateService(structure.SystemConf.ServicePort, app.ConfigRoute())

	select {}
}

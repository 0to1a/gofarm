package framework

import (
	"framework/app/structure"
	"github.com/go-co-op/gocron"
)

type Utils struct{}

var (
	listModule []*structure.ModularStruct
	utils      Utils
	webserver  WebServer
	dbMysql    MysqlDatabase
	cron       CronUtils
	scheduler  *gocron.Scheduler
)

const (
	SeedOK    = 200
	ErrNoRows = -404
	ErrQuery  = -500
)

const (
	okMigration1 = "no change"
	okMigration2 = "migration success"
)

const (
	errorEnv     = "Err Environment #U0000:"
	errorEnv1    = "Err Environment #U0001: config.json not exist"
	errorModule1 = "Module %s incompatible, target version: %d exist version: %d"
	errorModule2 = "Module '%s' incompatible, no depending '%s' included"
)

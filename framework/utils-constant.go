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
	mysql      MysqlDatabase
	scheduler  *gocron.Scheduler
)

const (
	SeedOK        = 200
	ErrSeedNoRows = -1
	ErrMigration  = -2
	ErrNoRows     = -404
	ErrQuery      = -500
)

const (
	okMigration1    = "No Migration Unit: "
	errorMigration1 = "Err Migration #S0001: no indexing function"
	errorMigration2 = "Err Migration #S0002: database not init"
	errorMigration3 = "Err Migration #S0003: database can't create table"
	errorMigration4 = "Err Migration #S0004: problem to insert data"
)

const (
	errorEnv  = "Err Environment #U0000:"
	errorEnv1 = "Err Environment #U0001: config.json not exist"
)

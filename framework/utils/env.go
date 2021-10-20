package utils

import (
	"encoding/json"
	"framework/app/model"
	"framework/app/structure"
	"framework/framework/database"
	"github.com/doug-martin/goqu/v9"
	"io/ioutil"
	"log"
	"os"
)

var (
	JwtSecret string
)

func ReloadSystem() {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		log.Fatalln(errorEnv1)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &structure.SystemConf)
	if err != nil {
		log.Fatalln(errorEnv, err)
	}

	JwtSecret = structure.SystemConf.SecretKey

	if structure.SystemConf.Database == "" {
		model.Database = nil
	} else if structure.SystemConf.Database == "mysql" {
		conf := structure.SystemConf
		model.Database = database.MysqlConnect(conf.DatabaseUsername, conf.DatabasePassword, conf.DatabaseHost, conf.DatabaseName)
		model.Dialect = goqu.Dialect("mysql")
	}
}

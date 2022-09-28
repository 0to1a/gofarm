package framework

import (
	"encoding/json"
	"framework/app/structure"
	"github.com/doug-martin/goqu/v9"
	"io/ioutil"
	"log"
	"os"
)

func (w *Utils) ReloadSystem() {
	// TODO: read ENV machine first

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

	jwtSecret = structure.SystemConf.SecretKey

	if structure.SystemConf.Database == "" {
		Database = nil
	} else if structure.SystemConf.Database == "mysql" {
		conf := structure.SystemConf
		Database = MysqlConnect(conf.DatabaseUsername, conf.DatabasePassword, conf.DatabaseHost, conf.DatabaseName)
		Dialect = goqu.Dialect("mysql")
	}
}

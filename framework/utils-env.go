package framework

import (
	"encoding/json"
	"framework/app/structure"
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
		DatabaseMysql = nil
	} else if structure.SystemConf.Database == "mysql" {
		mysql = MysqlDatabase{
			Username: structure.SystemConf.DatabaseUsername,
			Password: structure.SystemConf.DatabasePassword,
			Host:     structure.SystemConf.DatabaseHost,
			Database: structure.SystemConf.DatabaseName,
		}
	}
}

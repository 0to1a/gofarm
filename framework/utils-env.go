package framework

import (
	"encoding/json"
	"framework/app/structure"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func (w *Utils) reloadSystemByEnv() {
	v := reflect.ValueOf(&structure.SystemConf).Elem()
	for j := 0; j < v.NumField(); j++ {
		jsonName := v.Type().Field(j).Tag.Get("json")
		typeName := v.Field(j).Type().Name()
		input := os.Getenv(strings.ToUpper(jsonName))
		if typeName == "int" {
			tmp, _ := strconv.ParseInt(input, 10, 64)
			v.Field(j).SetInt(tmp)
		} else if typeName == "float" {
			tmp, _ := strconv.ParseFloat(input, 64)
			v.Field(j).SetFloat(tmp)
		} else if typeName == "bool" {
			if strings.ToUpper(input) == "TRUE" || input == "1" {
				v.Field(j).SetBool(true)
			} else {
				v.Field(j).SetBool(false)
			}
		} else {
			v.Field(j).SetString(input)
		}
	}
}

func (w *Utils) reloadDatabase() {
	if structure.SystemConf.Database == "" {
		DatabaseMysql = nil
	} else if structure.SystemConf.Database == "mysql" {
		dbMysql = MysqlDatabase{
			Username: structure.SystemConf.DatabaseUsername,
			Password: structure.SystemConf.DatabasePassword,
			Host:     structure.SystemConf.DatabaseHost,
			Database: structure.SystemConf.DatabaseName,
		}
	}
}

func (w *Utils) ReloadSystem() {
	w.reloadSystemByEnv()

	jsonFile, err := os.Open("config.json")
	if err != nil && structure.SystemConf.ServicePort == 0 {
		log.Panic(errorEnv1)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &structure.SystemConf)
	if err != nil {
		log.Panic(errorEnv, err)
	}

	cron.Setup()
	jwtSecret = structure.SystemConf.SecretKey
	w.reloadDatabase()
}

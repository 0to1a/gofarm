package utils

import (
	"encoding/json"
	"framework/app/structure"
	"io/ioutil"
	"log"
	"os"
)

var (
	JwtToken string
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

	JwtToken = structure.SystemConf.SecretKey
}

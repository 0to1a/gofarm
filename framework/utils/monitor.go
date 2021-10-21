package utils

import (
	"io"
	"log"
	"os"
)

func UseMonitor() {
	_ = os.MkdirAll("log", os.ModePerm)

	file, e := os.OpenFile("log/monitor.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if e != nil {
		log.Fatalln("Failed to open log file")
	}
	multi := io.MultiWriter(file, os.Stdout)
	log.SetOutput(multi)
}

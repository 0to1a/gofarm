package utils

import (
	"io"
	"log"
	"os"
	"strconv"
	"time"
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

func LogPrintln(s string) {
	t := time.Now()
	year, month, day := t.Date()
	hour, min, sec := t.Clock()

	monthStr := strconv.Itoa(int(month))
	if len(monthStr) < 2 {
		monthStr = "0" + monthStr
	}
	dayStr := strconv.Itoa(day)
	if len(dayStr) < 2 {
		dayStr = "0" + dayStr
	}
	hourStr := strconv.Itoa(hour)
	if len(hourStr) < 2 {
		hourStr = "0" + hourStr
	}
	minStr := strconv.Itoa(min)
	if len(minStr) < 2 {
		minStr = "0" + minStr
	}
	secStr := strconv.Itoa(sec)
	if len(secStr) < 2 {
		secStr = "0" + secStr
	}

	filename := "log-" + strconv.Itoa(year) + "-" + monthStr + "-" + dayStr + ".log"
	file, e := os.OpenFile("log/"+filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if e != nil {
		log.Println("LogPrintln: Failed to open log file")
		return
	}
	defer file.Close()

	formatTime := strconv.Itoa(year) + "/" + monthStr + "/" + dayStr + " " + hourStr + ":" + minStr + ":" + secStr + " "
	if _, err := file.WriteString(formatTime + s + "\n"); err != nil {
		log.Println("LogPrintln: " + err.Error())
		return
	}
}

func LogError(s string) {
	t := time.Now()
	year, month, day := t.Date()
	hour, min, sec := t.Clock()

	monthStr := strconv.Itoa(int(month))
	if len(monthStr) < 2 {
		monthStr = "0" + monthStr
	}
	dayStr := strconv.Itoa(day)
	if len(dayStr) < 2 {
		dayStr = "0" + dayStr
	}
	hourStr := strconv.Itoa(hour)
	if len(hourStr) < 2 {
		hourStr = "0" + hourStr
	}
	minStr := strconv.Itoa(min)
	if len(minStr) < 2 {
		minStr = "0" + minStr
	}
	secStr := strconv.Itoa(sec)
	if len(secStr) < 2 {
		secStr = "0" + secStr
	}

	filename := "error-" + strconv.Itoa(year) + "-" + monthStr + "-" + dayStr + ".log"
	file, e := os.OpenFile("log/"+filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if e != nil {
		log.Println("LogPrintln: Failed to open log file")
		return
	}
	defer file.Close()

	formatTime := strconv.Itoa(year) + "/" + monthStr + "/" + dayStr + " " + hourStr + ":" + minStr + ":" + secStr + " "
	if _, err := file.WriteString(formatTime + s + "\n"); err != nil {
		log.Println("LogPrintln: " + err.Error())
		return
	}
}

package webserver

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	logFile    *os.File = nil
	logErrFile *os.File = nil
)

func CreateService(port int, router *echo.Echo) {
	go func() {
		serverPort := ":" + strconv.Itoa(port)

		log.Print("Webserver: Service Running")
		router.HideBanner = true
		router.Logger.Fatal(router.Start(serverPort))
	}()

	select {}
}

func SetLogFile(logPath string, filenameLog string, filenameError string) {
	_ = os.MkdirAll(logPath, os.ModePerm)
	var err error

	logFile, err = os.OpenFile(logPath+"/"+filenameLog, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to create request log file:", err)
	}

	logErrFile, err = os.OpenFile(logPath+"/"+filenameError, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to create request log file:", err)
	}
}

func ResponseAPI(response int, message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"response":  response,
		"message":   message,
		"data":      data,
		"timestamp": time.Now().Unix(),
	}
}

func ResultAPI(c echo.Context, response int, message string, data interface{}) error {
	return c.JSON(response, ResponseAPI(response, message, data))
}

func ResultAPIFromJson(c echo.Context, mapJson map[string]interface{}) error {
	response := 0
	switch v := mapJson["response"].(type) {
	case int:
		response = v
	case float64:
		response = int(v)
	}
	return c.JSON(response, mapJson)
}

func Logger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			req := c.Request()
			timeFormat := "2006-01-02 15:04:05"
			start := time.Now()

			if err = next(c); err != nil {
				c.Error(err)
			}
			stop := time.Now()

			if err == nil {
				format := fmt.Sprintf("%s\t%s\t\t%s\n", time.Now().Format(timeFormat), req.RequestURI, stop.Sub(start).String())
				_, _ = logFile.WriteString(format)
			} else {
				format := fmt.Sprintf("%s\t%s\t%s\t\t%s\n", time.Now().Format(timeFormat), req.RequestURI, stop.Sub(start).String(), err.Error())
				_, _ = logErrFile.WriteString(format)
			}
			return nil
		}
	}
}

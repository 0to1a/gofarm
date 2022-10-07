package framework

import (
	"github.com/bmizerany/assert"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

func TestCreateService(t *testing.T) {
	t.Run("Running service", func(t *testing.T) {
		e := echo.New()
		timeout := time.After(1 * time.Second)
		done := make(chan bool)
		go func() {
			webserver.CreateService(18080, e)
			done <- true
		}()

		assert.Panic(t, nil, func() {
			select {
			case <-timeout:
				t.Fatal("Test didn't finish in time")
			case <-done:
			}
		})
	})
	t.Run("Failure spawn", func(t *testing.T) {
		e := echo.New()
		assert.Panic(t, "listen tcp: address 100000: invalid port", func() {
			webserver.runService(":100000", e)
		})
	})
}

func TestSetupLogFile(t *testing.T) {
	defer func() {
		err := os.RemoveAll("log-test")
		if err != nil {
			t.Fatal(err)
		}
	}()

	t.Run("Setup normally", func(t *testing.T) {
		webserver.SetupLogFile("log-test", "hello", "world")
		if _, err := os.Stat("log-test/hello"); err != nil {
			t.Fatal(err)
		}
		webserver.logFile.Close()
		webserver.logErrFile.Close()
	})
	t.Run("Filename log error", func(t *testing.T) {
		assert.Panic(t, "Failed to create request log file:open log-test/|hello: The filename, directory name, or volume label syntax is incorrect.", func() {
			webserver.SetupLogFile("log-test", "|hello", "world")
		})
	})
	t.Run("Filename error", func(t *testing.T) {
		assert.Panic(t, "Failed to create request log file:open log-test/|world: The filename, directory name, or volume label syntax is incorrect.", func() {
			webserver.SetupLogFile("log-test", "hello", "|world")
		})
		webserver.logFile.Close()
	})
}

func TestResultAPI(t *testing.T) {
	defer func() {
		webserver.logFile.Close()
		webserver.logErrFile.Close()
		err := os.RemoveAll("log-test")
		if err != nil {
			t.Fatal(err)
		}
	}()

	e := echo.New()
	e.Use(webserver.Logger("log-test", "hello", "world"))
	e.GET("/test", func(context echo.Context) error {
		return webserver.ResultAPI(context, 200, "ok", "ok")
	})

	webserver.CreateService(30080, e)

	t.Run("Send Json from scratch", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "http://localhost:30080/test", strings.NewReader(""))
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, http.StatusOK, res.StatusCode)
	})
	t.Run("Send Json invalid", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "http://localhost:30080/test-hello", strings.NewReader(""))
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
}

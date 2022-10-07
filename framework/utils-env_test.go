package framework

import (
	"framework/app/structure"
	"github.com/bmizerany/assert"
	"os"
	"testing"
)

func TestReloadSystem(t *testing.T) {
	os.Setenv("SERVICE_PORT", "8080")
	os.Setenv("USE_MIGRATION", "FALSE")
	os.Setenv("USE_WEBSERVER_LOG", "TRUE")
	defer os.Unsetenv("SERVICE_PORT")
	defer os.Unsetenv("USE_MIGRATION")
	defer os.Unsetenv("USE_WEBSERVER_LOG")

	t.Run("Run with ENV", func(t *testing.T) {
		assert.Panic(t, nil, func() {
			utils.ReloadSystem()
		})
	})
	t.Run("Run with ENV and mysql", func(t *testing.T) {
		os.Setenv("DATABASE", "mysql")
		defer os.Unsetenv("DATABASE")
		assert.Panic(t, nil, func() {
			utils.ReloadSystem()
		})
	})
	t.Run("Run with ENV and no port", func(t *testing.T) {
		os.Setenv("SERVICE_PORT", "0")
		assert.Panic(t, errorEnv1, func() {
			utils.ReloadSystem()
		})
	})
}

func TestReloadSystemByJson(t *testing.T) {
	stat := func(filename string) (os.FileInfo, error) {
		return nil, nil
	}

	t.Run("Using Valid Json", func(t *testing.T) {
		readFile := func(filename string) ([]byte, error) {
			return []byte(`{"service_port": 1111}`), nil
		}
		utils.reloadSystemByJson("config.json", stat, readFile)
		assert.Equal(t, 1111, structure.SystemConf.ServicePort)
	})
	t.Run("Invalid Json", func(t *testing.T) {
		readFile := func(filename string) ([]byte, error) {
			return []byte(`hello world`), nil
		}
		assert.Panic(t, errorEnv+"invalid character 'h' looking for beginning of value", func() {
			utils.reloadSystemByJson("config.json", stat, readFile)
		})
	})
}

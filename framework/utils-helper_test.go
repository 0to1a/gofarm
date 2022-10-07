package framework

import (
	"github.com/bmizerany/assert"
	"testing"
	"time"
)

func TestGenerateStandardAPI(t *testing.T) {
	t.Run("Output Test", func(t *testing.T) {
		output := utils.GenerateStandardAPI(200, "ok", "ok")
		want := map[string]interface{}{
			"response":  200,
			"message":   "ok",
			"data":      "ok",
			"timestamp": time.Now().Unix(),
		}
		assert.Equal(t, output, want)
	})
}

func TestSeedName(t *testing.T) {
	t.Run("Alphanumeric", func(t *testing.T) {
		output := utils.SeedName("hello world 1")
		assert.Equal(t, output, "hello world 1")
	})
	t.Run("Alphanumeric dash", func(t *testing.T) {
		output := utils.SeedName("hello-world")
		assert.Equal(t, output, "hello-world")
	})
	t.Run("Special character", func(t *testing.T) {
		output := utils.SeedName("hello-world!")
		assert.Equal(t, output, "c31fPscsiUf9y-cGCN9VhtQi6vg=")
	})
}

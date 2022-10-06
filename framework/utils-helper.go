package framework

import (
	"crypto/sha1"
	"encoding/base64"
	"regexp"
	"time"
)

func (w *Utils) GenerateStandardAPI(response int, message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"response":  response,
		"message":   message,
		"data":      data,
		"timestamp": time.Now().Unix(),
	}
}

func (w *Utils) SeedName(functionName string) string {
	if regexp.MustCompile(`^[a-zA-Z0-9 \-]*$`).MatchString(functionName) {
		return functionName
	}

	hashes := sha1.New()
	hashes.Write([]byte(functionName))
	seedName := base64.URLEncoding.EncodeToString(hashes.Sum(nil))

	return seedName
}

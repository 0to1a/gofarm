package framework

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var (
	jwtSecret string
)

func (w *WebServer) JWTCheckToken(tokenString string) (*jwt.Token, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, false
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return nil, false
	}
	return token, true
}

func (w *WebServer) JWTCreateToken(username string, timeInMinutes int) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["username"] = username
	atClaims["exp"] = time.Now().Add(time.Minute * time.Duration(timeInMinutes)).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	secret := jwtSecret
	token, err := at.SignedString([]byte(secret))
	if err != nil || len(secret) == 0 {
		return "", err
	}
	return token, nil
}

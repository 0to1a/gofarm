package framework

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestJWTCheckToken(t *testing.T) {
	t.Run("Valid JWT", func(t *testing.T) {
		jwtSecret = "1234"
		out, _ := webserver.JWTCreateToken("test", 1)

		_, isOkay := webserver.JWTCheckToken(out)
		assert.Equal(t, true, isOkay)
	})

	t.Run("Invalid secret", func(t *testing.T) {
		jwtSecret = "1234"
		out, _ := webserver.JWTCreateToken("test", 1)

		jwtSecret = "5678"
		_, isOkay := webserver.JWTCheckToken(out)
		assert.Equal(t, false, isOkay)
	})

	t.Run("Invalid algo", func(t *testing.T) {
		jwt := "eyJhbGciOiJQUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.iOeNU4dAFFeBwNj6qdhdvm-IvDQrTa6R22lQVJVuWJxorJfeQww5Nwsra0PjaOYhAMj9jNMO5YLmud8U7iQ5gJK2zYyepeSuXhfSi8yjFZfRiSkelqSkU19I-Ja8aQBDbqXf2SAWA8mHF8VS3F08rgEaLCyv98fLLH4vSvsJGf6ueZSLKDVXz24rZRXGWtYYk_OYYTVgR1cg0BLCsuCvqZvHleImJKiWmtS0-CymMO4MMjCy_FIl6I56NqLE9C87tUVpo1mT-kbg5cHDD8I7MjCW5Iii5dethB4Vid3mZ6emKjVYgXrtkOQ-JyGMh6fnQxEFN1ft33GX2eRHluK9eg"
		jwtSecret = "5678"
		_, isOkay := webserver.JWTCheckToken(jwt)
		assert.Equal(t, false, isOkay)
	})
}

func TestJWTCreateToken(t *testing.T) {
	t.Run("Valid JWT", func(t *testing.T) {
		jwtSecret = "1234"
		_, err := webserver.JWTCreateToken("test", 1)
		assert.Equal(t, nil, err)
	})

	t.Run("Invalid output JWT", func(t *testing.T) {
		jwtSecret = ""
		out, _ := webserver.JWTCreateToken("test", 1)
		assert.Equal(t, "", out)
	})
}

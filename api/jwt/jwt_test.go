package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"testing"
)

func TestSign(t *testing.T) {
	claims := jwt.MapClaims{
		"uid": 1000,
	}
	sign, err := Sign(claims)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(sign)
}

func TestVerify(t *testing.T) {
	tokenStr := "eyJhbGciOiJIUzI1NiJ9.eyJ1aWQiOjF9.eZMNx6N5ouOYo2wMlnU1qkhZkEI3PtWqN36x9O_5EQ4"
	claims, err := Verify(tokenStr)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(claims)
	t.Log(claims["uid"])
}

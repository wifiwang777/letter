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
	tokenStr := "eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEwMDB9.50JcnincZUSwL7nGtjCS7MzE9kggMDDwDupTem4cDAjZdRGv4UNqCWaXWRSmhvEyP6LJXyp-0IqEfI6s1pVxpw"
	claims, err := Verify(tokenStr)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(claims)
	t.Log(claims["uid"])
}

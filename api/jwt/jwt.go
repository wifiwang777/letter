package jwt

import (
	"github.com/aisuosuo/letter/config"
	"github.com/golang-jwt/jwt/v4"
)

func Sign(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString(config.JwtSignKey)
	if err != nil {
		return "", err
	}
	return signedString, nil
}

func Verify(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr,
		func(token *jwt.Token) (i interface{}, e error) {
			return config.JwtSignKey, nil
		})
	if err != nil {
		return nil, err
	}
	return token.Claims.(jwt.MapClaims), nil
}

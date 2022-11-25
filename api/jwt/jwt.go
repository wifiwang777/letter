package jwt

import (
	"github.com/aisuosuo/letter/config"
	"github.com/golang-jwt/jwt/v4"
)

func Sign(claims jwt.MapClaims) (string, error) {
	alg := jwt.GetSigningMethod("ES256")
	token := jwt.NewWithClaims(alg, claims)
	signedString, err := token.SignedString(config.JwtPrivateKey.PrivateKey)
	if err != nil {
		return "", err
	}
	return signedString, nil
}

func Verify(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr,
		func(t *jwt.Token) (interface{}, error) {
			return config.JwtPublicKey.PublicKey, nil
		})
	if err != nil {
		return nil, err
	}
	return token.Claims.(jwt.MapClaims), nil
}

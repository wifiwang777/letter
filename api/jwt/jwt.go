package jwt

import (
	"fmt"
	"github.com/aisuosuo/letter/config/apollo"
	"github.com/golang-jwt/jwt/v4"
)

func Sign(claims jwt.MapClaims) (string, error) {
	jwtSecretKey, err := apollo.GlobalApolloConfig.Get("jwt.secretKey")
	if err != nil {
		return "", err
	}
	key := []byte(jwtSecretKey.(string))
	alg := jwt.GetSigningMethod("HS256")
	token := jwt.NewWithClaims(alg, claims)
	signedString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return signedString, nil
}

func Verify(tokenStr string) (jwt.MapClaims, error) {
	jwtSecretKey, err := apollo.GlobalApolloConfig.Get("jwt.secretKey")
	if err != nil {
		return nil, err
	}
	key := []byte(jwtSecretKey.(string))
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims.(jwt.MapClaims), nil
}

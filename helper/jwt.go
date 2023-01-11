package helper

import "github.com/golang-jwt/jwt/v4"

func CreateJWT(secret string, claims jwt.MapClaims) (signedToken string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString([]byte(secret))
	if err != nil {
		return
	}

	return
}

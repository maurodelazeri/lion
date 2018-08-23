package token

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//https://echo.labstack.com/cookbook/jwt
//https://github.com/istio/istio/blob/8aa98402bffcdbf0dc18558ae4175529ca4d7dec/security/pkg/platform/gcp.go

// CreateToken create a jwt token
func CreateToken(account, subject string) (string, error) {
	now := time.Now()
	hostname, _ := os.Hostname()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &jwt.StandardClaims{
		IssuedAt: now.Unix(),
		Issuer:   hostname,
		Subject:  subject,
		Id:       account,
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT-TOKEN-SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

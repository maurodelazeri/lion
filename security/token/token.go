package token

import (
	"errors"
	"os"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CreateToken create a jwt token
// MODE:
// DEMO               = 0;
// LIVE               = 1;
// BACKTESTING        = 2;
func CreateToken(account, mode string) (string, error) {
	now := time.Now()
	expTime, _ := strconv.Atoi(os.Getenv("EXPIRATION_TIME_JWT_TOKEN"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &jwt.StandardClaims{
		IssuedAt:  now.Unix(),
		ExpiresAt: time.Now().Add(time.Minute * time.Duration(expTime)).Unix(),
		//Issuer:    mode,
		//Audience:  strategy,
		Subject: mode,
		Id:      account,
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SIGNED_JWT_KEY")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ValidateToken token
func ValidateToken(token string) (*jwt.StandardClaims, error) {
	refreshToken, reason := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGNED_JWT_KEY")), nil
	})
	if reason != nil {
		return nil, reason
	}
	if claims, ok := refreshToken.Claims.(*jwt.StandardClaims); ok && refreshToken.Valid {
		return claims, nil
	}
	return nil, errors.New("Invalid token")
}

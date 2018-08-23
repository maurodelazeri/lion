package token

import (
	"errors"
	"os"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CreateToken create a jwt token
func CreateToken(account, userid string) (string, error) {
	now := time.Now()
	hostname, _ := os.Hostname()
	expTime, _ := strconv.Atoi(os.Getenv("EXPIRATION_TIME_JWT_TOKEN"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &jwt.StandardClaims{
		IssuedAt:  now.Unix(),
		ExpiresAt: time.Now().Add(time.Minute * time.Duration(expTime)).Unix(),
		Issuer:    hostname,
		Subject:   userid,
		Id:        account,
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
		// return new access token and refresh token
		// id, _ := strconv.Atoi(claims.Subject)
		// context.JSON(http.StatusOK, gin.H{
		// 	"refresh_access": libs.GenerateRefreshToken(claims.Email, id),
		// 	"access_token":   libs.GenerateAccessToken(claims.Email, id),
		// })
		return claims, nil
	}
	return nil, errors.New("Invalid token")
}

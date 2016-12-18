package jwtauth

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	UserID    uint   `json:"user_id"`
	UserEmail string `json:"user_email"`
	jwt.StandardClaims
}

func GenerateToken(id uint, email string) (string, error) {
	now := time.Now()
	exp := now.Add(6 * time.Hour)

	claims := Claims{
		id,
		email,
		jwt.StandardClaims{
			ExpiresAt: exp.Unix(),
			Issuer:    "",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtSecret)

	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}

// func VerifyToken(tokenStr string) (*

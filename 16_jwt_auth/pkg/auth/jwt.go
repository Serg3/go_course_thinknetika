package jwtauth

import (
	"errors"
	users "go_course_thinknetika/16_jwt_auth/pkg/db"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type tokenClaims struct {
	UserID int
	Admin  bool
	jwt.StandardClaims
}

// NewToken generates and returns a new token as []byte.
func NewToken(usr *users.User) ([]byte, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserID: usr.ID(),
		Admin:  usr.Admin(),
	})

	key, _ := jwtKey(nil)
	tokenStr, err := token.SignedString(key)
	if err != nil {
		return nil, err
	}
	return []byte(tokenStr), nil
}

// VerifyToken checks if the token is valid.
func VerifyToken(tokenStr string) (*tokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &tokenClaims{}, jwtKey)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("token is not valid")
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return nil, errors.New("wrong token claims")
	}
	return claims, nil
}

func jwtKey(token *jwt.Token) (interface{}, error) {
	return []byte("ENV_SECRET_KEY"), nil
}

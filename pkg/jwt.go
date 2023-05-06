package pkg

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var ( 
	JWT_KEY = []byte("dashdfgahfgsdahjfgasdhjk")
	ErrParseClaims = errors.New("cannot parse jwt")
	ErrJWTNotValid = errors.New("jwt not valid")
)


type Claims struct {
	User string
	jwt.RegisteredClaims
}

type JWT struct{}

func (JWT) Generate(expr time.Duration, username string) (string, error) {
	claims := &Claims{
		User: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expr)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWT_KEY)

	return tokenString, err
}

func (JWT) Validate(token string) error {
	claims := &Claims{}

	tokenParsed, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (any, error) {
		return JWT_KEY, nil
	})

	if err != nil {
		return ErrParseClaims
	}

	if !tokenParsed.Valid {
		return ErrJWTNotValid
	}

	return nil
}

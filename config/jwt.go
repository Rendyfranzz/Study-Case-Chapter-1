package config

import "github.com/golang-jwt/jwt/v5"

var JWT_KEY = []byte("dashdfgahfgsdahjfgasdhjk")

type Claims struct {
	User string `json:"name" bson:"name"`
	jwt.RegisteredClaims
}

package models

import "github.com/dgrijalva/jwt-go"

type UserClaims struct {
	jwt.StandardClaims
	Id       string `json:"id"`
	Username string `json:"username"`
}

package handler

import "github.com/golang-jwt/jwt/v5"

type UserClaims struct {
	Identity string `json:"identity"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

type SignInBody struct {
	Identity string `json:"identity"`
	Password string `json:"password"`
}

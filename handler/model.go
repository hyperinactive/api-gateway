package handler

import (
	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	UserId string `json:"userId"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type SignInBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignUpBody struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

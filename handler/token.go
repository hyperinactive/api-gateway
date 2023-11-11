package handler

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// creates user claims and does user validation
func createUserClaim(identity string, role string) (*UserClaims, error) {
	// Validate role against allowed roles
	validRoles := map[string]bool{
		"admin": true,
		"user":  true,
	}

	if !validRoles[role] {
		return nil, errors.New("invalid role")
	}

	return &UserClaims{
		Identity: identity,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}, nil
}

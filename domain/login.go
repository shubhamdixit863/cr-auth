package domain

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func (l User) ClaimsForAccessToken() AccessTokenClaims {
	return l.claims()
}

func (l User) claims() AccessTokenClaims {
	return AccessTokenClaims{
		Username: l.Username,
		Role:     l.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ACCESS_TOKEN_DURATION).Unix(),
		},
	}
}

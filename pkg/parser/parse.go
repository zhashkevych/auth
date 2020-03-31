package parser

import (
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/zhashkevych/auth/pkg/auth"
)

func ParseToken(accessToken string, signingKey []byte) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &auth.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*auth.Claims); ok && token.Valid {
		return claims.Username, nil
	}

	return "", auth.ErrInvalidAccessToken
}

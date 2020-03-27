package auth

import (
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/zhashkevych/auth/pkg/models"
)

type Claims struct {
	jwt.StandardClaims
	User *models.User `json:"user"`
}

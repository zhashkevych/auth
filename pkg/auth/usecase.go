package auth

import (
	"context"
	"github.com/zhashkevych/auth/pkg/models"
)

type UseCase interface {
	SignUp(ctx context.Context, user *models.User) error
	SignIn(ctx context.Context, user *models.User) (string, error)
}

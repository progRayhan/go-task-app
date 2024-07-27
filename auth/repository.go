package auth

import (
	"context"

	"github.com/mchayapol/go-todo-app/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, username, password string) (*models.User, error)
}

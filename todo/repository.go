package todo

import (
	"context"

	"github.com/mchayapol/go-todo-app/models"
)

type Repository interface {
	CreateTodo(ctx context.Context, user *models.User, bm *models.Todo) error
	GetTodos(ctx context.Context, user *models.User) ([]*models.Todo, error)
	DeleteTodo(ctx context.Context, user *models.User, id string) error
}

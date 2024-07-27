package todo

import (
	"context"

	"github.com/mchayapol/go-todo-app/models"
)

type UseCase interface {
	CreateTodo(ctx context.Context, user *models.User, completed bool, title string) error
	GetTodos(ctx context.Context, user *models.User) ([]*models.Todo, error)
	DeleteTodo(ctx context.Context, user *models.User, id string) error
}

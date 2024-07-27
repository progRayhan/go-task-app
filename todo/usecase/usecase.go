package usecase

import (
	"context"

	"github.com/mchayapol/go-todo-app/models"
	"github.com/mchayapol/go-todo-app/todo"
)

type TodoUseCase struct {
	todoRepo todo.Repository
}

func NewTodoUseCase(todoRepo todo.Repository) *TodoUseCase {
	return &TodoUseCase{
		todoRepo: todoRepo,
	}
}

func (b TodoUseCase) CreateTodo(ctx context.Context, user *models.User, completed bool, title string) error {
	todoItem := &models.Todo{
		Completed: completed,
		Title:     title,
	}

	return b.todoRepo.CreateTodo(ctx, user, todoItem)
}

func (b TodoUseCase) GetTodos(ctx context.Context, user *models.User) ([]*models.Todo, error) {
	return b.todoRepo.GetTodos(ctx, user)
}

func (b TodoUseCase) DeleteTodo(ctx context.Context, user *models.User, id string) error {
	return b.todoRepo.DeleteTodo(ctx, user, id)
}

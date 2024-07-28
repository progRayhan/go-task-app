package task

import (
	"context"

	"github.com/mchayapol/go-task-app/models"
)

type UseCase interface {
	CreateTask(ctx context.Context, user *models.User, completed bool, title string) error
	GetTasks(ctx context.Context, user *models.User) ([]*models.Task, error)
	DeleteTask(ctx context.Context, user *models.User, id string) error
}

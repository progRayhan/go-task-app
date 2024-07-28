package task

import (
	"context"

	"github.com/mchayapol/go-task-app/models"
)

type Repository interface {
	CreateTask(ctx context.Context, user *models.User, task *models.Task) error
	GetTasks(ctx context.Context, user *models.User) ([]*models.Task, error)
	DeleteTask(ctx context.Context, user *models.User, id string) error
}

package usecase

import (
	"context"

	"github.com/mchayapol/go-task-app/models"
	"github.com/mchayapol/go-task-app/task"
)

type TaskUseCase struct {
	taskRepo task.Repository
}

func NewTaskUseCase(taskRepo task.Repository) *TaskUseCase {
	return &TaskUseCase{
		taskRepo: taskRepo,
	}
}

func (b TaskUseCase) CreateTask(ctx context.Context, user *models.User, completed bool, title string) error {
	taskItem := &models.Task{
		Completed: completed,
		Title:     title,
	}

	return b.taskRepo.CreateTask(ctx, user, taskItem)
}

func (b TaskUseCase) GetTasks(ctx context.Context, user *models.User) ([]*models.Task, error) {
	return b.taskRepo.GetTasks(ctx, user)
}

func (b TaskUseCase) DeleteTask(ctx context.Context, user *models.User, id string) error {
	return b.taskRepo.DeleteTask(ctx, user, id)
}

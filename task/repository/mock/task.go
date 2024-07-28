package mock

import (
	"context"

	"github.com/mchayapol/go-task-app/models"
	"github.com/stretchr/testify/mock"
)

type TaskStorageMock struct {
	mock.Mock
}

func (s *TaskStorageMock) CreateTask(ctx context.Context, user *models.User, bm *models.Task) error {
	args := s.Called(user, bm)

	return args.Error(0)
}

func (s *TaskStorageMock) GetTasks(ctx context.Context, user *models.User) ([]*models.Task, error) {
	args := s.Called(user)

	return args.Get(0).([]*models.Task), args.Error(1)
}

func (s *TaskStorageMock) DeleteTask(ctx context.Context, user *models.User, id string) error {
	args := s.Called(user, id)

	return args.Error(0)
}

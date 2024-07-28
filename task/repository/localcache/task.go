package localcache

import (
	"context"
	"sync"

	"github.com/mchayapol/go-task-app/models"

	"github.com/mchayapol/go-task-app/task"
)

type TaskLocalStorage struct {
	tasks map[string]*models.Task
	mutex *sync.Mutex
}

func NewTaskLocalStorage() *TaskLocalStorage {
	return &TaskLocalStorage{
		tasks: make(map[string]*models.Task),
		mutex: new(sync.Mutex),
	}
}

func (s *TaskLocalStorage) CreateTask(ctx context.Context, user *models.User, bm *models.Task) error {
	bm.UserID = user.ID

	s.mutex.Lock()
	s.tasks[bm.ID] = bm
	s.mutex.Unlock()

	return nil
}

func (s *TaskLocalStorage) GetTasks(ctx context.Context, user *models.User) ([]*models.Task, error) {
	tasks := make([]*models.Task, 0)

	s.mutex.Lock()
	for _, taskItem := range s.tasks {
		if taskItem.UserID == user.ID {
			tasks = append(tasks, taskItem)
		}
	}
	s.mutex.Unlock()

	return tasks, nil
}

func (s *TaskLocalStorage) DeleteTask(ctx context.Context, user *models.User, id string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	taskItem, ex := s.tasks[id]
	if ex && taskItem.UserID == user.ID {
		delete(s.tasks, id)
		return nil
	}

	return task.ErrTaskNotFound
}

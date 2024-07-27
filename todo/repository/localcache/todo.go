package localcache

import (
	"context"
	"sync"

	"github.com/mchayapol/go-todo-app/models"

	"github.com/mchayapol/go-todo-app/todo"
)

type TodoLocalStorage struct {
	todos map[string]*models.Todo
	mutex *sync.Mutex
}

func NewTodoLocalStorage() *TodoLocalStorage {
	return &TodoLocalStorage{
		todos: make(map[string]*models.Todo),
		mutex: new(sync.Mutex),
	}
}

func (s *TodoLocalStorage) CreateTodo(ctx context.Context, user *models.User, bm *models.Todo) error {
	bm.UserID = user.ID

	s.mutex.Lock()
	s.todos[bm.ID] = bm
	s.mutex.Unlock()

	return nil
}

func (s *TodoLocalStorage) GetTodos(ctx context.Context, user *models.User) ([]*models.Todo, error) {
	todos := make([]*models.Todo, 0)

	s.mutex.Lock()
	for _, todoItem := range s.todos {
		if todoItem.UserID == user.ID {
			todos = append(todos, todoItem)
		}
	}
	s.mutex.Unlock()

	return todos, nil
}

func (s *TodoLocalStorage) DeleteTodo(ctx context.Context, user *models.User, id string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	todoItem, ex := s.todos[id]
	if ex && todoItem.UserID == user.ID {
		delete(s.todos, id)
		return nil
	}

	return todo.ErrTodoNotFound
}

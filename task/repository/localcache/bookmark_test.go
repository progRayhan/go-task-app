package localcache

import (
	"context"
	"fmt"
	"testing"

	"github.com/mchayapol/go-task-app/models"
	"github.com/mchayapol/go-task-app/task"
	"github.com/stretchr/testify/assert"
)

func TestGetTasks(t *testing.T) {
	id := "id"
	user := &models.User{ID: id}

	s := NewTaskLocalStorage()

	for i := 0; i < 10; i++ {
		bm := &models.Task{
			ID:     fmt.Sprintf("id%d", i),
			UserID: user.ID,
		}

		err := s.CreateTask(context.Background(), user, bm)
		assert.NoError(t, err)
	}

	returnedTasks, err := s.GetTasks(context.Background(), user)
	assert.NoError(t, err)

	assert.Equal(t, 10, len(returnedTasks))
}

func TestDeleteTask(t *testing.T) {
	id1 := "id1"
	id2 := "id2"

	user1 := &models.User{ID: id1}
	user2 := &models.User{ID: id2}

	taskID := "bmID"
	tk := &models.Task{ID: taskID, UserID: user1.ID}

	s := NewTaskLocalStorage()

	err := s.CreateTask(context.Background(), user1, tk)
	assert.NoError(t, err)

	err = s.DeleteTask(context.Background(), user1, taskID)
	assert.NoError(t, err)

	err = s.CreateTask(context.Background(), user1, tk)
	assert.NoError(t, err)

	err = s.DeleteTask(context.Background(), user2, taskID)
	assert.Error(t, err)
	assert.Equal(t, err, task.ErrTaskNotFound)
}

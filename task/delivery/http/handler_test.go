package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mchayapol/go-task-app/auth"
	"github.com/mchayapol/go-task-app/models"
	"github.com/mchayapol/go-task-app/task/usecase"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	testUser := &models.User{
		Username: "testuser",
		Password: "testpass",
	}

	r := gin.Default()
	group := r.Group("/api", func(c *gin.Context) {
		c.Set(auth.CtxUserKey, testUser)
	})

	uc := new(usecase.TaskUseCaseMock)

	RegisterHTTPEndpoints(group, uc)

	inp := &createInput{
		Completed: false,
		Title:     "testtitle",
	}

	body, err := json.Marshal(inp)
	assert.NoError(t, err)

	uc.On("CreateTask", testUser, inp.Completed, inp.Title).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/tasks", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGet(t *testing.T) {
	testUser := &models.User{
		Username: "testuser",
		Password: "testpass",
	}

	r := gin.Default()
	group := r.Group("/api", func(c *gin.Context) {
		c.Set(auth.CtxUserKey, testUser)
	})

	uc := new(usecase.TaskUseCaseMock)

	RegisterHTTPEndpoints(group, uc)

	taskItems := make([]*models.Task, 5)
	for i := 0; i < 5; i++ {
		taskItems[i] = &models.Task{
			ID:        "id",
			Completed: false,
			Title:     "title",
		}
	}

	uc.On("GetTasks", testUser).Return(taskItems, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/tasks", nil)
	r.ServeHTTP(w, req)

	expectedOut := &getResponse{Tasks: toTasks(taskItems)}

	expectedOutBody, err := json.Marshal(expectedOut)
	assert.NoError(t, err)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(expectedOutBody), w.Body.String())
}

func TestDelete(t *testing.T) {
	testUser := &models.User{
		Username: "testuser",
		Password: "testpass",
	}

	r := gin.Default()
	group := r.Group("/api", func(c *gin.Context) {
		c.Set(auth.CtxUserKey, testUser)
	})

	uc := new(usecase.TaskUseCaseMock)

	RegisterHTTPEndpoints(group, uc)

	inp := &deleteInput{
		ID: "id",
	}

	body, err := json.Marshal(inp)
	assert.NoError(t, err)

	uc.On("DeleteTask", testUser, inp.ID).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/tasks", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

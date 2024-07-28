package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mchayapol/go-task-app/auth"
	"github.com/mchayapol/go-task-app/models"
	"github.com/mchayapol/go-task-app/task"
)

// Locally defined because we don't want to expose UserID
type Task struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type Handler struct {
	useCase task.UseCase
}

func NewHandler(useCase task.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

type createInput struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func (h *Handler) Create(c *gin.Context) {
	inp := new(createInput)
	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := c.MustGet(auth.CtxUserKey).(*models.User)

	if err := h.useCase.CreateTask(c.Request.Context(), user, inp.Completed, inp.Title); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

type getResponse struct {
	Tasks []*Task `json:"tasks"`
}

func (h *Handler) Get(c *gin.Context) {
	user := c.MustGet(auth.CtxUserKey).(*models.User)

	taskItems, err := h.useCase.GetTasks(c.Request.Context(), user)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, &getResponse{
		Tasks: toTasks(taskItems),
	})
}

type deleteInput struct {
	ID string `json:"id"`
}

func (h *Handler) Delete(c *gin.Context) {
	inp := new(deleteInput)
	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := c.MustGet(auth.CtxUserKey).(*models.User)

	if err := h.useCase.DeleteTask(c.Request.Context(), user, inp.ID); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func toTasks(bs []*models.Task) []*Task {
	out := make([]*Task, len(bs))

	for i, b := range bs {
		out[i] = toTask(b)
	}

	return out
}

func toTask(b *models.Task) *Task {
	return &Task{
		ID:        b.ID,
		Title:     b.Title,
		Completed: b.Completed,
	}
}

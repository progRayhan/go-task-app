package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mchayapol/go-todo-app/auth"
	"github.com/mchayapol/go-todo-app/models"
	"github.com/mchayapol/go-todo-app/todo"
)

// Locally defined because we don't want to expose UserID
type Todo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type Handler struct {
	useCase todo.UseCase
}

func NewHandler(useCase todo.UseCase) *Handler {
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

	if err := h.useCase.CreateTodo(c.Request.Context(), user, inp.Completed, inp.Title); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

type getResponse struct {
	Todos []*Todo `json:"todos"`
}

func (h *Handler) Get(c *gin.Context) {
	user := c.MustGet(auth.CtxUserKey).(*models.User)

	todoItems, err := h.useCase.GetTodos(c.Request.Context(), user)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, &getResponse{
		Todos: toTodos(todoItems),
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

	if err := h.useCase.DeleteTodo(c.Request.Context(), user, inp.ID); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func toTodos(bs []*models.Todo) []*Todo {
	out := make([]*Todo, len(bs))

	for i, b := range bs {
		out[i] = toTodo(b)
	}

	return out
}

func toTodo(b *models.Todo) *Todo {
	return &Todo{
		ID:        b.ID,
		Title:     b.Title,
		Completed: b.Completed,
	}
}

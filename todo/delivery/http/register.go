package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mchayapol/go-todo-app/todo"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, uc todo.UseCase) {
	h := NewHandler(uc)

	todos := router.Group("/todos")
	{
		todos.POST("", h.Create)
		todos.GET("", h.Get)
		todos.DELETE("", h.Delete)
	}
}

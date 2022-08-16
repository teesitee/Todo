package router

import (
	"todo-with-gig/api"

	"github.com/gin-gonic/gin"
)

func NewRouter(h api.Handler) *gin.Engine {
	router := gin.Default()
	router.GET("/todos", h.GetTodoList)
	router.GET("/todos/personName/:personName", h.GetTodo)
	router.POST("/todo", h.CreateTodo)
	router.DELETE("/todo", h.DeleteTodo)
	router.PATCH("/todo/:personName/:text", h.PatchTodo)
	return router
}

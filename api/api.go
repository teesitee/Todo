package api

import (
	"net/http"
	"todo-with-gig/repository"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Repo repository.MongoRepository
}

func (h Handler) GetTodoList(c *gin.Context) {
	todos, err := h.Repo.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"todos": todos})

}

type GetRequest struct {
	Text       string `json:"text"`
	PersonName string `json:"personName"`
}

func (h Handler) GetTodo(c *gin.Context) {

	todos, err := h.Repo.GetTodos(c.Param("personName"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"todo": todos})
}

func (h Handler) DeleteTodo(c *gin.Context) {

	err := h.Repo.DelTodos(c.Query("personName"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Result": "ok"})

	todos, err := h.Repo.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"todos": todos})
}

func (h Handler) PatchTodo(c *gin.Context) {

	err := h.Repo.PatTodos(c.Param("personName"), c.Param("text"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Result": "ok"})
}

type CreateRequest struct {
	Text       string `json:"text"`
	PersonName string `json:"personName"`
}

func (h Handler) CreateTodo(c *gin.Context) {
	var req CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.Repo.InsertTodo(req.Text, req.PersonName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	todos, err := h.Repo.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"todos": todos})

}

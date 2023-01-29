package controller

import (
	"test-api/entity"
	"test-api/service"

	"github.com/gin-gonic/gin"
)

type TodoController interface {
	GetAll() []entity.Todo
	Save(ctx *gin.Context) entity.Todo
	Delete(id int32) entity.Todo
	Update(entity.Todo) entity.Todo
}

type contoller struct {
	service service.TodoService
}

func New(service service.TodoService) TodoController {
	return &contoller{
		service: service,
	}
}
func (c *contoller) Update(todo entity.Todo) entity.Todo {
	return c.service.Update(todo)
}
func (c *contoller) Delete(id int32) entity.Todo {
	return c.service.Delete(id)
}
func (c *contoller) GetAll() []entity.Todo {
	return c.service.GetAll()
}

func (c *contoller) Save(ctx *gin.Context) entity.Todo {
	var todo entity.Todo
	ctx.BindJSON(&todo)
	c.service.Save(todo)
	return todo
}

package main

import (
	"net/http"
	"test-api/controller"
	"test-api/entity"
	"test-api/service"

	"github.com/gin-gonic/gin"
)

var (
	todoService    service.TodoService       = service.New()
	todoController controller.TodoController = controller.New(todoService)
)

func main() {
	r := gin.Default()
	r.GET("/todos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, todoController.GetAll())
	})

	r.POST("/create", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, todoController.Save(ctx))
	})
	r.DELETE("/todos", func(ctx *gin.Context) {
		type id struct {
			Id int32 `json:"id"`
		}
		myid := id{}
		ctx.BindJSON(&myid)
		ctx.JSON(http.StatusOK, todoController.Delete(myid.Id))
	})
	r.PUT("/todos", func(ctx *gin.Context) {
		todo := entity.Todo{}
		ctx.BindJSON(&todo)
		ctx.JSON(http.StatusOK, todoController.Update(todo))
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}

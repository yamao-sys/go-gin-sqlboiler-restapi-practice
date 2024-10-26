package routers

import (
	"app/controllers"

	"github.com/gin-gonic/gin"
)

type TodoRouter interface {
	SetRouting(r *gin.Engine)
}

type todoRouter struct {
	todoController controllers.TodoController
}

func NewTodoRouter(todoController controllers.TodoController) TodoRouter {
	return &todoRouter{todoController}
}

func (tr *todoRouter) SetRouting(r *gin.Engine) {
	r.POST("/todos/", tr.todoController.Create)
	r.GET("/todos/", tr.todoController.Index)
	r.GET("/todos/:id", tr.todoController.Show)
	r.PUT("/todos/:id", tr.todoController.Update)
	r.DELETE("/todos/:id", tr.todoController.Delete)
}

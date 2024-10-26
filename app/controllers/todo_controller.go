package controllers

import (
	"app/dto"
	"app/services"
	"app/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoController interface {
	Create(ctx *gin.Context)
	Index(ctx *gin.Context)
	Show(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type todoController struct {
	todoService services.TodoService
	authService services.AuthService
}

func NewTodoController(todoService services.TodoService, authService services.AuthService) TodoController {
	return &todoController{todoService, authService}
}

func (todoController *todoController) Create(ctx *gin.Context) {
	user, err := todoController.authService.GetAuthUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized error"})
		return
	}

	// NOTE: リクエストデータを構造体に変換
	requestParams := dto.CreateTodoRequest{}
	if err := ctx.ShouldBind(&requestParams); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	result := todoController.todoService.CreateTodo(ctx, requestParams, user.ID)

	if result.Error == nil {
		ctx.JSON(http.StatusOK, gin.H{"todo": result.Todo})
		return
	}

	switch result.ErrorType {
	case "internalServerError":
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
	case "validationError":
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.CoordinateValidationErrors(result.Error)})
	}
}

func (todoController *todoController) Index(ctx *gin.Context) {
	user, err := todoController.authService.GetAuthUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized error"})
		return
	}

	result := todoController.todoService.FetchTodosList(ctx, user.ID)

	if result.Error == nil {
		ctx.JSON(http.StatusOK, gin.H{"todos": result.Todos})
		return
	}

	switch result.ErrorType {
	case "notFound":
		ctx.JSON(http.StatusNotFound, gin.H{"error": result.Error})
	case "internalServerError":
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
	}
}

func (todoController *todoController) Show(ctx *gin.Context) {
	user, err := todoController.authService.GetAuthUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized error"})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	result := todoController.todoService.FetchTodo(ctx, id, user.ID)

	if result.Error == nil {
		ctx.JSON(http.StatusOK, gin.H{"todo": result.Todo})
		return
	}

	switch result.ErrorType {
	case "notFound":
		ctx.JSON(http.StatusNotFound, gin.H{"error": result.Error})
	case "internalServerError":
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
	}
}

func (todoController *todoController) Update(ctx *gin.Context) {
	user, err := todoController.authService.GetAuthUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized error"})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	// NOTE: リクエストデータを構造体に変換
	requestParams := dto.UpdateTodoRequest{}
	if err := ctx.ShouldBind(&requestParams); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	result := todoController.todoService.UpdateTodo(ctx, id, requestParams, user.ID)

	if result.Error == nil {
		ctx.JSON(http.StatusOK, gin.H{"todo": result.Todo})
		return
	}

	switch result.ErrorType {
	case "validationError":
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.CoordinateValidationErrors(result.Error)})
	case "notFound":
		ctx.JSON(http.StatusNotFound, gin.H{"error": result.Error})
	case "internalServerError":
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
	}
}

func (todoController *todoController) Delete(ctx *gin.Context) {
	user, err := todoController.authService.GetAuthUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized error"})
		return
	}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	result := todoController.todoService.DeleteTodo(ctx, id, user.ID)

	if result.Error == nil {
		ctx.JSON(http.StatusOK, gin.H{"result": "delete todo(ID: " + ctx.Param("id") + ") successfully"})
		return
	}

	switch result.ErrorType {
	case "notFound":
		ctx.JSON(http.StatusNotFound, gin.H{"error": result.Error})
	case "internalServerError":
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
	}
}

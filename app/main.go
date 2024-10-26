package main

import (
	"app/controllers"
	"app/db"
	"app/routers"
	"app/services"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbCon := db.Init()

	// service
	authService := services.NewAuthService(dbCon)
	todoService := services.NewTodoService(dbCon)

	// controller
	authController := controllers.NewAuthController(authService)
	todoController := controllers.NewTodoController(todoService, authService)
	authRouter := routers.NewAuthRouter(authController)
	todoRouter := routers.NewTodoRouter(todoController)

	// router
	r := gin.Default()
	authRouter.SetRouting(r)
	todoRouter.SetRouting(r)
	r.Run(":" + os.Getenv("SERVER_PORT"))
}

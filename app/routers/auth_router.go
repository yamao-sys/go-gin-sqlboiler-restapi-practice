package routers

import (
	"app/controllers"

	"github.com/gin-gonic/gin"
)

type AuthRouter interface {
	SetRouting(r *gin.Engine)
}

type authRouter struct {
	authController controllers.AuthController
}

func NewAuthRouter(authController controllers.AuthController) AuthRouter {
	return &authRouter{authController}
}

func (ar *authRouter) SetRouting(r *gin.Engine) {
	r.POST("/auth/sign_up", ar.authController.SignUp)
	r.POST("/auth/sign_in", ar.authController.SignIn)
}

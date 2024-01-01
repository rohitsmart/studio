package route

import (
	"github.com/gin-gonic/gin"
	"github.com/rohitsmart/studio/controller"
)

func SetupRouter(authController *controller.AuthController) *gin.Engine {
	r := gin.Default()

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/signup", authController.SignUp)
		authGroup.POST("/login", authController.Login)
	}

	return r
}

package routes

import (
	"github.com/MatheusMangueira/go/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(
	r *gin.RouterGroup,
	userController controller.UserControllerInterface,
) {

	r.GET("/v1/userById/:userId", userController.FindUserById)
	r.GET("/v1/userByEmail/:userEmail", userController.FindUserByEmail)
	r.GET("/v1/user", userController.FindUser)
	r.POST("/v1/user", userController.CreateUser)
	r.PUT("/v1/user/:userId", userController.UpdateUser)
	r.DELETE("/v1/user/:userId", userController.DeleteUser)
}

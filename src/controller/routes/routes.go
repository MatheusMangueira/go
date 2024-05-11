package routes

import (
	"github.com/MatheusMangueira/go/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/v1/userById/:userId", controller.FindUserById)
	r.GET("/v1/userByEmail/:userEmail", controller.FindUserByEmail)
	r.GET("/v1/user", controller.FindUser)
	r.POST("/v1/user", controller.CreateUser)
	r.PUT("/v1/user/:userId", controller.UpdateUser)
	r.DELETE("/v1/user/:userId", controller.DeleteUser)
}

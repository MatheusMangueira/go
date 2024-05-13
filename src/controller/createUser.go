package controller

import (
	"net/http"

	"github.com/MatheusMangueira/go/src/configuration/logger"
	"github.com/MatheusMangueira/go/src/configuration/validation"
	"github.com/MatheusMangueira/go/src/controller/model/request"
	"github.com/MatheusMangueira/go/src/model"
	"github.com/MatheusMangueira/go/src/model/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {

	logger.Info("Init Create User Controller",
		zap.String("journey", "CreateUser"))

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {

		logger.Error("Error trying to validate user info ", err,
			zap.String("journey", "CreateUser"))

		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	service := service.NewUserDomainService()

	if err := service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully",
		zap.String("journey", "CreateUser"))

	c.String(http.StatusOK, "")
}

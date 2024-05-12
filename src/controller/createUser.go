package controller

import (
	"net/http"

	"github.com/MatheusMangueira/go/src/configuration/logger"
	"github.com/MatheusMangueira/go/src/configuration/validation"
	"github.com/MatheusMangueira/go/src/model/request"
	"github.com/MatheusMangueira/go/src/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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

	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	logger.Info("User created successfully",
		zap.String("journey", "CreateUser"))

	c.JSON(http.StatusOK, response)

}

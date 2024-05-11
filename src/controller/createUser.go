package controller

import (
	"fmt"
	"net/http"

	"github.com/MatheusMangueira/go/src/configuration/rest_err"
	"github.com/MatheusMangueira/go/src/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("Invalid JSON body error%s\n", err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, userRequest)

}

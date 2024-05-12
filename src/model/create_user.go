package model

import (
	"fmt"

	"github.com/MatheusMangueira/go/src/configuration/logger"
	"github.com/MatheusMangueira/go/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Init CreateUser model", zap.String("journey", "createUser"))

	ud.EncryptPassword()

	fmt.Println("User created: ", ud)

	return nil
}

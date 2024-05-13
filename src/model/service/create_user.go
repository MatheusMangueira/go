package service

import (
	"fmt"

	"github.com/MatheusMangueira/go/src/configuration/logger"
	"github.com/MatheusMangueira/go/src/configuration/rest_err"
	"github.com/MatheusMangueira/go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init CreateUser model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	fmt.Println("User created: ", userDomain.GetEmail())

	return nil
}

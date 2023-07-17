package service

import (
	"github.com/Fccrelick/crudGO/src/configuration/logger"
	"github.com/Fccrelick/crudGO/src/configuration/rest_err"
	"github.com/Fccrelick/crudGO/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUserServices(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	user, _ := ud.FindUserByEmailServices(userDomain.GetEmail())
	if user != nil {
		return nil, rest_err.NewBadRequestError("Email is already registered in another account")
	}

	userDomain.EncryptPassword()
	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		return nil, err
	}

	return userDomainRepository, nil
}

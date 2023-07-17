package service

import (
	"github.com/Fccrelick/crudGO/src/configuration/logger"
	"github.com/Fccrelick/crudGO/src/configuration/rest_err"
	"github.com/Fccrelick/crudGO/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) LoginUserServices(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init LoginUser model",
		zap.String("journey", "LoginUser"))

	userDomain.EncryptPassword()

	user, err := ud.findUserByEmailAndPasswordServices(
		userDomain.GetEmail(),
		userDomain.GetPassword(),
	)
	if err != nil {
		return nil, rest_err.NewBadRequestError("Email is already registered in another account")
	}

	return user, nil
}

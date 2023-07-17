package service

import (
	"github.com/Fccrelick/crudGO/src/configuration/logger"
	"github.com/Fccrelick/crudGO/src/configuration/rest_err"
	"github.com/Fccrelick/crudGO/src/model"
)

func (ud *userDomainService) FindUserByEmailServices(
	email string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmail Services")

	return ud.userRepository.FindUserByEmail(email)
}

func (ud *userDomainService) FindUserByIDServices(
	id string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByID Services")

	return ud.userRepository.FindUserByID(id)
}

func (ud *userDomainService) findUserByEmailAndPasswordServices(
	email string,
	password string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmailAndPassword Services")

	return ud.userRepository.FindUserByEmailAndPassword(email, password)
}

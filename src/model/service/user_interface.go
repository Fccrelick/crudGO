package service

import (
	"github.com/Fccrelick/crudGO/src/configuration/rest_err"
	"github.com/Fccrelick/crudGO/src/model"
	"github.com/Fccrelick/crudGO/src/model/repository"
)

func NewUserDomainService(
	userRepository repository.UserRepository,
) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUserServices(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)

	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr

	FindUserByEmailServices(
		id string,
	) (model.UserDomainInterface, *rest_err.RestErr)

	FindUserByIDServices(
		id string,
	) (model.UserDomainInterface, *rest_err.RestErr)

	DeleteUser(string) *rest_err.RestErr

	LoginUserServices(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_err.RestErr)
}

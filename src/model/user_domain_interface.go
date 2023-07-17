package model

import "github.com/Fccrelick/crudGO/src/configuration/rest_err"

type UserDomainInterface interface {
	GetId() string
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8

	SetId(string)

	EncryptPassword()
	GenerateToken() (string, *rest_err.RestErr)
}

func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
		age:      age,
	}
}

func NewUserLoginDomain(
	email, password string,
) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
	}
}

func NewUserUpdateDomain(
	name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		name: name,
		age:  age,
	}
}

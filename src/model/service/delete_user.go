package service

import (
	"github.com/Fccrelick/crudGO/src/configuration/logger"
	"github.com/Fccrelick/crudGO/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUser(
	userId string) *rest_err.RestErr {
	logger.Info("Init DeleteUser model",
		zap.String("journey", "DeleteUser"))

	err := ud.userRepository.DeleteUser(userId)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "DeleteUser"))
		return err
	}

	logger.Info(
		"DeleteUser model service executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "DeleteUser"))
	return nil
}

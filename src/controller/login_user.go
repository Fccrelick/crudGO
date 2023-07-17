package controller

import (
	"net/http"

	"github.com/Fccrelick/crudGO/src/configuration/logger"
	"github.com/Fccrelick/crudGO/src/configuration/validation"
	"github.com/Fccrelick/crudGO/src/controller/model/request"
	"github.com/Fccrelick/crudGO/src/model"
	"github.com/Fccrelick/crudGO/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) LoginUser(c *gin.Context) {
	logger.Info("Init LoginUser controller",
		zap.String("journey", "LoginUser"),
	)
	var userRequest request.UserLogin

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "LoginUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserLoginDomain(
		userRequest.Email,
		userRequest.Password,
	)
	domainResult, token, err := uc.service.LoginUserServices(domain)
	if err != nil {
		logger.Error("Error trying to call LoginUser service", err,
			zap.String("journey", "LoginUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"LoginUser controller executed successfully",
		zap.String("userId", domainResult.GetId()),
		zap.String("journey", "LoginUser"))

	c.Header("Authorization", token)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		domainResult,
	))
}

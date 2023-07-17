package controller

import (
	"net/http"
	"net/mail"

	"github.com/Fccrelick/crudGO/src/configuration/logger"
	"github.com/Fccrelick/crudGO/src/configuration/rest_err"
	"github.com/Fccrelick/crudGO/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Init FindUserByEmail controller")

	userEmail := c.Param("userEmail")

	if _, err := mail.ParseAddress(userEmail); err != nil {
		logger.Error("Error trying to validate email",
			err)
		errorMessage := rest_err.NewBadRequestError(
			"userEmail is not a valid email",
		)

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmailServices(userEmail)
	if err != nil {
		logger.Error("Error trying to call findUserByEmail services",
			err)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserByEmail controller executed succesfully")
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		userDomain,
	))
}

func (uc *userControllerInterface) FindUserByID(c *gin.Context) {
	logger.Info("Init FindUserByID controller")

	userId := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error trying to validate userId",
			err)
		errorMessage := rest_err.NewBadRequestError(
			"UserId is not a valid id",
		)

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByIDServices(userId)
	if err != nil {
		logger.Error("Error trying to call findUserById services",
			err)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserById controller executed succesfully")
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		userDomain,
	))
}

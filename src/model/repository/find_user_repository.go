package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/Fccrelick/crudGO/src/configuration/logger"
	"github.com/Fccrelick/crudGO/src/configuration/rest_err"
	"github.com/Fccrelick/crudGO/src/model"
	"github.com/Fccrelick/crudGO/src/model/repository/entity"
	"github.com/Fccrelick/crudGO/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmail(
	email string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmail Repository")

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf(
				"User not found with this Email: %s", email)
			logger.Error(errorMessage,
				err)
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "Error trying to find user by email"
		logger.Error(errorMessage,
			err)
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("FindUserByEmail repository executed successfully",
		zap.String("journey", "findUserByEmail"),
		zap.String("email", email),
		zap.String("userId", userEntity.ID.Hex()))
	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserByID(
	id string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByID Repository")

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	objectID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectID}}
	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf(
				"User not found with this ID: %s", id)
			logger.Error(errorMessage,
				err)
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "Error trying to find user by ID"
		logger.Error(errorMessage,
			err)
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("FindUserByEmail repository executed successfully",
		zap.String("journey", "findUserByID"),
		zap.String("userId", userEntity.ID.Hex()))
	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserByEmailAndPassword(
	email string,
	password string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmailAndPassword Repository")

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{
		{Key: "email", Value: email},
		{Key: "password", Value: password},
	}
	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := "User or password is invalid"
			logger.Error(errorMessage,
				err)
			return nil, rest_err.NewForbiddenError(errorMessage)
		}
		errorMessage := "Error trying to find user by email and password"
		logger.Error(errorMessage,
			err)
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("FindUserByEmailAndPassword repository executed successfully",
		zap.String("journey", "FindUserByEmailAndPassword"),
		zap.String("email", email),
		zap.String("userId", userEntity.ID.Hex()))
	return converter.ConvertEntityToDomain(*userEntity), nil
}

package main

import (
	"github.com/Fccrelick/crudGO/src/controller"
	"github.com/Fccrelick/crudGO/src/model/repository"
	"github.com/Fccrelick/crudGO/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}

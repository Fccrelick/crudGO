package main

import (
	"log"

	"github.com/Fccrelick/crudGO/src/configuration/database/mongodb"
	"github.com/Fccrelick/crudGO/src/controller"
	"github.com/Fccrelick/crudGO/src/controller/routes"
	"github.com/Fccrelick/crudGO/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongodb.InitConnection()

	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

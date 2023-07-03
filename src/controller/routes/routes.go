package routes

import (
	"github.com/Fccrelick/crudGO/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/getUserById/:userId", controller.findUserByID)
	r.GET("/getUserByEmail/:userId", controller.findUserByEmail)
	r.POST("/createUser/", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.updateUser)
	r.DELETE("/deleteUser/:userId", controller.deleteUser)
}
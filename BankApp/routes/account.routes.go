package routes

import (
	"bankDemo/controllers"

	"github.com/gin-gonic/gin"
)

func AccountRoute(router *gin.Engine, controller controllers.AccountController) {
	router.POST("/api/accounts/create", controller.CreateAccount)
	router.GET("/api/accounts/get/:id", controller.GetAccountById)
	router.PUT("/api/accounts/update/:id", controller.UpdateAccountById)
	router.DELETE("/api/accounts/delete/:id", controller.DeleteAccountById)
	router.POST("/api/accounts/createMany", controller.CreateManyAccount)
}

package routes

import (
	"bankDemo/controllers"

	"github.com/gin-gonic/gin"
)

func CustRoute(router *gin.Engine, controller controllers.TransactionController) {
	router.POST("/api/profile/create", controller.CreateTransaction)
	router.GET("/api/profile/get/:id", controller.GetCustomerById)
	router.PUT("/api/profile/update/:id", controller.UpdateCustomerById)
	router.DELETE("/api/profile/delete/:id", controller.DeleteCustomerById)
	router.POST("/api/profile/createMany", controller.CreateManyCustomer)

}
package routes

import (
	"bankDemo/controllers"

	"github.com/gin-gonic/gin"
)

func BankRoute(router *gin.Engine, controller controllers.BankController) {
	router.POST("/api/banks/create", controller.CreateBankid)
	router.GET("/api/banks/get/:id", controller.GetBankid)
	router.PUT("/api/banks/update/:id", controller.UpdateBankid)
	router.DELETE("/api/banks/delete/:id", controller.DeleteBankid)
	router.POST("/api/banks/createMany", controller.CreateManyBankid)
	router.GET("/api/banks/getall", controller.GetallCustomer)
	router.GET("/api/banks/getbyid/:id", controller.GetCustomerbyid)
	router.GET("/api/banks/getbank", controller.GetBank)
}

package routes

import (
	"bankDemo/controllers"

	"github.com/gin-gonic/gin"
)

func Transactionroute(router *gin.Engine,controller controllers.TransactionC){
	router.POST("/api/transfer",controller.Transfer)
}
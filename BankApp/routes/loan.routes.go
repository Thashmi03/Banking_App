package routes

import (
	"bankDemo/controllers"

	"github.com/gin-gonic/gin"
)

func LoanRoute(router *gin.Engine, controller controllers.LoanController) {
	router.POST("/api/loan/createloan", controller.CreateLoan)
	router.GET("/api/loan/getloan/:id", controller.GetLoanById)
	router.PUT("/api/loan/updateloan/:id", controller.UpdateLoanById)
	router.DELETE("/api/loan/deleteloan/:id", controller.DeleteLoanById)
	router.POST("/api/loan/createMany", controller.CreateManyLoan)

}
/*{
        "_id":123,
        "name":"menaha",
        "amount":6000,
        "type":"education"
    }*/
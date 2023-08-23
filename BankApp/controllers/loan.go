package controllers

import (
	"bankDemo/interfaces"
	"bankDemo/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LoanController struct{
	LoanService interfaces.Iloan
}

func InitLoanController (loanService interfaces.Iloan)(LoanController){
	return LoanController{loanService}
}

func (t *LoanController)CreateLoan(ctx *gin.Context){
    var loans *models.Loan 
    if err := ctx.ShouldBindJSON(&loans); err != nil {
        ctx.JSON(http.StatusBadRequest, err.Error())
        return
    }
    newtrans, err := t.LoanService.CreateLoan(loans)
    if(err!=nil){
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

    }
    ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newtrans})

}

func (t *LoanController)GetLoanById(ctx *gin.Context){
    id:= ctx.Param("id")
    id1,err := strconv.ParseInt(id,10,64)
    if(err!=nil){
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

    }
    val, err := t.LoanService.GetLoanById(id1)
    if(err!=nil){
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

    }
    ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": val})
}

func (t *LoanController)UpdateLoanById(ctx *gin.Context){
    id:= ctx.Param("id")
    loan := &models.Loan{}
    if err := ctx.ShouldBindJSON(&loan); err != nil {
        ctx.JSON(http.StatusBadRequest, err.Error())
        return
    }
    id1,err := strconv.ParseInt(id,10,64)
    if(err!=nil){
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
    }
    res,err := t.LoanService.UpdateLoanById(id1,loan)
    if err!=nil{
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
    }
    ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": res})
}

func (t *LoanController)DeleteLoanById(ctx *gin.Context){
    id:= ctx.Param("id")
    id1,err := strconv.ParseInt(id,10,64)
    if(err!=nil){
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
    }
    res,err := t.LoanService.DeleteLoanById(id1)
    if err!=nil{
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
    }
    ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": res})
}
func (t *LoanController)CreateManyLoan(ctx *gin.Context){
    var loans []*models.Loan
    var post *models.Loan
    loans = append(loans, post)
    if err := ctx.ShouldBindJSON(&loans); err != nil {
        fmt.Println("error on controller")
        ctx.JSON(http.StatusBadRequest, err.Error())
        return
    }
    res,err := t.LoanService.CreateManyLoan(loans)
    if err!=nil{
        fmt.Println("error on controller1")

        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
    }
    ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": res})
}
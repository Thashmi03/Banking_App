package controllers

import (
	"bankDemo/interfaces"
	"bankDemo/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	
)

type AccountController struct{
     AccountService  interfaces.IAccount
}


func InitAccountController(accountService interfaces.IAccount) AccountController {
    return AccountController{accountService}
}

func (t *AccountController)CreateAccount(ctx *gin.Context){
	var trans *models.Account  
    if err := ctx.ShouldBindJSON(&trans); err != nil {
        ctx.JSON(http.StatusBadRequest, err.Error())
        return
    }
    newtrans, err := t.AccountService.CreateAccount(trans)
    if(err!=nil){
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

    }
    ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newtrans})

}



func (t *AccountController)GetAccountById(ctx *gin.Context){
    id:= ctx.Param("id")
    id1,err := strconv.ParseInt(id,10,64)
    if(err!=nil){
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
    }
    val, err := t.AccountService.GetAccountById(id1)
    if(err!=nil){
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

    }
    ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": val})
}

func (t *AccountController)UpdateAccountById(ctx *gin.Context){
    id:= ctx.Param("id")
    account := &models.Account{}
    if err := ctx.ShouldBindJSON(&account); err != nil {
        ctx.JSON(http.StatusBadRequest, err.Error())
        return
    }
    id1,err := strconv.ParseInt(id,10,64)
    if(err!=nil){
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
    }
    res,err := t.AccountService.UpdateAccountById(id1,account)
    if err!=nil{
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
    }
    ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": res})
}

func (t *AccountController)DeleteAccountById(ctx *gin.Context){
    id:= ctx.Param("id")
    id1,err := strconv.ParseInt(id,10,64)
    if(err!=nil){
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
    }
    res,err := t.AccountService.DeleteAccountById(id1)
    if err!=nil{
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
    }
    ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": res})
}

func (t *AccountController)CreateManyAccount(ctx *gin.Context){
    var accounts []*models.Account
    if err := ctx.ShouldBindJSON(&accounts); err != nil {
        fmt.Println("error on controller")
        ctx.JSON(http.StatusBadRequest, err.Error())
        return
    }
    res,err := t.AccountService.CreateManyAccount(accounts)
    if err!=nil{
        fmt.Println("error on controller1")

        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
    }
    ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": res})
}
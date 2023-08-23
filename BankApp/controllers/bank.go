package controllers

import (
	"bankDemo/interfaces"
	"bankDemo/models"
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type BankController struct{
     BankService  interfaces.IBank
}


func InitBankController(bankService interfaces.IBank) BankController {
    return BankController{bankService}
}

func (t *BankController)CreateBankid(ctx *gin.Context){
	var trans *models.Bank 
    if err := ctx.ShouldBindJSON(&trans); err != nil {
        ctx.JSON(http.StatusBadRequest, err.Error())
        return
    }
    newtrans, err := t.BankService.CreateBankid(trans)
    if(err!=nil){
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

    }
    ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newtrans})

}



func (t *BankController)GetBankid(ctx *gin.Context){
    id:= ctx.Param("id")
	id1,err := strconv.ParseInt(id,10,64)
   
    if(err!=nil){
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
    }
    val, err := t.BankService.GetBankid(id1)
    if(err!=nil){
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

    }
    ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": val})
}

func (t *BankController)UpdateBankid(ctx *gin.Context){
    id:= ctx.Param("id")
    account := &models.Bank{}
    if err := ctx.ShouldBindJSON(&account); err != nil {
        ctx.JSON(http.StatusBadRequest, err.Error())
        return
    }
    id1,err := strconv.ParseInt(id,10,64)
    if(err!=nil){
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
    }
    res,err := t.BankService.UpdateBankid(id1,account)
    if err!=nil{
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
    }
    ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": res})
}

func (t *BankController)DeleteBankid(ctx *gin.Context){
    id:= ctx.Param("id")
	id1,err := strconv.ParseInt(id,10,64)
    if(err!=nil){
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
    }
    res,err := t.BankService.DeleteBankid(id1)
    if err!=nil{
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
    }
    ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": res})
}

func (t *BankController)CreateManyBankid(ctx *gin.Context){
    var banks []*models.Bank
    if err := ctx.ShouldBindJSON(&banks); err != nil {
        fmt.Println("error on controller")
        ctx.JSON(http.StatusBadRequest, err.Error())
        return
    }
    res,err := t.BankService.CreateManyBankid(banks)
    if err!=nil{
        fmt.Println("error on controller1")

        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
    }
    ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": res})
}

func (t*BankController)GetallCustomer(ctx *gin.Context){
    var orders []bson.M
    new,_:=t.BankService.GetallCustomer()
    if err := new.All(context.Background(), &orders); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
        log.Println("Error:", err)
        return
    }
    ctx.JSON(http.StatusOK, orders)

}
func (t*BankController)GetCustomerbyid(ctx *gin.Context){
    var orders []bson.M
    id:= ctx.Param("id")
	id1,err := strconv.ParseInt(id,10,64)
    if(err!=nil){
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
    }
    new,err:=t.BankService.GetCustomerbyid(id1)
    if err := new.All(context.Background(), &orders); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
        log.Println("Error:", err)
        return
    }
    ctx.JSON(http.StatusOK, orders)
}
func (a *BankController)GetBank(ctx *gin.Context){
	accounts,_:=a.BankService.GetBank()
	ctx.JSON(http.StatusOK,accounts)
}
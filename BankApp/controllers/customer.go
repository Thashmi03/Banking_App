package controllers

import (
	"bankDemo/interfaces"
	"bankDemo/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionController struct{
     TransactionService  interfaces.Icustomer
}

// func InitialiseTransactionController( transaction interfaces.Itransaction)(TransactionController){
//  return TransactionController{transactionService} 
// }

func InitTransController(transactionService interfaces.Icustomer) TransactionController {
    return TransactionController{transactionService}
}

func (t *TransactionController)CreateTransaction(ctx *gin.Context){
    var trans *models.Customer  
    if err := ctx.ShouldBindJSON(&trans); err != nil {
        ctx.JSON(http.StatusBadRequest, err.Error())
        return
    }
    newtrans, err := t.TransactionService.CreateCustomer(trans)
    if(err!=nil){
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

    }
    ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newtrans})

}



func (t *TransactionController)GetCustomerById(ctx *gin.Context){
    id:= ctx.Param("id")
    id1,err := primitive.ObjectIDFromHex(id)
    if(err!=nil){
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

    }
    val, err := t.TransactionService.GetCustomerById(id1)
    if(err!=nil){
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

    }
    ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": val})
}

func (t *TransactionController)UpdateCustomerById(ctx *gin.Context){
    id:= ctx.Param("id")
    customer := &models.Customer{}
    if err := ctx.ShouldBindJSON(&customer); err != nil {
        ctx.JSON(http.StatusBadRequest, err.Error())
        return
    }
    id1,err := primitive.ObjectIDFromHex(id)
    if(err!=nil){
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
    }
    res,err := t.TransactionService.UpdateCustomerById(id1,customer)
    if err!=nil{
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
    }
    ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": res})
}

func (t *TransactionController)DeleteCustomerById(ctx *gin.Context){
    id:= ctx.Param("id")
    id1,err := primitive.ObjectIDFromHex(id)
    if(err!=nil){
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
    }
    res,err := t.TransactionService.DeleteCustomerById(id1)
    if err!=nil{
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
    }
    ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": res})
}

func (t *TransactionController)CreateManyCustomer(ctx *gin.Context){
    var customers []*models.Customer
    var post *models.Customer
    customers = append(customers, post)
    if err := ctx.ShouldBindJSON(&customers); err != nil {
        fmt.Println("error on controller")
        ctx.JSON(http.StatusBadRequest, err.Error())
        return
    }
    res,err := t.TransactionService.CreateManyCustomer(customers)
    if err!=nil{
        fmt.Println("error on controller1")

        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
    }
    ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": res})
}
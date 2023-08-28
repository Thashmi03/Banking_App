package controllers

import (
	"bankDemo/interfaces"
	
	"net/http"

	"github.com/gin-gonic/gin"
)

type transact struct{
	From int64 `json:"from_id" `
	To int64 `json:"to_id"`
    Amount int64`json:"amount"`
}
type TransactionC struct{
	TransactionService  interfaces.Itransact
}
func InitTransactionC(transactionService interfaces.Itransact) TransactionC {
    return TransactionC{transactionService}
}

func (t *TransactionC)Transfer(ctx  *gin.Context ){
	var transfer *transact
    if err := ctx.ShouldBindJSON(&transfer); err != nil {
        ctx.JSON(http.StatusBadRequest, err.Error())
        return
    }
    newtrans, err := t.TransactionService.Transfer(transfer.From,transfer.To,transfer.Amount)
    if(err!=nil){
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

    }
    ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newtrans})
}
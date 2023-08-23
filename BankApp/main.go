package main

import (
	"bankDemo/config"
	"bankDemo/constants"
	"bankDemo/controllers"
	"bankDemo/routes"
	"bankDemo/service"
	"context"
	"fmt"
	"log"

	//	"rest-api/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)


var (
	mongoclient *mongo.Client
	ctx         context.Context
	server         *gin.Engine
)
func initRoutes(){
	routes.Default(server)
}

func initApp(mongoClient *mongo.Client){
	ctx = context.TODO()
	profileCollection := mongoClient.Database("banking").Collection("customer")
	profileService := service.InitCustomer(profileCollection, ctx)
	profileController := controllers.InitTransController(profileService)
	routes.CustRoute(server,profileController)
}
func initloan(mongoClient *mongo.Client){
	ctx = context.TODO()
	profileCollection := mongoClient.Database("banking").Collection("Loan")
	profileService := service.InitLoan(profileCollection, ctx)
	profileController := controllers.InitLoanController(profileService)
	routes.LoanRoute(server,profileController)
}
func InitAccount(mongoClient *mongo.Client){
	ctx = context.TODO()
	profileCollection := mongoClient.Database("banking").Collection("Account")
	profileService := service.InitAccount(profileCollection, ctx)
	profileController := controllers.InitAccountController(profileService)
	routes.AccountRoute(server,profileController)
}
func InitBank(mongoClient *mongo.Client){
	ctx = context.TODO()
	profileCollection := mongoClient.Database("banking").Collection("Bank")
	profileService := service.InitBank(profileCollection, ctx)
	profileController := controllers.InitBankController(profileService)
	routes.BankRoute(server,profileController)
}



func main(){
	server = gin.Default()
	mongoclient,err :=config.ConnectDataBase()
	defer   mongoclient.Disconnect(ctx)
	if err!=nil{
		panic(err)
	}
	initRoutes()
	initApp(mongoclient)
	initloan(mongoclient)
	InitAccount(mongoclient)
	InitBank(mongoclient)
	fmt.Println("server running on port",constants.Port)
	log.Fatal(server.Run(constants.Port))
}
package service

import (
	"bankDemo/interfaces"
	"bankDemo/models"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Account struct{
	ctx context.Context
	mongoCollection *mongo.Collection
}

func InitAccount(collection *mongo.Collection, ctx context.Context) interfaces.IAccount{
	return &Account{ctx,collection}
}
func(c *Account) CreateAccount(user *models.Account)(*mongo.InsertOneResult,error){

	indexModel := mongo.IndexModel{
		Keys:    bson.M{"account_id": 1}, // 1 for ascending, -1 for descending
		Options: options.Index().SetUnique(true),
	}
	_, err := c.mongoCollection.Indexes().CreateOne(c.ctx, indexModel)
	if err != nil {
		log.Fatal(err)
	}
	
	res,err := c.mongoCollection.InsertOne(c.ctx, &user)
	if err!=nil{
		if mongo.IsDuplicateKeyError(err){
			log.Fatal("Duplicate key error")
		}
		return nil,err
	}
	
	return res,nil
}


func(c *Account) GetAccountById(id int64) (*models.Account, error) {
	filter := bson.D{{Key: "account_id", Value: id}}
	var account *models.Account
	res := c.mongoCollection.FindOne(c.ctx, filter)
	err := res.Decode(&account)
	if err!=nil{
		return nil,err
	}
	return account,nil
}

func(c *Account) UpdateAccountById(id int64, account *models.Account) (*mongo.UpdateResult, error){
	iv := bson.M{"account_id": id}
	fv := bson.M{"$set": &account}
	res,err := c.mongoCollection.UpdateOne(c.ctx, iv, fv)
	if err!=nil{
		return nil,err
	}
	return res,nil
}

func (c *Account) DeleteAccountById(id int64) (*mongo.DeleteResult, error){
	del := bson.M{"account_id": id}
	res,err := c.mongoCollection.DeleteOne(c.ctx, del)
	if err!=nil{
		return nil,err
	}
	return res,nil
}

func (c *Account) CreateManyAccount(post []*models.Account)(*mongo.InsertManyResult,error){
	var users []interface{}
	for _,user := range post{		
		users = append(users, user)
	}
	res,err := c.mongoCollection.InsertMany(c.ctx, users)
	// fmt.Println(user)
	if err!=nil{
		fmt.Println("error in service")
		return nil,err
	}
	return res,nil
}
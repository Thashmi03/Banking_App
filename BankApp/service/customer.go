package service

import (
	"bankDemo/interfaces"
	"bankDemo/models"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type Cust struct{
	ctx context.Context
	mongoCollection *mongo.Collection
}

func InitCustomer(collection *mongo.Collection, ctx context.Context) interfaces.Icustomer{
	return &Cust{ctx,collection}
}
func(c *Cust) CreateCustomer(user *models.Customer)(*mongo.InsertOneResult,error){
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"account_id": 1}, // 1 for ascending, -1 for descending
		Options: options.Index().SetUnique(true),
	}
	_, err := c.mongoCollection.Indexes().CreateOne(c.ctx, indexModel)
	if err != nil {
		log.Fatal(err)
	}
	user.Customer_ID = primitive.NewObjectID()
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password),7)
	user.Password = string(hashedPassword)
	res,err := c.mongoCollection.InsertOne(c.ctx, &user)
	if err!=nil{
		if mongo.IsDuplicateKeyError(err){
			log.Fatal("Duplicate key error")
		}
		return nil,err
	}
	
	return res,nil
}


func(c *Cust) GetCustomerById(id primitive.ObjectID) (*models.Customer, error) {
	filter := bson.D{{Key: "customer_id", Value: id}}
	var customer *models.Customer
	res := c.mongoCollection.FindOne(c.ctx, filter)
	err := res.Decode(&customer)
	if err!=nil{
		return nil,err
	}
	return customer,nil
}

func(c *Cust) UpdateCustomerById(id primitive.ObjectID, customer *models.Customer) (*mongo.UpdateResult, error){
	iv := bson.M{"customer_id": id}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(customer.Password),8)
	customer.Password = string(hashedPassword)
	fv := bson.M{"$set": &customer}
	res,err := c.mongoCollection.UpdateOne(c.ctx, iv, fv)
	if err!=nil{
		return nil,err
	}
	return res,nil
}

func (c *Cust) DeleteCustomerById(id primitive.ObjectID) (*mongo.DeleteResult, error){
	del := bson.M{"customer_id": id}
	res,err := c.mongoCollection.DeleteOne(c.ctx, del)
	if err!=nil{
		return nil,err
	}
	return res,nil
}

func (c *Cust) CreateManyCustomer(post []*models.Customer)(*mongo.InsertManyResult,error){
	var users []interface{}
	for _,user := range post{
		user.Customer_ID = primitive.NewObjectID()
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password),8)
		user.Password = string(hashedPassword)
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


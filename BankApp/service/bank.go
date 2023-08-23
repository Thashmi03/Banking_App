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

type Bank struct{
	ctx context.Context
	mongoCollection *mongo.Collection
}

func InitBank(collection *mongo.Collection, ctx context.Context) interfaces.IBank{
	return &Bank{ctx,collection}
}
func(c *Bank) CreateBankid(user *models.Bank)(*mongo.InsertOneResult,error){

	indexModel := mongo.IndexModel{
		Keys:    bson.M{"bank_id": 1}, // 1 for ascending, -1 for descending
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

func (c*Bank) GetallCustomer()(*mongo.Cursor,error){
	pipeline := mongo.Pipeline{
        {
            {Key: "$lookup", Value: bson.D{
                {Key: "from", Value: "customer"},
                {Key: "localField", Value: "bank_id"},
                {Key: "foreignField", Value: "bank_id"},
                {Key: "as", Value: "new_bank_id"},
            }},
        },
        {
            {Key: "$unwind", Value: "$new_bank_id"},
        },
        {
            {Key: "$project", Value: bson.D{  
				{Key: "new_bank_id.customer_id", Value: 1},
                {Key: "new_bank_id.name", Value: 1},
				{Key: "new_bank_id.account_id", Value: 1},
            }},
        },
    }
    
    // Perform aggregation
    res, err := c.mongoCollection.Aggregate(c.ctx, pipeline)
    if err != nil {
        log.Fatal(err)
	}
	return res,nil
}
func(a *Bank)GetBank()([]*models.Bank,error){
	filter:=bson.D{}
	options:=options.Find()
	res,_:=a.mongoCollection.Find(a.ctx,filter,options)
	var bank[]*models.Bank
	for res.Next(a.ctx){
		acc:=&models.Bank{}
		err:=res.Decode(acc)
		if err!=nil{
			return nil,err
		}
		bank=append(bank, acc)
	}
	return bank,nil
}
func (c*Bank) GetCustomerbyid(id int64)(*mongo.Cursor,error){
	pipeline := mongo.Pipeline{
        {
            {Key: "$lookup", Value: bson.D{
                {Key: "from", Value: "customer"},
                {Key: "localField", Value: "bank_id"},
                {Key: "foreignField", Value: "bank_id"},
                {Key: "as", Value: "new_bank_id"},
            }},
        },
        {
            {Key: "$unwind", Value: "$new_bank_id"},
        },
		{
			{Key: "$match",Value: bson.D{{Key: "bank_id",Value:12345 }}},
		},
        {
            {Key: "$project", Value: bson.D{  
				{Key: "new_bank_id.customer_id", Value: 1},
                {Key: "new_bank_id.name", Value: 1},
				{Key: "new_bank_id.account_id", Value: 1},
            }},
        },
    }
    
    // Perform aggregation
    res, err := c.mongoCollection.Aggregate(c.ctx, pipeline)
    if err != nil {
        log.Fatal(err)
	}
	return res,nil
}

func(c *Bank) GetBankid(id int64) (*models.Bank, error) {
	filter := bson.D{{Key: "bank_id", Value: id}}
	var bank *models.Bank
	res := c.mongoCollection.FindOne(c.ctx, filter)
	err := res.Decode(&bank)
	if err!=nil{
		return nil,err
	}
	return bank,nil
}

func(c *Bank) UpdateBankid(id int64, bank *models.Bank) (*mongo.UpdateResult, error){
	iv := bson.M{"bank_id": id}
	fv := bson.M{"$set": &bank}
	res,err := c.mongoCollection.UpdateOne(c.ctx, iv, fv)
	if err!=nil{
		return nil,err
	}
	return res,nil
}

func (c *Bank) DeleteBankid(id int64) (*mongo.DeleteResult, error){
	del := bson.M{"bank_id": id}
	res,err := c.mongoCollection.DeleteOne(c.ctx, del)
	if err!=nil{
		return nil,err
	}
	return res,nil
}

func (c *Bank) CreateManyBankid(post []*models.Bank)(*mongo.InsertManyResult,error){
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
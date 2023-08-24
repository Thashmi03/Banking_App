package service

import (
	"bankDemo/interfaces"
	"bankDemo/models"
	"context"
	"fmt"
	"log"
	"time"

	// "time"

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
	date := time.Now()
	for i:=0;i<len(user.Transaction);i++{
		user.Transaction[i].Date = date.Format("2006-01-02 12.50.00.000000000")
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
		date := time.Now()
		for i:=0;i<len(user.Transaction);i++{
			user.Transaction[i].Date = date.Format("2006-01-02 12.50.00.000000000")
		}
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

func (b *Cust) GetAllBankTransSum(date1 string, date2 string, id primitive.ObjectID) (int64, error) {
	// pipeline := []bson.M{
		
	// 	{
	// 		"$match": bson.M{
	// 			"date": bson.M{
	// 				"$gte": date1,
	// 				"$lte": date2,
	// 			},
	// 		},
	// 	},
	// 	{            "$unwind": "$transaction",        },
	// 	{
	// 		"$group": bson.M{
	// 			"_id": "",
	// 			"totalAmount": bson.M{"$sum": "$transaction.transaction_amount"},
	// 		},
	// 	},
	
		
	// }
	// cursor, err := b.mongoCollection.Aggregate(b.ctx, pipeline)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	

	// var results []bson.M
	// if err := cursor.All(b.ctx, &results); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("hi",results)
	// if len(results) > 0 {
	// 	totalSum := results[0]["total"].(float64)
	// 	fmt.Println(totalSum)
	// }
	// return 0,nil
	pipeline := []bson.M{
		{
			"$match": bson.M{
				"date": bson.M{
					"$gte": date1,
					"$lte": date2,
				},
				// "customer_id":id, // Replace with the actual customer ID
			},
		},     
		{           
			"$unwind": "$transaction",        
		}, 
		{            
			"$group": bson.M{               
				 "_id": "$customer_id",               
				  "total": bson.M{"$sum": "$transaction.transaction_amount",},           
				   },        
				   },    
		}   
		// res1, err:= b.mongoCollection.Aggregate(b.ctx, pipeline)   
		// if err!=nil{        
		// 	return 0, err   
		//  }    
		//  var re []bson.M   

		//   if err := res1.All(b.ctx, &re); err != nil {   
		//     return 0, err    
		// 	}  

		// total:=re[0]["total"].(int64)
		// return total,nil
		cursor, err := b.mongoCollection.Aggregate(context.Background(), pipeline)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	var results []bson.M
	if err := cursor.All(context.Background(), &results); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Total sum of amounts by customer within the specified date range:")
	for _, result := range results {
		totalAmount := result["totalAmount"].(int64)
		fmt.Printf("Customer: %s, Total Amount: %d\n", result["_id"], totalAmount)
	}
	return 0,nil
}


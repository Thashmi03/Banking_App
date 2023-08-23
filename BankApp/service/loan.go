package service

import (
	"bankDemo/interfaces"
	"bankDemo/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Loan struct{
	ctx context.Context
	mongoCollection *mongo.Collection
}

func InitLoan(collection *mongo.Collection, ctx context.Context) interfaces.Iloan{
	return &Loan{ctx,collection}
}

func (c* Loan)CreateLoan(user *models.Loan)(*mongo.InsertOneResult,error){

	res,err := c.mongoCollection.InsertOne(c.ctx,&user)
	 if err!= nil{
		return nil,err
	 }
	 return res, nil
}

func (c *Loan)GetLoanById(id int64)([]*models.Loan,error){
	match:=bson.D{{Key:"_id",Value: id}}
	result,err:=c.mongoCollection.Find(c.ctx,match)
	if err!=nil{
		return nil, err
	}else{
		var Loan_detail[] *models.Loan
		for result.Next(c.ctx){
			detail:=&models.Loan{}
			err:=result.Decode(detail)
			if err!=nil{
				return nil, err
			}
			Loan_detail=append(Loan_detail, detail)
		}
		return Loan_detail,nil
	}
}

func (c *Loan) UpdateLoanById(id int64, loan *models.Loan) (*mongo.UpdateResult, error){
	iv := bson.M{"_id": id}
	fv := bson.M{"$set": &loan}
	res,err := c.mongoCollection.UpdateOne(c.ctx, iv, fv)
	if err!=nil{
		return nil,err
	}
	return res,nil
}

func (c *Loan) DeleteLoanById(id int64) (*mongo.DeleteResult, error){	
	del := bson.M{"_id": id}
    res,err:=c.mongoCollection.DeleteOne(c.ctx,del)
	if err!=nil{
		return nil,err
	}
	return res, nil
}
func (c *Loan) CreateManyLoan(post []*models.Loan)(*mongo.InsertManyResult,error){
	
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
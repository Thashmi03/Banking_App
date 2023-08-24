package models

import (


	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
	Customer_ID primitive.ObjectID `json:"customer_id" bson:"customer_id"`
	Bank_ID     int64              `json:"bank_id" bson:"bank_id"`
	Name        string             `json:"name" bson:"name"`
	Password    string             `json:"password" bson:"password"`
	Account_ID  int64              `json:"account_id" bson:"account_id"`
	Transaction []CustTransaction  `json:"transaction" bson:"transaction"`
	// Accounts	[]Account		   `json:"account" bson:"account"`
}

type CustTransaction struct {
	Transaction_id     int64     `json:"transaction_id" bson:"transaction_id"`
	Transaction_amount int64     `json:"transaction_amount" bson:"transaction_amount"`
	Date               string     `json:"date" bson:"date"`
	Transaction_type   string    `json:"transaction_type" bson:"transaction_type"`
}

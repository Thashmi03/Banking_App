package models

// import "time"

type Transaction struct{
	Transaction_id  string `json:"transaction_id" bson:"transaction_id"`
	From_account int64	 `json:"from_account" bson:"from_account"`
	To_account int64	 `json:"to_account" bson:"to_account"`
	Amount int64	     `json:"amount" bson:"amount"`
	// Time time.Time	     `json:"time" bson:"time"`
}
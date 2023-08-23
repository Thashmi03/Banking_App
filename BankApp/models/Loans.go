package models


type Loan struct{
	Id int64 `json:"_id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	Amount int64 `json:"amount" bson:"amount"`
	Type string `json:"type" bson:"type"`
}
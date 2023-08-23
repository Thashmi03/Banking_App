package models


type Account struct{
	Account_ID int64 `json:"account_id" bson:"account_id"`
	Account_type string `json:"account_type" bson:"account_type"`
	Branch string `json:"branch" bson:"branch"`
}
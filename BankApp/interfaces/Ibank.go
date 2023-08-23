package interfaces

import (
	"bankDemo/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type IBank interface{
	CreateBankid(*models.Bank)(*mongo.InsertOneResult,error)
	CreateManyBankid([]*models.Bank)(*mongo.InsertManyResult,error)
	GetBankid(int64) (*models.Bank, error)
	UpdateBankid(int64, *models.Bank) (*mongo.UpdateResult, error)
	DeleteBankid(int64) (*mongo.DeleteResult, error)
	GetallCustomer()(*mongo.Cursor,error)
	GetCustomerbyid(int64)(*mongo.Cursor,error)
	GetBank()([]*models.Bank,error)
}
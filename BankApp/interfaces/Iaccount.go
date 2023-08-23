package interfaces

import (
	"bankDemo/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type IAccount interface {
	CreateAccount(*models.Account) (*mongo.InsertOneResult, error)
	CreateManyAccount([]*models.Account) (*mongo.InsertManyResult, error)
	GetAccountById(int64) (*models.Account, error)
	UpdateAccountById(int64, *models.Account) (*mongo.UpdateResult, error)
	DeleteAccountById(int64) (*mongo.DeleteResult, error)
}


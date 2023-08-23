package interfaces

import (
	"bankDemo/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Icustomer interface {
	CreateCustomer(*models.Customer)(*mongo.InsertOneResult,error)
	CreateManyCustomer([]*models.Customer)(*mongo.InsertManyResult,error)
	GetCustomerById(primitive.ObjectID) (*models.Customer, error)
	UpdateCustomerById(primitive.ObjectID, *models.Customer) (*mongo.UpdateResult, error)
	DeleteCustomerById(primitive.ObjectID) (*mongo.DeleteResult, error)
}
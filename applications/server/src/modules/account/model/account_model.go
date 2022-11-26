package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateServiceModel struct {
	AccountID   primitive.ObjectID
	Name        string
	Description string
}

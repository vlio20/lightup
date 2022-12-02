package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateServingModel struct {
	AccountID primitive.ObjectID
	Name      string
	Email     string
	Password  string
}

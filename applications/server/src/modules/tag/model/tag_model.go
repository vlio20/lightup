package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateTagModel struct {
	AccountID   primitive.ObjectID
	Name        string
	Description string
}

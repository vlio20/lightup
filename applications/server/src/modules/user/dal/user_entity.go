package dal

import (
	"lightup/src/common/db"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserEntity struct {
	db.BaseEntity `bson:",inline"`
	AccountID     primitive.ObjectID `bson:"accountId, omitempty"`
	Name          string             `bson:"name, omitempty"`
	Email         string             `bson:"email, omitempty"`
	Password      string             `bson:"password, omitempty"`
	Archived      bool               `bson:"archived, omitempty"`
}

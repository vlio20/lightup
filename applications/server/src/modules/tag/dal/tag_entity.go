package dal

import (
	"lightup/src/common/db"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TagEntity struct {
	db.BaseEntity `bson:",inline"`
	AccountID     primitive.ObjectID `bson:"accountId, omitempty"`
	Name          string             `bson:"name, omitempty"`
	Description   string             `bson:"description, omitempty"`
	Archived      bool               `bson:"archived, omitempty"`
}

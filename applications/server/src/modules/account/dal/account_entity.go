package dal

import (
	"lightup/src/common/db"
)

type AccountEntity struct {
	db.BaseEntity `bson:",inline"`
	Name          string `bson:"name, omitempty"`
	Description   string `bson:"description, omitempty"`
	Archived      bool   `bson:"archived, omitempty"`
}

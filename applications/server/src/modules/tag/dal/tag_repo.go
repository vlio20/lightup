package dal

import (
	"lightup/src/common/db"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var getDb = db.GetDB

type TagRepo struct {
	db.Repository[TagEntity]
}

func NewTagRepository() *TagRepo {
	return &TagRepo{
		Repository: db.Repository[TagEntity]{
			DB:         getDb(),
			Collection: getDb().Collection("tag"),
		},
	}
}

func (repo *TagRepo) GetTag(accountId primitive.ObjectID, name string) (*TagEntity, error) {

	var entity = &TagEntity{
		AccountID: accountId,
		Name:      name,
	}

	return repo.FindOne(entity)
}

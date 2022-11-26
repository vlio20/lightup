package dal

import (
	"lightup/src/common/db"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var getDb = db.GetDB

type ServiceRepo struct {
	db.Repository[ServiceEntity]
}

func NewServiceRepository() *ServiceRepo {
	return &ServiceRepo{
		Repository: db.Repository[ServiceEntity]{
			DB:         getDb(),
			Collection: getDb().Collection("service"),
		},
	}
}

func (repo *ServiceRepo) GetService(accountId primitive.ObjectID, name string) (*ServiceEntity, error) {

	var entity = &ServiceEntity{
		AccountID: accountId,
		Name:      name,
	}

	return repo.FindOne(entity)
}

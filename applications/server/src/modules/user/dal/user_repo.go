package dal

import (
	"go.mongodb.org/mongo-driver/bson"
	"lightup/src/common/db"
)

type UserRepo struct {
	db.Repository[UserEntity]
}

func New() *UserRepo {
	return &UserRepo{
		Repository: db.Repository[UserEntity]{
			DB:         db.GetDB(),
			Collection: db.GetDB().Collection("user"),
		},
	}
}

func (repo *UserRepo) FindByEmail(email string) (*UserEntity, error) {
	return repo.FindOne(bson.M{"email": email})
}

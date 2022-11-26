package dal

import (
	"lightup/src/common/db"
)

type UserRepo struct {
	db.Repository[UserEntity]
}

func NewUserRepository() *UserRepo {
	return &UserRepo{
		Repository: db.Repository[UserEntity]{
			DB:         db.GetDB(),
			Collection: db.GetDB().Collection("user"),
		},
	}
}

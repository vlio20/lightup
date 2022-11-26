package dal

import (
	"lightup/src/common/db"
)

type AccountRepo struct {
	db.Repository[AccountEntity]
}

func NewAccountRepository() *AccountRepo {
	return &AccountRepo{
		Repository: db.Repository[AccountEntity]{
			DB:         db.GetDB(),
			Collection: db.GetDB().Collection("account"),
		},
	}
}

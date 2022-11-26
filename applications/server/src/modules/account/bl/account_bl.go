package bl

import (
	"lightup/src/common/db"
	"lightup/src/common/log"
	"lightup/src/modules/account/dal"
	"lightup/src/modules/account/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccountImpl struct {
	log         log.Logger
	AccountRepo *dal.AccountRepo
}

type AccountBl interface {
	GetAccountById(id primitive.ObjectID) (*dal.AccountEntity, error)
	CreateAccount(input *model.CreateAccountModel) (*dal.AccountEntity, error)
}

func New() AccountBl {
	return &AccountImpl{
		log:         log.GetLogger("AccountBl"),
		AccountRepo: dal.NewAccountRepository(),
	}
}

func (impl *AccountImpl) GetAccountById(id primitive.ObjectID) (*dal.AccountEntity, error) {
	return impl.AccountRepo.GetByObjectId(&id)
}

func (impl *AccountImpl) CreateAccount(input *model.CreateAccountModel) (*dal.AccountEntity, error) {
	entity := &dal.AccountEntity{
		BaseEntity:  *db.GetBaseEntity(),
		Name:        input.Name,
		Description: input.Description,
		Archived:    false,
	}

	return impl.AccountRepo.Add(entity)
}

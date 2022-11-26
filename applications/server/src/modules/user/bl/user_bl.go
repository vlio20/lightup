package bl

import (
	"lightup/src/common/db"
	"lightup/src/common/log"
	"lightup/src/modules/user/dal"
	"lightup/src/modules/user/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserImpl struct {
	log      log.Logger
	UserRepo *dal.UserRepo
}

type UserBl interface {
	GetUserById(id primitive.ObjectID) (*dal.UserEntity, error)
	CreateUser(input *model.CreateUserModel) (*dal.UserEntity, error)
}

func New() UserBl {
	return &UserImpl{
		log:      log.GetLogger("UserBl"),
		UserRepo: dal.NewUserRepository(),
	}
}

func (impl *UserImpl) GetUserById(id primitive.ObjectID) (*dal.UserEntity, error) {
	return impl.UserRepo.GetByObjectId(&id)
}

func (impl *UserImpl) CreateUser(input *model.CreateUserModel) (*dal.UserEntity, error) {
	entity := &dal.UserEntity{
		BaseEntity: *db.GetBaseEntity(),
		Name:       input.Name,
		Email:      input.Email,
		Password:   input.Password,
		Archived:   false,
	}

	return impl.UserRepo.Add(entity)
}

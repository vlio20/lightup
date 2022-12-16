package bl

import (
	"lightup/src/common/db"
	"lightup/src/common/hasher"
	"lightup/src/common/http"
	"lightup/src/common/log"
	"lightup/src/modules/user/dal"
	"lightup/src/modules/user/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserImpl struct {
	log      log.Logger
	userRepo *dal.UserRepo
	hasher   hasher.Hasher
}

type UserBl interface {
	GetUserById(id primitive.ObjectID) (*dal.UserEntity, error)
	CreateUser(input *model.CreateUserModel) (*dal.UserEntity, error)
}

func New() *UserImpl {
	return &UserImpl{
		log:      log.GetLogger("UserBl"),
		userRepo: dal.New(),
		hasher:   hasher.New(),
	}
}

func (impl *UserImpl) GetUserById(id primitive.ObjectID) (*dal.UserEntity, error) {
	return impl.userRepo.GetByObjectId(&id)
}

func (impl *UserImpl) CreateUser(input *model.CreateUserModel) (*dal.UserEntity, error) {
	existingUser, err := impl.userRepo.FindByEmail(input.Email)

	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		return nil, http.Error{StatusCode: 400, Message: "User already exists"}
	}

	hashedPassword, err := impl.hasher.Hash(input.Password)

	if err != nil {
		return nil, err
	}

	entity := &dal.UserEntity{
		BaseEntity: *db.GetBaseEntity(),
		Name:       input.Name,
		Email:      input.Email,
		Password:   hashedPassword,
		Archived:   false,
	}

	return impl.userRepo.Add(entity)
}

package bl

import (
	"lightup/src/common/db"
	"lightup/src/common/log"
	"lightup/src/modules/service/dal"
	"lightup/src/modules/service/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ServiceImpl struct {
	log         log.Logger
	ServiceRepo *dal.ServiceRepo
}

type ServiceBl interface {
	GetServiceById(id primitive.ObjectID) (*dal.ServiceEntity, error)
	CreateService(input *model.CreateServiceModel) (*dal.ServiceEntity, error)
	GetService(accountId primitive.ObjectID, name string) (*dal.ServiceEntity, error)
}

func New() ServiceBl {
	return &ServiceImpl{
		log:         log.GetLogger("ServiceBl"),
		ServiceRepo: dal.NewServiceRepository(),
	}
}

func (impl *ServiceImpl) GetServiceById(id primitive.ObjectID) (*dal.ServiceEntity, error) {
	return impl.ServiceRepo.GetByObjectId(&id)
}

func (impl *ServiceImpl) CreateService(input *model.CreateServiceModel) (*dal.ServiceEntity, error) {
	entity := &dal.ServiceEntity{
		BaseEntity:  *db.GetBaseEntity(),
		AccountID:   input.AccountID,
		Name:        input.Name,
		Description: input.Description,
		Archived:    false,
	}

	return impl.ServiceRepo.Add(entity)
}

func (impl *ServiceImpl) GetService(accountId primitive.ObjectID, name string) (*dal.ServiceEntity, error) {
	entity, err := impl.ServiceRepo.GetService(accountId, name)

	if err != nil {
		return nil, err
	}

	return entity, nil
}

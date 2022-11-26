package bl

import (
	"lightup/src/common/db"
	"lightup/src/common/log"
	"lightup/src/modules/feature_flag/dal"
	"lightup/src/modules/feature_flag/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FeatureFlagImpl struct {
	log             log.Logger
	FeatureFlagRepo *dal.FeatureFlagRepo
}

type FeatureFlagBl interface {
	GetFeatureFlagById(id primitive.ObjectID) (*dal.FeatureFlagEntity, error)
	CreateFeatureFlag(input *model.CreateFeatureFlagDto) (*dal.FeatureFlagEntity, error)
	GetFeatureFlag(accountId primitive.ObjectID, serviceId primitive.ObjectID, name string) (*dal.FeatureFlagEntity, error)
}

func New() FeatureFlagBl {
	return &FeatureFlagImpl{
		log:             log.GetLogger("FeatureFlagBl"),
		FeatureFlagRepo: dal.NewFeatureFlagRepository(),
	}
}

func (impl *FeatureFlagImpl) GetFeatureFlagById(id primitive.ObjectID) (*dal.FeatureFlagEntity, error) {
	return impl.FeatureFlagRepo.GetByObjectId(&id)
}

func (impl *FeatureFlagImpl) CreateFeatureFlag(input *model.CreateFeatureFlagDto) (*dal.FeatureFlagEntity, error) {
	entity := &dal.FeatureFlagEntity{
		BaseEntity:  *db.GetBaseEntity(),
		AccountID:   input.AccountID,
		Name:        input.Name,
		Description: input.Description,
		ServiceID:   input.ServiceID,
		Archived:    input.Archived,
		Config:      input.Config,
	}

	return impl.FeatureFlagRepo.Add(entity)
}

func (impl *FeatureFlagImpl) GetFeatureFlag(accountId primitive.ObjectID, serviceId primitive.ObjectID, name string) (*dal.FeatureFlagEntity, error) {
	entity, err := impl.FeatureFlagRepo.GetFeatureFlag(accountId, serviceId, name)

	if err != nil {
		return nil, err
	}

	return entity, nil
}

package bl

import (
	"lightup/src/common/db"
	"lightup/src/common/log"
	"lightup/src/modules/feature_flag/ff_dal"
	"lightup/src/modules/feature_flag/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FeatureFlagImpl struct {
	log             log.Logger
	FeatureFlagRepo *ff_dal.FeatureFlagRepo
}

type FeatureFlagBl interface {
	GetFeatureFlagById(id primitive.ObjectID) (*ff_dal.FeatureFlagEntity, error)
	CreateFeatureFlag(input *model.CreateFeatureFlagInput) (*ff_dal.FeatureFlagEntity, error)
	GetFeatureFlag(accountId primitive.ObjectID, name string) (*ff_dal.FeatureFlagEntity, error)
}

func New() FeatureFlagBl {
	return &FeatureFlagImpl{
		log:             log.GetLogger("FeatureFlagBl"),
		FeatureFlagRepo: ff_dal.NewFeatureFlagRepository(),
	}
}

func (impl *FeatureFlagImpl) GetFeatureFlagById(id primitive.ObjectID) (*ff_dal.FeatureFlagEntity, error) {
	return impl.FeatureFlagRepo.GetByObjectId(&id)
}

func (impl *FeatureFlagImpl) CreateFeatureFlag(input *model.CreateFeatureFlagInput) (*ff_dal.FeatureFlagEntity, error) {
	entity := &ff_dal.FeatureFlagEntity{
		BaseEntity:  *db.GetBaseEntity(),
		AccountID:   input.AccountID,
		Name:        input.Name,
		Description: input.Description,
		Tags:        input.Tags,
		Archived:    input.Archived,
		Config:      input.Config,
	}

	return impl.FeatureFlagRepo.Add(entity)
}

func (impl *FeatureFlagImpl) GetFeatureFlag(accountId primitive.ObjectID, name string) (*ff_dal.FeatureFlagEntity, error) {
	entity, err := impl.FeatureFlagRepo.GetFeatureFlag(accountId, name)

	if err != nil {
		return nil, err
	}

	return entity, nil
}

package bl

import (
	"lightup/src/common/db"
	"lightup/src/common/log"
	"lightup/src/modules/feature_flag/dal"
	"lightup/src/modules/feature_flag/model"
)

type FeatureFlagImpl struct {
	log             log.Logger
	FeatureFlagRepo *dal.FeatureFlagRepo
}

type FeatureFlagBl interface {
	GetFeatureFlagById(id string) (*dal.FeatureFlagEntity, error)
	CreateFeatureFlag(input *model.CreateFeatureFlagDto) (*dal.FeatureFlagEntity, error)
}

func New() FeatureFlagBl {
	return &FeatureFlagImpl{
		log:             log.GetLogger("FeatureFlagBl"),
		FeatureFlagRepo: dal.NewFeatureFlagRepository(),
	}
}

func (bl *FeatureFlagImpl) GetFeatureFlagById(id string) (*dal.FeatureFlagEntity, error) {
	return bl.FeatureFlagRepo.GetById(id)
}

func (bl *FeatureFlagImpl) CreateFeatureFlag(input *model.CreateFeatureFlagDto) (*dal.FeatureFlagEntity, error) {
	entity := &dal.FeatureFlagEntity{
		BaseEntity:  *db.GetBaseEntity(),
		Name:        input.Name,
		Description: input.Description,
		Archived:    input.Archived,
		Config:      input.Config,
	}

	return bl.FeatureFlagRepo.Add(entity)
}

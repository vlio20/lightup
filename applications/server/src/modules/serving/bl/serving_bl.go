package bl

import (
	"lightup/src/common/log"
	"lightup/src/modules/feature_flag/ff_dal"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ServingImpl struct {
	log    log.Logger
	ffRepo ff_dal.FeatureFlagRepo
}

type ServingBl interface {
}

func New() ServingBl {
	return &ServingImpl{
		log:    log.GetLogger("ServingBl"),
		ffRepo: *ff_dal.NewFeatureFlagRepository(),
	}
}

func (bl *ServingImpl) GetFlagStatus(accountId primitive.ObjectID, ffName string) (bool, error) {
	return true, nil
}

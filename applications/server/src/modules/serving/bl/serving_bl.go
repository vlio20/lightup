package bl

import (
	"lightup/src/common/http"
	"lightup/src/common/log"
	"lightup/src/modules/feature_flag/ff_dal"
	"lightup/src/modules/serving/model"
	"math/rand"
	"net/url"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ServingImpl struct {
	log    log.Logger
	ffRepo ff_dal.FlagRepo
	rand   func() float32
}

type ServingBl interface {
	GetFeatureFlagState(accountID primitive.ObjectID, params *model.FlagStateParams, query url.Values) (bool, error)
}

func New() ServingBl {
	return &ServingImpl{
		log:    log.GetLogger("ServingBl"),
		ffRepo: *ff_dal.NewFlagRepository(),
		rand:   rand.Float32,
	}
}

func (impl *ServingImpl) GetFeatureFlagState(
	accountID primitive.ObjectID,
	params *model.FlagStateParams,
	query url.Values,
) (bool, error) {
	flag, err := impl.ffRepo.GetFlagByName(accountID, params.FlagName)

	if err != nil {
		return false, err
	}

	if flag == nil {
		return false, http.Error{
			StatusCode: 404,
			Message:    "Flag not found"}
	}

	match, err := flag.Config.Segment.IsMatch(query)

	if err != nil {
		return false, err
	}

	random := rand.Float32()

	return match && random <= flag.Config.Percentage, nil
}

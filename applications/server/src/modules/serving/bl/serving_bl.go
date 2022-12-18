package bl

import (
	"lightup/src/common/hasher"
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
	hasher hasher.Hasher
}

type ServingBl interface {
	GetFeatureFlagState(accountID primitive.ObjectID, params *model.FlagStateParams, query url.Values) (bool, error)
}

func New() ServingBl {
	return &ServingImpl{
		log:    log.GetLogger("ServingBl"),
		ffRepo: *ff_dal.NewFlagRepository(),
		rand:   rand.Float32,
		hasher: hasher.New(),
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

	if flag.Config.Persistent {
		return impl.getPersistentFlagState(flag, query, match)
	} else {
		return impl.getNonPersistentFlagState(flag, match)
	}

}

func (impl *ServingImpl) getPersistentFlagState(
	flag *ff_dal.FeatureFlagEntity,
	query url.Values,
	match bool,
) (bool, error) {
	identifier := query.Get(flag.Config.Identifier)

	if identifier == "" {
		return false, http.Error{
			StatusCode: 400,
			Message:    "Missing identifier. Key: " + flag.Config.Identifier,
		}
	}

	seed := flag.ID.Hex()
	randVal, err := impl.hasher.HashStringToFloat(identifier, seed)

	if err != nil {
		return false, err
	}

	return match && randVal <= flag.Config.Percentage, nil
}

func (impl *ServingImpl) getNonPersistentFlagState(flag *ff_dal.FeatureFlagEntity, match bool) (bool, error) {
	return match && impl.rand() <= flag.Config.Percentage, nil
}

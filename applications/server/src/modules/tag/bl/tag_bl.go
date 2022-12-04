package bl

import (
	"lightup/src/common/db"
	"lightup/src/common/log"
	"lightup/src/modules/tag/dal"
	"lightup/src/modules/tag/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TagImpl struct {
	log         log.Logger
	TagRepo *dal.TagRepo
}

type TagBl interface {
	GetTagById(id primitive.ObjectID) (*dal.TagEntity, error)
	CreateTag(input *model.CreateTagModel) (*dal.TagEntity, error)
	GetTag(accountId primitive.ObjectID, name string) (*dal.TagEntity, error)
}

func New() TagBl {
	return &TagImpl{
		log:         log.GetLogger("TagBl"),
		TagRepo: dal.NewTagRepository(),
	}
}

func (impl *TagImpl) GetTagById(id primitive.ObjectID) (*dal.TagEntity, error) {
	return impl.TagRepo.GetByObjectId(&id)
}

func (impl *TagImpl) CreateTag(input *model.CreateTagModel) (*dal.TagEntity, error) {
	entity := &dal.TagEntity{
		BaseEntity:  *db.GetBaseEntity(),
		AccountID:   input.AccountID,
		Name:        input.Name,
		Description: input.Description,
		Archived:    false,
	}

	return impl.TagRepo.Add(entity)
}

func (impl *TagImpl) GetTag(accountId primitive.ObjectID, name string) (*dal.TagEntity, error) {
	entity, err := impl.TagRepo.GetTag(accountId, name)

	if err != nil {
		return nil, err
	}

	return entity, nil
}

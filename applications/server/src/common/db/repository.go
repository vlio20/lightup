package db

import (
	"context"
	"errors"
	logging "lightup/src/common/log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BaseEntity struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt int64              `bson:"createdAt"`
	UpdatedAt int64              `bson:"updatedAt"`
}

func GetBaseEntity() *BaseEntity {
	return &BaseEntity{
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
}

type Repository[T any] struct {
	DB         *mongo.Database
	Collection *mongo.Collection
	logger     logging.Logger
}

func (r *Repository[T]) StrIdToObjectID(id string) primitive.ObjectID {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		r.logger.Error("StrIdToObjectID is not valid, provided, id: "+id, err)
		return primitive.NilObjectID
	}

	return objectId
}

func (r *Repository[T]) Add(entity *T) (*T, error) {
	result, err := r.Collection.InsertOne(context.Background(), entity)

	if err != nil {
		return nil, err
	}

	objId := result.InsertedID.(primitive.ObjectID)

	return r.GetByObjectId(&objId)
}

// func (r *repository[T]) AddAll(entity *[]T, ctx context.Context) error {
// 	return r.db.WithContext(ctx).Create(&entity).Error
// }

func (r *Repository[T]) GetByObjectId(objectId *primitive.ObjectID) (*T, error) {
	var entity T
	err := r.Collection.FindOne(context.Background(), bson.M{"_id": objectId}).Decode(&entity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, err
	}

	return &entity, nil
}

func (r *Repository[T]) FindByStrID(id string) (*T, error) {
	objectID := r.StrIdToObjectID(id)

	if objectID == primitive.NilObjectID {
		return nil, errors.New("invalid entity id")
	}

	return r.GetByObjectId(&objectID)
}

func (r *Repository[T]) FindOne(filter bson.M) (*T, error) {
	var entity T
	filterMarshaled, err := bson.Marshal(filter)

	if err != nil {
		return nil, err
	}

	err = r.Collection.FindOne(context.Background(), filterMarshaled).Decode(&entity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, err
	}

	return &entity, nil
}

// func (r *repository[T]) GetAll(ctx context.Context) (*[]T, error) {
// 	var entities []T
// 	err := r.db.WithContext(ctx).Find(&entities).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &entities, nil
// }

// func (r *repository[T]) Where(params *T, ctx context.Context) (*[]T, error) {
// 	var entities []T
// 	err := r.db.WithContext(ctx).Where(&params).Find(&entities).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &entities, nil
// }

// func (r *repository[T]) Update(entity *T, ctx context.Context) error {
// 	return r.db.WithContext(ctx).Save(&entity).Error
// }

// func (r repository[T]) UpdateAll(entities *[]T, ctx context.Context) error {
// 	return r.db.WithContext(ctx).Save(&entities).Error
// }

// func (r *repository[T]) Delete(id int, ctx context.Context) error {
// 	var entity T
// 	return r.db.WithContext(ctx).FirstOrInit(&entity).UpdateColumn("is_active", false).Error
// }

// func (r *repository[T]) SkipTake(skip int, take int, ctx context.Context) (*[]T, error) {
// 	var entities []T
// 	err := r.db.WithContext(ctx).Offset(skip).Limit(take).Find(&entities).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &entities, nil
// }

// func (r *repository[T]) Count(ctx context.Context) int64 {
// 	var entity T
// 	var count int64
// 	r.db.WithContext(ctx).Model(&entity).Count(&count)
// 	return count
// }

// func (r *repository[T]) CountWhere(params *T, ctx context.Context) int64 {
// 	var entity T
// 	var count int64
// 	r.db.WithContext(ctx).Model(&entity).Where(&params).Count(&count)
// 	return count
// }

type IRepository[T any] interface {
	Add(entity *T) (*T, error)
	GetByObjectId(objectId *primitive.ObjectID) (*T, error)
	GetById(id string) (*T, error)
}

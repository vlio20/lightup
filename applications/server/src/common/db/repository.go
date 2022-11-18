package db

import (
	"context"
	"lightup/src/common/http"
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
}

func NewRepository[T any](db *mongo.Database, collection string) *Repository[T] {
	return &Repository[T]{
		DB:         db,
		Collection: db.Collection(collection),
	}
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
	result := r.Collection.FindOne(context.Background(), bson.M{"_id": objectId})

	err := result.Decode(&entity)

	if err != nil {
		return nil, err
	}

	return &entity, nil
}

func (r *Repository[T]) GetById(id string) (*T, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, &http.HttpError{StatusCode: 409, Message: "Invalid id", OriginalError: err}
	}

	return r.GetByObjectId(&objectId)
}

// func (r *repository[T]) Get(params *T, ctx context.Context) *T {
// 	var entity T
// 	r.db.WithContext(ctx).Where(&params).FirstOrInit(&entity)
// 	return &entity
// }

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

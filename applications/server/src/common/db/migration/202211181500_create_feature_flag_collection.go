package migration

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var CreateFeatureFlagMigration = &Migration{
	Name:      "create_feature_flag_collection",
	CreatedAt: 202211181500,
	Up: func(DB *mongo.Database) {
		CreateCollection("featureFlag", DB)

		indexModel := []mongo.IndexModel{
			{
				Keys: bson.D{{Key: "accountId", Value: 1}},
			},
			{
				Keys: bson.D{
					{Key: "accountId", Value: 1},
					{Key: "serviceId", Value: 1},
					{Key: "name", Value: -1},
				},
				Options: options.Index().SetUnique(true),
			},
			{
				Keys: bson.D{{Key: "updatedAt", Value: -1}},
			},
			{
				Keys: bson.D{{Key: "archived", Value: 1}},
			},
		}

		DB.Collection("featureFlag").Indexes().CreateMany(context.Background(), indexModel)
	},
}

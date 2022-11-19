package migration

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var CreateFeatureFlagMigration = &Migration{
	Name:      "CreateFeatureFlag",
	CreatedAt: 20221118,
	Up: func(DB *mongo.Database) {
		indexModel := []mongo.IndexModel{
			{
				Keys:    bson.D{{Key: "name", Value: -1}},
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

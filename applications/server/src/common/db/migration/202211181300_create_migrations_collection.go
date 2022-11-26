package migration

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var CreateMigrationsCollection = &Migration{
	Name:      "CreateMigrationsCollection",
	CreatedAt: 202211181300,
	Up: func(DB *mongo.Database) {
		CreateCollection("featureFlags", DB)

		indexModel := mongo.IndexModel{
			Keys:    bson.D{{Key: "name", Value: 1}},
			Options: options.Index().SetUnique(true),
		}

		if _, err := DB.Collection("migration").Indexes().CreateOne(context.Background(), indexModel); err != nil {
			panic(err)
		}
	},
}

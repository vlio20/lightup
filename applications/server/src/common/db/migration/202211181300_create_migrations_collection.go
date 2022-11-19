package migration

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var CreateMigrationsCollection = &Migration{
	Name:      "CreateMigrationsCollection",
	CreatedAt: 20221118,
	Up: func(DB *mongo.Database) {
		if err := DB.CreateCollection(context.Background(), "migration"); err != nil {
			panic(err)
		}

		indexModel := mongo.IndexModel{
			Keys:    bson.D{{Key: "name", Value: 1}},
			Options: options.Index().SetUnique(true),
		}

		if _, err := DB.Collection("migration").Indexes().CreateOne(context.Background(), indexModel); err != nil {
			panic(err)
		}
	},
}
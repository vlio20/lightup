package migration

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var CreateServiceMigration = &Migration{
	Name:      "create_service_collection",
	CreatedAt: 202211261330,
	Up: func(DB *mongo.Database) {
		CreateCollection("service", DB)

		indexModel := []mongo.IndexModel{
			{
				Keys: bson.D{{Key: "accountId", Value: 1}},
			},
			{
				Keys: bson.D{
					{Key: "accountId", Value: 1},
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

		DB.Collection("service").Indexes().CreateMany(context.Background(), indexModel)
	},
}

package migration

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var CreateUserMigration = &Migration{
	Name:      "create_user_collection",
	CreatedAt: 202211262016,
	Up: func(DB *mongo.Database) {
		CreateCollection("user", DB)

		indexModel := []mongo.IndexModel{
			{
				Keys: bson.D{{Key: "accountId", Value: 1}},
			},
			{
				Keys:    bson.D{{Key: "email", Value: 1}},
				Options: options.Index().SetUnique(true),
			},
			{
				Keys: bson.D{{Key: "updatedAt", Value: -1}},
			},
			{
				Keys: bson.D{{Key: "archived", Value: 1}},
			},
		}

		DB.Collection("user").Indexes().CreateMany(context.Background(), indexModel)
	},
}

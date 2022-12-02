package migration

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var CreateAccountMigration = &Migration{
	Name:      "create_account_collection",
	CreatedAt: 202211261536,
	Up: func(DB *mongo.Database) {
		CreateCollection("account", DB)

		indexModel := []mongo.IndexModel{
			{
				Keys: bson.D{{Key: "updatedAt", Value: -1}},
			},
			{
				Keys: bson.D{{Key: "archived", Value: 1}},
			},
		}

		DB.Collection("account").Indexes().CreateMany(context.Background(), indexModel)
	},
}

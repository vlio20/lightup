package migration

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateCollection(migrationName string, db *mongo.Database) {
	cols, err := db.ListCollections(context.Background(), bson.M{})

	if err != nil {
		panic(err)
	}

	var exists = false

	for cols.Next(context.Background()) {
		var col bson.M
		if err := cols.Decode(&col); err != nil {
			panic(err)
		}

		if col["name"] == migrationName {
			exists = true
		}
	}

	if !exists {
		if err := db.CreateCollection(context.Background(), migrationName); err != nil {
			panic(err)
		}
	}

}

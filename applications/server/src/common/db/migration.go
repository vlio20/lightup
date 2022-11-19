package db

import (
	"context"
	"fmt"
	"lightup/src/common/db/migration"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var getDb = func() *mongo.Database {
	return client.Database("lightup")
}
var migrationsList = []migration.Migration{
	*migration.CreateMigrationsCollection,
	*migration.CreateFeatureFlagMigration,
}

func RunMigrations() {
	col := getDb().Collection("migration")

	for _, migration := range migrationsList {
		if !checkIfMigrationExist(col, migration.Name) {
			fmt.Println("Creating migration: ", migration.Name)
			migration.Up(getDb())
			col.InsertOne(context.Background(), bson.M{"name": migration.Name, "createdAt": migration.CreatedAt})
		} else {
			fmt.Println("Migration already exists: ", migration.Name)
		}
	}
}

func checkIfMigrationExist(col *mongo.Collection, migtationName string) bool {
	result, err := col.CountDocuments(context.Background(), bson.M{"name": migtationName})

	if err != nil {
		panic(err)
	}

	return result > 0
}

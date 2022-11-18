package db

import (
	"context"
	"fmt"
	"lightup/src/common/db/migrations"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var getDb = func() *mongo.Database {
	return client.Database("lightup")
}
var migrationsList []migrations.Migration

func RunMigrations() {
	col := getDb().Collection("migration")
	registerMigration()

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

func registerMigration() {
	migrationsList = append(migrationsList,
		*migrations.CreateMigrationsCollection,
		*migrations.CreateFeatureFlagMigration,
	)
}

func checkIfMigrationExist(col *mongo.Collection, migtationName string) bool {
	result, err := col.CountDocuments(context.Background(), bson.M{"name": migtationName})

	if err != nil {
		panic(err)
	}

	return result > 0
}

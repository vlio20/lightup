package db

import (
	"context"
	"lightup/src/common/db/migration"
	"lightup/src/common/log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var getDb = func() *mongo.Database {
	return client.Database("lightup")
}

var migrationsList = []migration.Migration{
	*migration.CreateMigrationsCollection,
	*migration.CreateFeatureFlagMigration,
	*migration.CreateServiceMigration,
	*migration.CreateAccountMigration,
}

func RunMigrations() {
	logger := log.GetLogger("migration")
	col := getDb().Collection("migration")

	for _, migration := range migrationsList {
		if !checkIfMigrationExist(col, migration.Name) {
			logger.Info("Creating migration: " + migration.Name)
			migration.Up(getDb())
			col.InsertOne(context.Background(), bson.M{"name": migration.Name, "createdAt": migration.CreatedAt})
		} else {
			logger.Info("Migration already exists: ", migration.Name)
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

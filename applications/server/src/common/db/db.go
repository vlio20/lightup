package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client

func InitDB() {
	const uri = "mongodb://root:root@localhost:27017/?maxPoolSize=20&w=majority"

	c, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	client = c

	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected and pinged.")
	RunMigrations()
}

func GetDBClient() *mongo.Client {
	return client
}

func GetDB(name string) *mongo.Database {
	return client.Database(name)
}

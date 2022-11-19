package db

import (
	"context"
	"fmt"
	"lightup/src/common/config"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

var client *mongo.Client
var getConfig = config.UnmarshalKey

type dbConfig struct {
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	User        string `mapstructure:"username"`
	Pass        string `mapstructure:"password"`
	MaxPoolSize uint64 `mapstructure:"maxPoolSize"`
}

func Init() {
	dbConfig := &dbConfig{}
	err := getConfig("db", dbConfig)

	if err != nil {
		panic(err)
	}

	clientOptions := options.ClientOptions{
		Hosts:        []string{dbConfig.Host + ":" + strconv.Itoa(dbConfig.Port)},
		Auth:         &options.Credential{Username: dbConfig.User, Password: dbConfig.Pass},
		MaxPoolSize:  &dbConfig.MaxPoolSize,
		WriteConcern: writeconcern.New(writeconcern.WMajority()),
	}

	c, err := mongo.NewClient(&clientOptions)
	if err != nil {
		panic(err)
	}

	c.Connect(context.Background())

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

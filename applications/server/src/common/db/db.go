package db

import (
	"context"
	"lightup/src/common/config"
	"lightup/src/common/log"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

type dbConfig struct {
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	User        string `mapstructure:"username"`
	Pass        string `mapstructure:"password"`
	MaxPoolSize uint64 `mapstructure:"maxPoolSize"`
	DBName      string `mapstructure:"dbName"`
}

var client *mongo.Client
var getConfig = config.UnmarshalKey
var dbConf = &dbConfig{}

func Init() {
	getConfig("db", dbConf)
	logger := log.GetLogger("db")

	clientOptions := options.ClientOptions{
		Hosts:        []string{dbConf.Host + ":" + strconv.Itoa(dbConf.Port)},
		Auth:         &options.Credential{Username: dbConf.User, Password: dbConf.Pass},
		MaxPoolSize:  &dbConf.MaxPoolSize,
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

	logger.Info("Successfully connected and pinged.")
	RunMigrations()
}

func GetDBClient() *mongo.Client {
	return client
}

func GetDB() *mongo.Database {
	return client.Database(dbConf.DBName)
}

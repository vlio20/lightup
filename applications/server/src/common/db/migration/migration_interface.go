package migration

import "go.mongodb.org/mongo-driver/mongo"

type Migration struct {
	Name      string
	CreatedAt uint64
	Up        func(DB *mongo.Database)
}

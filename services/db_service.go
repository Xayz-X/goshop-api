package services

import (
	"context"
	"log"

	"github.com/Xayz-X/goshop-api/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoCollections struct {
	UserCol *mongo.Collection
}

func ConnectToDb() (*mongo.Client, *mongo.Database) {
	config.LoadEnv()
	mongoURI := config.GetEnv("MONGO_URI")
	databaseName := config.GetEnv("DATABASE_NAME")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("mongo client error", err)
	}

	database := client.Database(databaseName)
	return client, database
}

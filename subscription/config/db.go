package config

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))

	if err != nil {
		log.Fatal(err)
	}
	return client

}

var DB *mongo.Client = ConnectDB()

func GetCollection(client *mongo.Client, collectionname string) *mongo.Collection {
	return client.Database(os.Getenv("DBNAME")).Collection(collectionname)
}

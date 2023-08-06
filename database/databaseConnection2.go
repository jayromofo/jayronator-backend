package database

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

func GetClient() *mongo.Client {
	err := godotenv.Load("./env/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mongoString := os.Getenv("MONGO_STRING")
	MongoDB := "mongodb://localhost:27017"
	println(mongoString)

	fmt.Println(MongoDB + " | " + time.Now().String())

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoString))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	println("Connected to Mongo DB")

	return client
}

var Client *mongo.Client = GetClient()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("cookbook").Collection(collectionName)

	return collection
}

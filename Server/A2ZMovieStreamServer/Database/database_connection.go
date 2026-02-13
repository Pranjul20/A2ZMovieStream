package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func Connect() *mongo.Client {

	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Error loading .env file")
	}

	MongoDb := os.Getenv("MONGODB_URI")

	if MongoDb == "" {
		log.Fatal("MONGODB_URI not found in environment variables")
	}

	fmt.Println("MongoDB URI:", MongoDb)

	clientOptions := options.Client().ApplyURI(MongoDb)

	client, err := mongo.Connect(clientOptions)

	if err != nil {
		return nil
	}
	return client
}

var Client *mongo.Client = Connect()

func OpenCollection(CollectionName string, client *mongo.Client) *mongo.Collection {

	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Error loading .env file")
	}

	DatabaseName := os.Getenv("DATABASE_NAME")

	fmt.Println("DATABASE_NAME:", DatabaseName)

	collection := Client.Database(DatabaseName).Collection(CollectionName)

	if collection == nil {
		return nil
	}

	return collection
}

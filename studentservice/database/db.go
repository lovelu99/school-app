package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var Database *mongo.Database

func Connect() error {
	// From Environment variables 
	MONGO_HOST := os.Getenv("MONGO_HOST")
	if MONGO_HOST == "" {
		MONGO_HOST = "localhost"
	}
	
	MONGO_PORT := os.Getenv("MONGO_PORT")
	if MONGO_PORT == "" {
		MONGO_PORT = "27017"
	}
	
	MONGO_DB := os.Getenv("MONGO_DB")
	if MONGO_DB == "" {
		MONGO_DB = "kindergarten"
	}
	
	MONGO_USER := os.Getenv("MONGO_USER")
	if MONGO_USER == "" {
		MONGO_USER = "admin"
	}
	
	MONGO_PASSWORD := os.Getenv("MONGO_PASSWORD")
	if MONGO_PASSWORD == "" {
		MONGO_PASSWORD = "password123"
	}
	

	MONGO_AUTH_SOURCE := os.Getenv("MONGO_AUTH_SOURCE")
	if MONGO_AUTH_SOURCE == "" {
		MONGO_AUTH_SOURCE = "admin"
	}

   // mongodb_uri=f"mongodb://{MONGO_USER}:{MONGO_PASSWORD}@{MONGO_HOST}:{MONGO_PORT}/{MONGO_DB}?authSource={MONGO_AUTH_SOURCE}"
	mongodbURI := "mongodb://" + MONGO_USER + ":" + MONGO_PASSWORD + "@" + MONGO_HOST + ":" + MONGO_PORT + "/" + MONGO_DB + "?authSource=" + MONGO_AUTH_SOURCE
	connectionString := mongodbURI
	//connectionString := os.Getenv("MONGODB_URI")
	// if connectionString == "" {
	// 	connectionString = "mongodb://mongo:27017" // default value
	// }

	// databaseName := os.Getenv("MONGO_DB")
	// if databaseName == "" {
	// 	databaseName = "kindergarten" // default value
	// }

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return err
	}

	Client = client
	Database = client.Database("kindergarten")
	log.Println("Connected to MongoDB successfully!")
	return nil
}

func GetCollection(collectionName string) *mongo.Collection {
	return Database.Collection(collectionName)
}
package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDB() *mongo.Database {
	// Set client options
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@%s", os.Getenv("MONGODB_USERNAME"), os.Getenv("MONGODB_PASSWORD"), os.Getenv("MONGODB_HOST")))

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the primary to verify that the client can connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client.Database(os.Getenv("MONGODB_DATABASE"))
}

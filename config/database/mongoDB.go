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

	// // specify the collection and field name to rename
	// collection := client.Database("location_db").Collection("countries")
	// oldFieldName := "iso2"
	// newFieldName := "cca2"

	// // update all documents in the collection with the renamed field
	// filter := bson.M{}
	// update := bson.M{"$rename": bson.M{oldFieldName: newFieldName}}
	// result, err := collection.UpdateMany(ctx, filter, update)
	// if err != nil {
	// 	fmt.Println("Error updating documents in MongoDB:", err)
	// }

	// fmt.Printf("Modified %v documents\n", result.ModifiedCount)

	return client.Database(os.Getenv("MONGODB_DATABASE"))
}

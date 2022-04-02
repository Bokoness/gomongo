package db

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connect() (*mongo.Client, context.Context) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO"))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err) //TODO: change to normal error
	}
	return client, ctx
}

func InsertOne(col string, data interface{}) (*mongo.InsertOneResult, error) {
	var dbName = os.Getenv("DB")
	client, ctx := connect()
	db := client.Database(dbName)
	defer client.Disconnect(ctx)
	collection := db.Collection(col)
	res, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func FindOne(col string, filter map[string]string) *mongo.SingleResult {
	var dbName = os.Getenv("DB")
	client, ctx := connect()
	db := client.Database(dbName)
	defer client.Disconnect(ctx)
	collection := db.Collection(col)
	result := collection.FindOne(ctx, filter)
	return result
}

func UpdateOne(col string, filter map[string]string, update interface{}) error {
	var dbName = os.Getenv("DB")
	client, ctx := connect()
	db := client.Database(dbName)
	defer client.Disconnect(ctx)
	collection := db.Collection(col)
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

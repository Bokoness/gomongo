package db

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connect() (*mongo.Client, *mongo.Database, context.Context) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO"))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err) //TODO: change to normal error
	}
	var dbName = os.Getenv("DB")
	db := client.Database(dbName)
	return client, db, ctx
}

func InsertOne(col string, data interface{}) (*mongo.InsertOneResult, error) {
	client, db, ctx := connect()
	defer client.Disconnect(ctx)
	collection := db.Collection(col)
	res, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func FindOne(col string, filter map[string]string) *mongo.SingleResult {
	client, db, ctx := connect()
	defer client.Disconnect(ctx)
	collection := db.Collection(col)
	result := collection.FindOne(ctx, filter)
	return result
}

func FindMany(col string, filter interface{}) (*mongo.Cursor, error) {
	client, db, ctx := connect()
	defer client.Disconnect(ctx)
	collection := db.Collection(col)
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	return cursor, nil
}

func FindOneById(col string, id string) *mongo.SingleResult {
	objectId, _ := primitive.ObjectIDFromHex(id)
	client, db, ctx := connect()
	defer client.Disconnect(ctx)
	collection := db.Collection(col)
	result := collection.FindOne(ctx, bson.D{{Key: "_id", Value: objectId}})
	return result
}

func FindByIdAndUpdate(col string, id string, updates interface{}) (*mongo.UpdateResult, error) {
	objectId, _ := primitive.ObjectIDFromHex(id)
	client, db, ctx := connect()
	defer client.Disconnect(ctx)
	collection := db.Collection(col)
	result, err := collection.UpdateByID(ctx, objectId, bson.D{{Key: "$set", Value: updates}})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func UpdateOne(col string, filter map[string]string, updates interface{}) (*mongo.UpdateResult, error) {
	client, db, ctx := connect()
	defer client.Disconnect(ctx)
	collection := db.Collection(col)
	result, err := collection.UpdateOne(ctx, filter, updates)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func FindByIdAndDelete(col string, id string) error {
	client, db, ctx := connect()
	defer client.Disconnect(ctx)
	collection := db.Collection(col)
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectId}}
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

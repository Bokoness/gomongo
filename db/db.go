package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	m := map[string]string{
		"username": "bokoness_updated",
	}
	fmt.Println(m)
	mm := map[string]interface{}{
		"$set": m,
	}
	var dbName = os.Getenv("DB")
	client, ctx := connect()
	db := client.Database(dbName)
	defer client.Disconnect(ctx)
	collection := db.Collection(col)
	_, err := collection.UpdateOne(context.TODO(), filter, mm)
	if err != nil {
		return err
	}
	return nil
}

func FindByIdAndUpdate(col string, id string, update interface{}) error {
	_id, _ := primitive.ObjectIDFromHex(id)
	// filter := map[string]primitive.ObjectID{
	// 	"_id": _id,
	// }
	filter := bson.D{{"_id", _id}}
	var dbName = os.Getenv("DB")
	client, ctx := connect()
	db := client.Database(dbName)
	defer client.Disconnect(ctx)
	collection := db.Collection(col)
	tempUpdate := bson.D{
		{"$set", bson.D{
			{"username", "yael"},
		}},
	}
	_, err := collection.UpdateOne(ctx, filter, tempUpdate)
	if err != nil {
		return err
	}
	return nil
}

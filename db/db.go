package db

import (
	"context"
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connect() (*mongo.Client, *mongo.Database, context.Context, error) {
	ctx := context.Background()
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO"))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		err := errors.New("cannot connect to database")
		return nil, nil, nil, err
	}
	var dbName = os.Getenv("DB")
	db := client.Database(dbName)
	return client, db, ctx, nil
}

func InsertOne(col string, data interface{}) (*mongo.InsertOneResult, error) {
	client, db, ctx, err := connect()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(ctx)
	collection := db.Collection(col)
	res, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func FindOne(col string, filter map[string]string) (*mongo.SingleResult, error) {
	client, db, ctx, err := connect()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(ctx)
	collection := db.Collection(col)
	result := collection.FindOne(ctx, filter)
	return result, err
}

func FindMany(col string, filter interface{}) (*mongo.Cursor, error) {
	client, db, ctx, err := connect()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(ctx)
	collection := db.Collection(col)
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	return cursor, nil
}

func FindOneById(col string, id string) (*mongo.SingleResult, error) {
	objectId, _ := primitive.ObjectIDFromHex(id)
	client, db, ctx, err := connect()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(ctx)
	collection := db.Collection(col)
	result := collection.FindOne(ctx, bson.D{{Key: "_id", Value: objectId}})
	return result, nil
}

func FindByIdAndUpdate(col string, id primitive.ObjectID, updates interface{}) (*mongo.UpdateResult, error) {
	client, db, ctx, err := connect()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(ctx)
	collection := db.Collection(col)
	result, err := collection.UpdateByID(ctx, id, bson.D{{Key: "$set", Value: updates}})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func UpdateOne(col string, filter map[string]string, updates interface{}) (*mongo.UpdateResult, error) {
	client, db, ctx, err := connect()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(ctx)
	collection := db.Collection(col)
	result, err := collection.UpdateOne(ctx, filter, updates)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func FindByIdAndDelete(col string, id primitive.ObjectID) error {
	client, db, ctx, err := connect()
	if err != nil {
		return err
	}
	defer client.Disconnect(ctx)
	collection := db.Collection(col)
	filter := bson.D{{Key: "_id", Value: id}}
	_, err = collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

// InsertedIdToObjectId convert Result.InsertedId into primitive.ObjectId
func InsertedIdToObjectId(id interface{}) (*primitive.ObjectID, error) {
	if oid, ok := id.(primitive.ObjectID); ok {
		return &oid, nil
	} else {
		return nil, errors.New("cannot convert into objectId")
	}
}

func InsertedIdToString(id interface{}) (string, error) {
	if oid, ok := id.(string); ok {
		return oid, nil
	} else {
		return "", errors.New("cannot convert into objectId")
	}
}

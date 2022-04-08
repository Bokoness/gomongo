package db_test

import (
	"gomongo/db"
	"testing"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TestModel struct {
	ID        string `json:"_id" bson:"_id"`
	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName"`
}

var testModel TestModel

const myid = "624ff033ec3ff69bf51cbbf8"

func (t TestModel) toMap(showId bool) map[string]string {
	mapData := map[string]string{
		"firstName": t.FirstName,
		"lastName":  t.LastName,
	}
	if showId {
		mapData["_id"] = t.ID
	}
	return mapData
}

func TestInsertOne(t *testing.T) {
	loadEnv()
	testModel.ID = myid
	testModel.FirstName = "Ness"
	testModel.LastName = "Bokobza"
	result, err := db.InsertOne("testing", testModel.toMap(true))
	if err != nil {
		t.Error(err)
	}
	resultStringId, err := db.InsertedIdToString(result.InsertedID)
	if err != nil {
		t.Error(err)
	}
	if resultStringId != testModel.ID {
		t.Error("result id not equals to inserted id")
	}
}

func TestFindByIdAndDelete(t *testing.T) {
	loadEnv()
	id, err := primitive.ObjectIDFromHex(myid)
	if err != nil {
		t.Error(err)
	}
	err = db.FindByIdAndDelete("testing", id)
	if err != nil {
		t.Error(err)
	}
}

func loadEnv() {
	godotenv.Load("../.env")
}

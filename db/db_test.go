package db_test

import (
	"gomongo/db"
	"testing"

	"github.com/joho/godotenv"
)

type TestModel struct {
	ID        string `json:"_id" bson:"_id"`
	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName"`
}

var testModel TestModel

const myid = "624ff033ec3ff69bf51cbbf8"
const dbname = "testing"

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
	result, err := db.InsertOne(dbname, testModel.toMap(true))
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

func TestFindByIdAndUpdate(t *testing.T) {
	loadEnv()
	testModel.ID = myid
	testModel.FirstName = "name updated"
	testModel.LastName = "lastname updated"
	result, err := db.FindByIdAndUpdate(dbname, myid, testModel.toMap(true))
	if err != nil {
		t.Error(err)
	}
	if result.MatchedCount <= 0 {
		t.Error("couldn't find record to udpate")
	} else if result.ModifiedCount <= 0 {
		t.Error("document is found but did not updated")
	}
}

func TestFindByIdAndDelete(t *testing.T) {
	loadEnv()
	result, err := db.FindByIdAndDelete(dbname, myid)
	if err != nil {
		t.Error(err)
	}
	if result.DeletedCount <= 0 {
		t.Error("item is not deleted")
	}
}

func TestDropDb(t *testing.T) {
	loadEnv()
	err := db.DropDatabase(dbname)
	if err != nil {
		t.Error(err)
	}
}

func loadEnv() {
	godotenv.Load("../.env")
}

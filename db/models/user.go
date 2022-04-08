package models

import (
	"context"
	"errors"
	"gomongo/db"
	"gomongo/services"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const collection = "users"

type User struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"`
	Todos    []Todo             `json:"todos" bson:"todos"`
}

func (u User) AsMap(showId, showPass bool) map[string]string {
	uData := map[string]string{
		"username": u.Username,
	}
	if showPass {
		uData["password"] = u.Password
	}
	if showId {
		uData["_id"] = u.ID.Hex()
	}
	return uData
}

func (u User) AsMapNoPwd() map[string]string {
	uData := map[string]string{
		"username": u.Username,
	}
	return uData
}

func (u *User) FindByIdAndUpdate(id primitive.ObjectID) {
	update := map[string]string{
		"username": "bokoness is the update!!!!!",
	}
	db.FindByIdAndUpdate(collection, id, update)
}

func (u *User) FindById(id string) error {
	result, err := db.FindOneById(collection, id)
	if err != nil {
		return err
	}
	err = result.Decode(u)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) FindByUsername(username string) error {
	filter := map[string]string{
		"username": username,
	}
	result, err := db.FindOne(collection, filter)
	if err != nil {
		return err
	}
	result.Decode(u)
	return nil
}

func (u *User) Store() error {
	u.Password = services.Hash(u.Password)
	uData := u.AsMap(false, true)
	result, err := db.InsertOne(collection, uData)
	if err != nil {
		return err
	}
	id, err := db.InsertedIdToObjectId(result.InsertedID)
	if err != nil {
		return err
	}
	u.ID = *id
	return nil
}

func (u *User) Save() error {
	if u.ID.IsZero() {
		return errors.New("user id is required")
	}
	_, err := db.FindByIdAndUpdate(collection, u.ID, u)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Destroy() error {
	err := db.FindByIdAndDelete(collection, u.ID)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) LoadTodos() error {
	if u.ID.IsZero() {
		return errors.New("user id is required")
	}
	filter := map[string]primitive.ObjectID{"user": u.ID}
	cursor, err := db.FindMany("todos", filter)
	if err != nil {
		return err
	}
	for cursor.Next(context.Background()) {
		var t Todo
		if err := cursor.Decode(&t); err != nil {
			return err
		}
		u.Todos = append(u.Todos, t)
	}
	return nil
}

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
	ID       string `json:"_id" bson:"_id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Todos    []Todo `json:"todos" bson:"todos"`
}

func (u User) AsMap() map[string]string {
	uData := map[string]string{
		"username": u.Username,
		"password": u.Password,
	}
	return uData
}

func (u User) AsMapNoPwd() map[string]string {
	uData := map[string]string{
		"username": u.Username,
	}
	return uData
}

func (u *User) FindByIdAndUpdate(id string) {
	update := map[string]string{
		"username": "bokoness is the update!!!!!",
	}
	db.FindByIdAndUpdate(collection, id, update)
}

func (u *User) FindById(id string) error {
	result := db.FindOneById(collection, id)
	err := result.Decode(u)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) FindByUsername(username string) {
	filter := map[string]string{
		"username": username,
	}
	result := db.FindOne(collection, filter)
	result.Decode(u)
}

func (u *User) Store() {
	u.Password = services.Hash(u.Password)
	uData := u.AsMap()
	db.InsertOne(collection, uData)
}

func (u *User) Save() error {
	if u.ID == "" {
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
	if u.ID == "" {
		return errors.New("user id is required")
	}
	userId, _ := primitive.ObjectIDFromHex(u.ID)
	filter := map[string]primitive.ObjectID{"user": userId}
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

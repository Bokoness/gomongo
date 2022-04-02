package models

import (
	"gomongo/db"
	"gomongo/services"
)

const collection = "users"

type User struct {
	Id       int64  `json:"id"`
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password"`
	Todos    []Todo
}

func (u *User) FindByIdAndUpdate(id string) {
	update := map[string]string{
		"username": "bokoness is the update",
	}
	db.FindByIdAndUpdate(collection, id, update)

}

func (u *User) FindById(id int64) {

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
	db.InsertOne(collection, u)
}

func (u *User) Save() {

}

func (u *User) Destroy() {

}

func (u *User) LoadTodos() {

}

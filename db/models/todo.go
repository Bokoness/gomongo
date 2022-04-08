package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	ID      primitive.ObjectID `json:"_id" bson:"_id"`
	Title   string             `json:"title" bson:"totle"`
	Content string             `json:"content" bson:"content"`
	User    primitive.ObjectID `json:"user" bson:"user"`
}

func (t *Todo) FindById(id int64) {

}

func (t *Todo) Create() {

}

func (t *Todo) Save() {

}

func (t *Todo) Destroy() {

}

func (t Todo) Index() []Todo {
	var tt []Todo
	return tt
}

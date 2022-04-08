package models

type Todo struct {
	ID      string `json:"_id" bson:"_id"`
	Title   string `json:"title" bson:"totle"`
	Content string `json:"content" bson:"content"`
	User    string `json:"user" bson:"user"`
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

package models

type Todo struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  int64  `json:"userId"`
	User    User
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

package todo

type TodoList struct {
	Id          int    `json:"-" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type TodoItem struct {
	Id          int    `json:"-"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}

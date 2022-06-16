package todo

import "errors"

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

type UpdateTodoListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (i UpdateTodoListInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update structure hasn't values")
	}
	return nil
}

package repository

import (
	"fmt"
	"strings"

	"github.com/Includeoyasi/todo-app"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (r *TodoItemPostgres) Create(listId int, item todo.TodoItem) (int, error) {
	tr, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int
	query := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoItemsTable)
	row := tr.QueryRow(query, item.Title, item.Description)
	if err := row.Scan(&itemId); err != nil {
		tr.Rollback()
		return 0, err
	}

	createListItemsQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES ($1, $2) RETURNING id", listsItemsTable)
	if _, err := tr.Exec(createListItemsQuery, listId, itemId); err != nil {
		tr.Rollback()
		return 0, err
	}

	return itemId, tr.Commit()
}

func (r *TodoItemPostgres) GetAll(userId, listId int) ([]todo.TodoItem, error) {
	var items []todo.TodoItem

	query := fmt.Sprintf(`SELECT ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li ON li.item_id = ti.id 
							INNER JOIN %s ul ON li.list_id = ul.list_id WHERE li.list_id = $1 AND ul.user_id = $2`,
		todoItemsTable, listsItemsTable, usersListsTable)
	if err := r.db.Select(&items, query, listId, userId); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *TodoItemPostgres) GetById(userId, itemId int) (todo.TodoItem, error) {
	var item todo.TodoItem

	query := fmt.Sprintf(`SELECT ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li ON li.item_id = ti.id 
							INNER JOIN %s ul ON li.list_id = ul.list_id WHERE ti.id = $1 AND ul.user_id = $2`,
		todoItemsTable, listsItemsTable, usersListsTable)
	if err := r.db.Get(&item, query, itemId, userId); err != nil {
		return item, err
	}

	return item, nil
}

func (r *TodoItemPostgres) Delete(userId, itemId int) error {
	query := fmt.Sprintf(`DELETE FROM %s ti USING %s li, %s ul 
							WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ti.id = $1 AND ul.user_id = $2`,
		todoItemsTable, listsItemsTable, usersListsTable)
	_, err := r.db.Exec(query, userId, itemId)

	return err
}

func (r *TodoItemPostgres) Update(userId, itemId int, input todo.UpdateTodoItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argKey := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argKey))
		args = append(args, *input.Title)
		argKey++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argKey))
		args = append(args, *input.Description)
		argKey++
	}

	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argKey))
		args = append(args, *input.Done)
		argKey++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf(`UPDATE %s ti SET %s FROM %s li, %s ul 
							WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ti.id = $%d AND ul.user_id = $%d`,
		todoItemsTable, setQuery, listsItemsTable, usersListsTable, argKey, argKey+1)
	args = append(args, itemId, userId)
	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)

	return err
}

package repository

import (
	"fmt"
	"github.com/AlibekDalgat/todo-app"
	"github.com/jmoiron/sqlx"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db}
}

func (itemPostgres *TodoItemPostgres) Create(listId int, item todo.TodoItem) (int, error) {
	tx, err := itemPostgres.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int

	createQuery := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", todoItemsTable)
	row := tx.QueryRow(createQuery, item.Title, item.Description)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createListItemsQuery := fmt.Sprintf("INSERT INTO %s (item_id, list_id) values ($1, $2) RETURNING id", todoItemsTable)
	_, err = tx.Exec(createListItemsQuery, itemId, listId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

func (itemPostgres *TodoItemPostgres) GetAll(userId, listId int) ([]todo.TodoItem, error) {
	var items []todo.TodoItem
	query := fmt.Sprintf("SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti "+
		"INNER JOIN %s li on li.item_id = ti.id INNER JOIN %s ul on ul.list_id = li.list_id WHERE li.list_id = &1 AND ul.user_id = $2",
		todoItemsTable, listsItemsTable, usersListsTable)
	if err := itemPostgres.db.Select(&items, query, listId, userId); err != nil {
		return nil, err
	}
	return items, nil
}

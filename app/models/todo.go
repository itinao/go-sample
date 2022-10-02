package models

import (
	"fmt"
)

type Todo struct {
	Id   int
	Todo string
}

func NewTodo(id int, todo string) *Todo {
	return &Todo{
		id,
		todo,
	}
}

func (t *Todo) TableName() string {
	return tableNameTodo
}

func (t *Todo) Create() error {
	cmd := fmt.Sprintf("INSERT INTO %s (id, todo) VALUES (?, ?)", t.TableName())
	_, err := DbConnection.Exec(cmd, t.Id, t.Todo)
	if err != nil {
		return err
	}
	return err
}

func (t *Todo) Save() error {
	cmd := fmt.Sprintf("UPDATE %s SET todo = ? WHERE id = ?", t.TableName())
	_, err := DbConnection.Exec(cmd, t.Todo, t.Id)
	if err != nil {
		return err
	}
	return err
}

func GetTodo(id int) *Todo {
	tableName := tableNameTodo
	cmd := fmt.Sprintf("SELECT id, todo FROM  %s WHERE id = ?", tableName)
	row := DbConnection.QueryRow(cmd, id)

	var todo Todo
	err := row.Scan(&todo.Id, &todo.Todo)
	if err != nil {
		return nil
	}
	return NewTodo(id, todo.Todo)
}

func GetAllTodo(limit int) (dfTodo *DataFrameTodo, err error) {
	tableName := tableNameTodo
	cmd := fmt.Sprintf(`
		SELECT id, todo FROM %s LIMIT ?`, tableName)
	rows, err := DbConnection.Query(cmd, limit)
	if err != nil {
		return
	}
	defer rows.Close()

	dfTodo = &DataFrameTodo{}
	for rows.Next() {
		var todo Todo
		rows.Scan(&todo.Id, &todo.Todo)
		dfTodo.Todos = append(dfTodo.Todos, todo)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return dfTodo, nil
}

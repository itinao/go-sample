package models

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/itinao/go-sample/config"
	_ "github.com/mattn/go-sqlite3"
)

const (
	tableNameTodo = "todo_list"
)

var DbConnection *sql.DB

func init() {
	var err error

	DbConnection, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	cmd := fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %s (
            id INT PRIMARY KEY NOT NULL,
            todo STRING)`, tableNameTodo)
	_, err = DbConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}
}

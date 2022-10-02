package controllers

import (
	"github.com/itinao/go-sample/app/models"
)

func IngestionData() {
	minId := 100
	m := []string{
		"hogehoge1",
		"hogehoge2",
		"hogehoge3",
		"hogehoge4",
		"hogehoge5",
		"hogehoge6",
		"hogehoge7",
		"hogehoge8",
		"hogehoge9",
	}

	for i, v := range m {
		models.NewTodo(minId+i, v).Create()
	}
}

package main

import (
	"todo/internal/command"
	"todo/internal/storage"
	"todo/internal/todo"
)

func main() {
	todos := todo.Todos{}
	storage := storage.NewStorage[todo.Todos]("data/data.json")

	storage.Load(&todos)

	commandFlags := command.NewCommandFlags()

	commandFlags.Execute(&todos)

	storage.Save(todos)
}

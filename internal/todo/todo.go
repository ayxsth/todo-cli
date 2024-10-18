package todo

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	IsCompleted bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) Add(title string) {
	todo := Todo{Title: title, IsCompleted: false, CreatedAt: time.Now(), CompletedAt: nil}

	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("invalid index")

		fmt.Println(err)

		return err
	}

	return nil
}

func (todos *Todos) Delete(index int) error {
	if err := todos.validateIndex(index); err != nil {
		return err
	}

	*todos = append((*todos)[:index], (*todos)[index+1:]...)

	return nil
}

func (todos *Todos) Toggle(index int) error {
	if err := todos.validateIndex(index); err != nil {
		return err
	}

	todo := &(*todos)[index]

	todo.IsCompleted = !todo.IsCompleted

	if todo.IsCompleted {
		completedAt := time.Now()

		todo.CompletedAt = &completedAt
	} else {
		todo.CompletedAt = nil
	}

	return nil
}

func (todos *Todos) Update(index int, title string) error {
	if err := todos.validateIndex(index); err != nil {
		return err
	}

	(*todos)[index].Title = title

	return nil
}

func (todos *Todos) List() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Is Completed", "Created At", "Completed At")

	for i, todo := range *todos {
		isCompleted := "No"
		completedAt := ""

		if todo.IsCompleted {
			isCompleted = "Yes"
			completedAt = todo.CompletedAt.Format(time.RFC3339)
		}

		table.AddRow(strconv.Itoa(i+1), todo.Title, isCompleted, todo.CreatedAt.Format(time.RFC3339), completedAt)
	}

	table.Render()
}

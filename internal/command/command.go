package command

import (
	"flag"
	"strconv"
	"strings"

	"todo/internal/todo"
)

type CommandFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCommandFlags() *CommandFlags {
	cf := CommandFlags{}

	// Take in command line arguments
	flag.StringVar(&cf.Add, "add", "", "Add a new task")
	flag.IntVar(&cf.Del, "del", -1, "Delete a task")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a task")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle a task")
	flag.BoolVar(&cf.List, "list", false, "List all tasks")

	// Parse the flags
	flag.Parse()

	return &cf
}

func (cf *CommandFlags) Execute(todos *todo.Todos) {
	switch {
	case cf.List:
		todos.List()
	case cf.Add != "":
		todos.Add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)

		if len(parts) != 2 {
			panic("Invalid edit format, should be <index>:<new task>")
		}

		index, err := strconv.Atoi(parts[0])

		if err != nil {
			panic("Invalid edit format, should be <index>:<new task>")
		}

		todos.Update(index-1, parts[1])
	case cf.Toggle != -1:
		todos.Toggle(cf.Toggle - 1)
	case cf.Del != -1:
		todos.Delete(cf.Del - 1)
	default:
		panic("No command specified")
	}
}

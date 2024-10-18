# Todo CLI Application

This is a simple command-line interface (CLI) application for managing a todo list. The application allows you to add, delete, edit, toggle, and list tasks.

## Project Structure

cmd/ app/ main.go go.mod go.sum internal/ command/ command.go storage/ storage.go todo/ todo.go todos.json

```
todo-cli/
├── cmd/
│   └── app/
│       └── main.go
├── internal/
│   ├── command/
│   │   └── command.go
│   ├── storage/
│   │   └── storage.go
│   └── todo/
│       ├── todo.go
│       └── todos.json
├── go.mod
└── go.sum
```

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/hello/todo-cli.git
   cd todo-cli
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```

## Usage

Run the application:

```sh
go run cmd/app/main.go
```

## Command Line Flags

- `-add`: Add a new task
- `-del`: Delete a task by index
- `-edit`: Edit a task by index
- `-toggle`: Toggle the completion status of a task by index
- `-list`: List all tasks

## Examples

### Add a new task:

```sh
go run cmd/app/main.go -add "Buy groceries"
```

### Delete a task by index:

```sh
go run cmd/app/main.go -del 1
```

### Edit a task by index:

```sh
go run cmd/app/main.go -edit 1 "Buy groceries and eggs"
```

### Toggle the completion status of a task by index:

```sh
go run cmd/app/main.go -toggle 1
```

### List all tasks:

```sh
go run cmd/app/main.go -list
```

## License

This project is licensed under the MIT License.

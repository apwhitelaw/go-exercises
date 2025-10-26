# To-do List

 A simple command line to-do list application built using only Go's standard library. Demonstrates key Go features including: 
   * File I/O with JSON
   * Command line arguments
   * Struct and slice manipulation
   * Conventional Go error handling

### Usage

Manage your to-do list from the command line.
All data stored locally in `tasks.json` in the current directory, while will be created automatically if it does not exist.

| Command          | Description                                           | Example                          |
| ---------------- | ----------------------------------------------------- | -------------------------------- |
| `list`           | Display all current tasks                             | `go run todo.go list`                |
| `add <task>`     | Add a new task                                        | `go run todo.go add "Buy groceries"` |
| `done <index>`   | Toggle the completion status of a task                | `go run todo.go done 2`              |
| `remove <index>` | Remove a task from the list                           | `go run todo.go remove 3`            |


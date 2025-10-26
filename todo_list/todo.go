package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	ID    int
	Title string
	Done  bool
}

const fileName = "tasks.json"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing command")
		return
	}

	tasks, err := readOrCreateFile(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	if err := runCommand(tasks, os.Args[1], os.Args[2:]); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}

func readOrCreateFile(filePath string) ([]Task, error) {
	data, err := os.ReadFile(fileName)
	if os.IsNotExist(err) {
		if err := os.WriteFile(filePath, []byte("[]"), 0644); err != nil {
			return nil, fmt.Errorf("failed to create file: %w", err)
		}
		return []Task{}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	tasks := []Task{}

	if len(data) == 0 {
		return tasks, nil
	}

	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("invalid JSON: %w", err)
	}

	return tasks, nil
}

func runCommand(tasks []Task, cmd string, args []string) error {
	switch cmd {
	case "list":
		listTasks(tasks)
	case "add":
		if len(args) < 1 {
			return fmt.Errorf("missing argument")
		}
		name := args[0]
		tasks = addTask(tasks, name)
		listTasks(tasks)
		if err := save(tasks); err != nil {
			return fmt.Errorf("failed saving: %w", err)
		}
	case "remove":
		if len(args) < 1 {
			return fmt.Errorf("missing argument'")
		}
		delIndex, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("failed to get argument: %w", err)
		}
		tasks, err := deleteTask(tasks, delIndex)
		if err != nil {
			return fmt.Errorf("failed to remove task: %w", err)
		}
		if err := save(tasks); err != nil {
			return fmt.Errorf("failed saving: %w", err)
		}
		listTasks(tasks)
	case "done":
		if len(args) < 1 {
			return fmt.Errorf("missing argument")
		}
		doneIndex, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("failed to get argument: %w", err)
		}
		toggleDone(tasks, int(doneIndex))
		if err := save(tasks); err != nil {
			return fmt.Errorf("failed saving: %w", err)
		}
		listTasks(tasks)
	default:
		return fmt.Errorf("invalid command. expected: list, add, remove, done")
	}

	return nil
}

func listTasks(tasks []Task) {
	for k, t := range tasks {
		done := func(b bool) string {
			if b {
				return "X"
			}
			return " "
		}(t.Done)
		fmt.Printf("[%s] %d. %s\n", done, k+1, t.Title)
	}
}

func addTask(tasks []Task, name string) []Task {
	fmt.Println("Adding task...")

	if len(tasks) == 0 {
		return append(tasks, Task{
			ID:    1,
			Title: name,
		})
	}

	maxIdIndex := 0
	for k, t := range tasks {
		if t.ID > tasks[maxIdIndex].ID {
			maxIdIndex = k
		}
	}

	return append(tasks, Task{
		ID:    tasks[maxIdIndex].ID + 1,
		Title: name,
	})
}

func toggleDone(tasks []Task, index int) error {
	if index < 1 || index > len(tasks) {
		return fmt.Errorf("index %d out of range", index)
	}

	tasks[index-1].Done = !tasks[index-1].Done
	return nil
}

func deleteTask(tasks []Task, index int) ([]Task, error) {
	fmt.Println("Removing task...")

	if index > len(tasks) || index < 1 {
		return nil, fmt.Errorf("index %d out of range", index)
	}

	return append(tasks[:index-1], tasks[index-1+1:]...), nil
}

func save(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "\t")
	if err != nil {
		return err
	}
	os.WriteFile(fileName, data, 0644)
	return nil
}

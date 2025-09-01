package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

var (
	todos    []Todo
	dataFile = filepath.Join(os.Getenv("HOME"), ".todo-mcp", "todos.json")
)

func LoadTodos() error {
	data, err := os.ReadFile(dataFile)
	if err != nil {
		todos = []Todo{}
		return nil
	}
	return json.Unmarshal(data, &todos)
}

func SaveTodos() error {
	if err := os.MkdirAll(filepath.Dir(dataFile), 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(dataFile, data, 0644)
}

func CreateTodo(todo Todo) error {
	todos = append(todos, todo)
	return SaveTodos()
}

func GetAllTodos() []Todo {
	return todos
}

func GetTodoByID(id string) (*Todo, int) {
	for idx, todo := range todos {
		if todo.ID == id {
			return &todos[idx], idx
		}
	}
	return nil, -1
}

func UpdateTodo(id string, updatedTodo Todo) error {
	for idx, todo := range todos {
		if todo.ID == id {
			todos[idx] = updatedTodo
			return SaveTodos()
		}
	}
	return fmt.Errorf("todo not found: %s", id)
}

func DeleteTodo(id string) error {
	for idx, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:idx], todos[idx+1:]...)
			return SaveTodos()
		}
	}
	return fmt.Errorf("todo not found: %s", id)
}

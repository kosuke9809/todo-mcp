package todo

import (
	"os"
	"path/filepath"
)

var (
	todos    = make(map[string]*Todo)
	dataFile = filepath.Join(os.Getenv("HOME"), ".todo-mcp", "todos.json")
)

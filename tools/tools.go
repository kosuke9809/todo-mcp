package tools

import (
	"context"
	"fmt"
	"mcp-todo/todo"
	"mcp-todo/utils"
	"strings"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func CreateTodo(ctx context.Context, req *mcp.CallToolRequest, args todo.CreateTodoParams) (*mcp.CallToolResult, any, error) {
	if args.Title == "" {
		return nil, nil, fmt.Errorf("title is required")
	}

	priority := "medium"
	if args.Priority != "" {
		switch strings.ToLower(args.Priority) {
		case "high":
			priority = "high"
		case "low":
			priority = "low"
		}
	}

	todo := &todo.Todo{
		ID:          utils.GenerateID(),
		Title:       args.Title,
		Description: args.Description,
		Priority:    priority,
		DueDate:     args.DueDate,
		Tags:        args.Tags,
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: fmt.Sprintf("✅ TODO作成: %s", todo.Title)},
		},
	}, nil, nil
}

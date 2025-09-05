package tools

import (
	"context"
	"fmt"
	"mcp-todo/todo"
	"mcp-todo/utils"
	"slices"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func Create(ctx context.Context, req *mcp.CallToolRequest, args todo.CreateTodoParams) (*mcp.CallToolResult, any, error) {
	if args.Title == "" {
		return nil, nil, fmt.Errorf("title is required")
	}

	todoItem := todo.Todo{
		ID:          utils.GenerateID(),
		Title:       args.Title,
		Description: args.Description,
		Tags:        args.Tags,
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := todo.CreateTodo(todoItem); err != nil {
		return nil, nil, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: fmt.Sprintf("✅ TODO作成: %s\n📋 ID: %s", todoItem.Title, todoItem.ID)},
		},
	}, nil, nil
}

func List(ctx context.Context, req *mcp.CallToolRequest, args todo.ListTodoParams) (*mcp.CallToolResult, any, error) {
	todos := todo.GetAllTodos()

	var filtered []todo.Todo
	for _, t := range todos {
		if !args.ShowCompleted && t.Completed {
			continue
		}

		if args.Tag != "" && !slices.Contains(t.Tags, args.Tag) {
			continue
		}

		filtered = append(filtered, t)
	}

	if args.Limit > 0 && len(filtered) > args.Limit {
		filtered = filtered[:args.Limit]
	}

	if len(filtered) == 0 {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: "📝 TODOが見つかりませんでした"},
			},
		}, nil, nil
	}

	var output string
	for i, t := range filtered {
		status := "⭕"
		if t.Completed {
			status = "✅"
		}

		tagsStr := ""
		if len(t.Tags) > 0 {
			tagsStr = fmt.Sprintf(" [%s]", fmt.Sprintf("%v", t.Tags))
		}

		output += fmt.Sprintf("%d. %s %s%s\n", i+1, status, t.Title, tagsStr)
		output += fmt.Sprintf("   🆔 ID: %s\n", t.ID)
		if t.Description != "" {
			output += fmt.Sprintf("   📋 %s\n", t.Description)
		}
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: fmt.Sprintf("📝 TODO一覧 (%d件):\n%s", len(filtered), output)},
		},
	}, nil, nil
}

func Get(ctx context.Context, req *mcp.CallToolRequest, args todo.GetTodoParams) (*mcp.CallToolResult, any, error) {
	var todoItem *todo.Todo

	if args.ID != "" {
		todoItem, _ = todo.GetTodoByID(args.ID)
	} else if args.Index > 0 {
		todos := todo.GetAllTodos()
		if args.Index <= len(todos) {
			todoItem = &todos[args.Index-1]
		}
	} else {
		return nil, nil, fmt.Errorf("ID or Index is required")
	}
	if todoItem == nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("❌ TODO not found: %s", args.ID)},
			},
		}, nil, nil
	}

	status := "⭕ 未完了"
	if todoItem.Completed {
		status = "✅ 完了"
	}

	tagsStr := ""
	if len(todoItem.Tags) > 0 {
		tagsStr = fmt.Sprintf("\n🏷️ タグ: %v", todoItem.Tags)
	}

	output := fmt.Sprintf("📝 TODO詳細:\n💡 タイトル: %s\n📋 説明: %s\n%s 状態: %s\n📅 作成日: %s\n📅 更新日: %s%s",
		todoItem.Title,
		todoItem.Description,
		status[:4],
		status[4:],
		todoItem.CreatedAt.Format("2006-01-02 15:04"),
		todoItem.UpdatedAt.Format("2006-01-02 15:04"),
		tagsStr,
	)

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: output},
		},
	}, nil, nil
}

func Update(ctx context.Context, req *mcp.CallToolRequest, args todo.UpdateTodoParams) (*mcp.CallToolResult, any, error) {
	var existingTodo *todo.Todo
	var todoID string

	if args.ID != "" {
		existingTodo, _ = todo.GetTodoByID(args.ID)
		todoID = args.ID
	} else if args.Index > 0 {
		todos := todo.GetAllTodos()
		if args.Index <= len(todos) {
			existingTodo = &todos[args.Index-1]
			todoID = existingTodo.ID
		}
	} else {
		return nil, nil, fmt.Errorf("ID or Index is required")
	}
	if existingTodo == nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("❌ TODO not found: %s", args.ID)},
			},
		}, nil, nil
	}

	updatedTodo := *existingTodo
	updatedTodo.UpdatedAt = time.Now()

	if args.Title != "" {
		updatedTodo.Title = args.Title
	}
	if args.Description != "" {
		updatedTodo.Description = args.Description
	}
	if len(args.Tags) > 0 {
		updatedTodo.Tags = args.Tags
	}

	if err := todo.UpdateTodo(todoID, updatedTodo); err != nil {
		return nil, nil, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: fmt.Sprintf("✅ TODO更新: %s", updatedTodo.Title)},
		},
	}, nil, nil
}

func Delete(ctx context.Context, req *mcp.CallToolRequest, args todo.DeleteTodoParams) (*mcp.CallToolResult, any, error) {
	var existingTodo *todo.Todo
	var todoID string

	if args.ID != "" {
		existingTodo, _ = todo.GetTodoByID(args.ID)
		todoID = args.ID
	} else if args.Index > 0 {
		todos := todo.GetAllTodos()
		if args.Index <= len(todos) {
			existingTodo = &todos[args.Index-1]
			todoID = existingTodo.ID
		}
	} else {
		return nil, nil, fmt.Errorf("ID or Index is required")
	}
	if existingTodo == nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("❌ TODO not found: %s", args.ID)},
			},
		}, nil, nil
	}

	if err := todo.DeleteTodo(todoID); err != nil {
		return nil, nil, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: fmt.Sprintf("🗑️ TODO削除: %s", existingTodo.Title)},
		},
	}, nil, nil
}

func Complete(ctx context.Context, req *mcp.CallToolRequest, args todo.CompleteTodoParams) (*mcp.CallToolResult, any, error) {
	var existingTodo *todo.Todo
	var todoID string

	if args.ID != "" {
		existingTodo, _ = todo.GetTodoByID(args.ID)
		todoID = args.ID
	} else if args.Index > 0 {
		todos := todo.GetAllTodos()
		if args.Index <= len(todos) {
			existingTodo = &todos[args.Index-1]
			todoID = existingTodo.ID
		}
	} else {
		return nil, nil, fmt.Errorf("ID or Index is required")
	}
	if existingTodo == nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("❌ TODO not found: %s", args.ID)},
			},
		}, nil, nil
	}

	updatedTodo := *existingTodo
	updatedTodo.Completed = args.Completed
	updatedTodo.UpdatedAt = time.Now()

	if err := todo.UpdateTodo(todoID, updatedTodo); err != nil {
		return nil, nil, err
	}

	status := "⭕ 未完了"
	if args.Completed {
		status = "✅ 完了"
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: fmt.Sprintf("%s TODO: %s", status, updatedTodo.Title)},
		},
	}, nil, nil
}

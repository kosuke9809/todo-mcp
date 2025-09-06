package server

import (
	"github.com/kosuke9809/todo-mcp/tools"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func NewServer() *mcp.Server {
	server := mcp.NewServer(
		&mcp.Implementation{
			Name:    "todo",
			Version: "v1.0.0",
		}, nil)

	registerTools(server)

	return server
}

func registerTools(server *mcp.Server) {

	mcp.AddTool(server, &mcp.Tool{
		Name:        "todo__create",
		Description: "TODOを作成します。タイトルは必須です。例: '買い物に行く'",
	}, tools.Create)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "todo__list",
		Description: "TODOを一覧表示します。",
	}, tools.List)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "todo__delete",
		Description: "TODOを削除します。例: '1'",
	}, tools.Delete)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "todo__update",
		Description: "TODOを更新します。例: '1' '買い物に行く'",
	}, tools.Update)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "todo__complete",
		Description: "TODOを完了します。例: '1'",
	}, tools.Complete)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "todo__get",
		Description: "TODOを取得します。例: '1'",
	}, tools.Get)
}

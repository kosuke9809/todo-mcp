package main

import (
	"context"
	"log"

	"github.com/kosuke9809/todo-mcp/server"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	mcpServer := server.NewServer()

	if err := mcpServer.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"context"
	"log"
	"mcp-todo/server"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	mcpServer := server.NewServer()

	if err := mcpServer.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}

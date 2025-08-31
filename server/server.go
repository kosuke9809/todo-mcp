package server

import "github.com/modelcontextprotocol/go-sdk/mcp"

func NewServer() *mcp.Server {
	server := mcp.NewServer(
		&mcp.Implementation{
			Name:    "todo",
			Version: "v1.0.0",
		}, nil)

	return server

}

func registerTools() error {
	return nil
}

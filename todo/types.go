package todo

import "time"

type Todo struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"descriptioni"`
	Tags        []string  `json:"tags"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateTodoParams struct {
	Title       string
	Description string
	Tags        []string
}

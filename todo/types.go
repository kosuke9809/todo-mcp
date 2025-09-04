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

type ListParams struct {
	Tag           string `json:"tag,omitempty" jsonschema:"特定のタグで絞り込み"`
	ShowCompleted bool   `json:"showCompleted,omitempty" jsonschema:"完了済みTODOも表示するか"`
	Limit         int    `json:"limit,omitempty" jsonschema:"取得件数の上限"`
}

type GetTodoParams struct {
	ID string `json:"id" jsonschema:"取得するTODOのID"`
}

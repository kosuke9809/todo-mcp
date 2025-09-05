package todo

import "time"

type Todo struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Tags        []string  `json:"tags"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateTodoParams struct {
	Title       string   `json:"title" jsonschema:"TODOのタイトル（必須）"`
	Description string   `json:"description,omitempty" jsonschema:"TODOの詳細説明"`
	Tags        []string `json:"tags,omitempty" jsonschema:"タグのリスト"`
}

type ListTodoParams struct {
	Tag           string `json:"tag,omitempty" jsonschema:"特定のタグで絞り込み"`
	ShowCompleted bool   `json:"showCompleted,omitempty" jsonschema:"完了済みTODOも表示するか"`
	Limit         int    `json:"limit,omitempty" jsonschema:"取得件数の上限"`
}

type GetTodoParams struct {
	ID    string `json:"id,omitempty" jsonschema:"取得するTODOのID"`
	Index int    `json:"index,omitempty" jsonschema:"取得するTODOのインデックス番号（1から開始）"`
}

type UpdateTodoParams struct {
	ID          string   `json:"id,omitempty" jsonschema:"更新するTODOのID"`
	Index       int      `json:"index,omitempty" jsonschema:"更新するTODOのインデックス番号（1から開始）"`
	Title       string   `json:"title,omitempty" jsonschema:"新しいタイトル"`
	Description string   `json:"description,omitempty" jsonschema:"新しい詳細説明"`
	Tags        []string `json:"tags,omitempty" jsonschema:"新しいタグのリスト"`
}

type DeleteTodoParams struct {
	ID    string `json:"id,omitempty" jsonschema:"削除するTODOのID"`
	Index int    `json:"index,omitempty" jsonschema:"削除するTODOのインデックス番号（1から開始）"`
}

type CompleteTodoParams struct {
	ID        string `json:"id,omitempty" jsonschema:"完了状態を変更するTODOのID"`
	Index     int    `json:"index,omitempty" jsonschema:"完了状態を変更するTODOのインデックス番号（1から開始）"`
	Completed bool   `json:"completed" jsonschema:"完了状態（true: 完了、false: 未完了）"`
}

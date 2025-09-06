# TODO MCP Server

TODOを管理するためのMCPサーバーです

## 機能

- TODOの作成・更新・削除
- TODOの一覧表示・個別取得
- タグによる分類・絞り込み
- 完了状態の管理


## インストール

```bash
go install github.com/kosuke9809/todo-mcp@latest
```

## 使用方法

### Claude Desktop/Codeでの設定

`claude_desktop_config.json` に以下を追加：

```json
{
  "mcpServers": {
    "todo-mcp": {
      "command": "todo-mcp"
    }
  }
}
```

## Tools

### 1. create-todo
新しいTODOを作成します。

**パラメータ:**
- `title` (必須): TODOのタイトル
- `description` (オプション): TODOの詳細説明
- `tags` (オプション): タグのリスト

**例:**
```
create-todo title:"買い物に行く" description:"牛乳とパンを買う" tags:["日常", "買い物"]
```

### 2. list-todos
TODOの一覧を表示します。

**パラメータ:**
- `tag` (オプション): 特定のタグで絞り込み
- `showCompleted` (オプション): 完了済みTODOも表示するか (default: false)
- `limit` (オプション): 取得件数の上限

**例:**
```
list-todos
list-todos tag:"買い物" showCompleted:true limit:5
```

### 3. get-todo
特定のTODOの詳細を取得します。

**パラメータ (どちらか一つが必須):**
- `id`: TODOのID
- `index`: TODOのインデックス番号（1から開始）

**例:**
```
get-todo index:1
get-todo id:"abc123..."
```

### 4. update-todo
既存のTODOを更新します。

**パラメータ:**
- `id` または `index` (どちらか必須): 更新するTODOの指定
- `title` (オプション): 新しいタイトル
- `description` (オプション): 新しい詳細説明
- `tags` (オプション): 新しいタグのリスト

**例:**
```
update-todo index:1 title:"新しいタイトル"
update-todo id:"abc123..." description:"更新された説明"
```

### 5. delete-todo
TODOを削除します。

**パラメータ (どちらか一つが必須):**
- `id`: TODOのID
- `index`: TODOのインデックス番号（1から開始）

**例:**
```
delete-todo index:3
delete-todo id:"abc123..."
```

### 6. complete-todo
TODOの完了状態を変更します。

**パラメータ:**
- `id` または `index` (どちらか必須): 対象TODOの指定
- `completed` (必須): 完了状態 (true: 完了、false: 未完了)

**例:**
```
complete-todo index:1 completed:true
complete-todo index:2 completed:false
```

## データストレージ

TODOデータは `~/.todo-mcp/todos.json` に保存されます。

## 使用例

### 基本的なワークフロー

1. **TODOを作成**
   ```
   create-todo title:"レポート作成" tags:["仕事", "重要"]
   ```

2. **一覧を確認**
   ```
   list-todos
   ```
   出力例:
   ```
   📝 TODO一覧 (1件):
   1. ⭕ レポート作成 [仕事 重要]
      🆔 ID: abc12345-...
   ```

3. **インデックス番号で操作**
   ```
   update-todo index:1 description:"明日までに完成させる"
   complete-todo index:1 completed:true
   ```

4. **完了済みも含めて確認**
   ```
   list-todos showCompleted:true
   ```

### タグを使った管理

```
create-todo title:"食材購入" tags:["買い物", "日常"]
create-todo title:"プレゼント選び" tags:["買い物", "特別"]
list-todos tag:"買い物"
```

## ライセンス

MIT License

# 開発メモ
- sqlboilerのモデル自動生成
`touch sqlboiler.toml`
```
pkname="models"
output="db/models"
[psql]
  dbname = "dynamic_novel"
  host   = "db"
  port   = 5432
  user   = "postgres"
  pass   = "pass"
  sslmode = "disable"
  schema = "public"
  blacklist = ["migrations", "other"]
```
`sqlboiler psql`

# connect-goのコード生成
1. protoファイル準備
`touch proto/todo/v1/todo.proto`
```
syntax = "proto3";

package proto.todo.v1;

option go_package = "github.com/mamoruuji/dynamic_novel_api/gen/proto/todo/v1;todov1";

message TodoData {
  int32 id = 1;
  string title = 2;
  string context = 3;
}

message ListTodosRequest {}

message ListTodosResponse {
  repeated TodoData todos = 1;
}

message AddTodoRequest {
  string title = 1;
  string context = 2;
}

message AddTodoResponse {
  int32 id = 1;
}

message DeleteTodoRequest {
  int32 id = 1;
}

message DeleteTodoResponse {}

message UpdateTodoStatusRequest {
  int32 id = 1;
  string title = 2;
  string context = 3;
}

message UpdateTodoStatusResponse {}

service TodoService {
  rpc ListTodos(ListTodosRequest) returns (ListTodosResponse) {}
  rpc AddTodo(AddTodoRequest) returns (AddTodoResponse) {}
  rpc DeleteTodo(DeleteTodoRequest) returns (DeleteTodoResponse) {}
  rpc UpdateTodoStatus(UpdateTodoStatusRequest) returns (UpdateTodoStatusResponse) {}
}
```

2. bufファイル準備
`buf mod init`
buf.yaml生成

`touch buf.gen.yaml`
```
version: v1
plugins:
  - plugin: buf.build/protocolbuffers/go
    out: gen
    opt: paths=source_relative
  - plugin: buf.build/bufbuild/connect-go
    out: gen
    opt: paths=source_relative
```
3. protoファイルからコード生成
`buf lint`
`buf mod update`
`buf generate`
## goライブラリ管理
`touch go.mod`
```
module github.com/mamoruuji/dynamic_novel_api

go 1.21
```
`touch go.sum`

`go mod tidy`
## 動作確認
`go run ./cmd/server/main.go`

```
curl \
--header "Content-Type: application/json" \
--data '{}' \
http://localhost:8080/proto.todo.v1.TodoService/ListTodos
```

- 別ターミナルから実行
```
curl \
--header "Content-Type: application/json" \
--data '{}' \
http://dynamic_novel_server:8080/proto.todo.v1.TodoService/ListTodos
```

## パスの重複削除
```
export PATH="$PATH:$(go env GOPATH)/bin"
echo $PATH | tr : \\n | sort
export PATH="$(printf %s "$PATH" | awk -v RS=: -v ORS=: '!arr[$0]++')"
```

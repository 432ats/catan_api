# catan_api

## memo
- `gqlgen.yml` : gqlgen の設定ファイル
- `graph/generated/generated.go` : 神様が自動生成（触っちゃダメ❤️）
- `graph/model/models_gen.go` : `schema.graphqls` に基づいて神様が自動生成（触っちゃダメ❤️）
- `graph/resolver.go` : GraphQLサーバーが読み込むリゾルバー（ゴリゴリ書くぞ🐯）
- `graph/schema.graphqls` : GraphQLのスキーマファイル（ゴリゴリ書くぞ🐯）
- `graph/schema.resolvers.go` : まだよう分からん🤔

## 実装フロー
スキーマファイル(`schema.graphqls`)の更新を行ったら、generated.goに反映するために以下を実行
```
go run github.com/99designs/gqlgen generate
```
もしくは
```
go generate ./...
```
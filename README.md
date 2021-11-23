# go-echo-bbs

## セットアップ

postgresを起動しスキーマを適用します。

```shell
$docker-compose up -d
```

初回の場合はスキーマを適用します。データベースの初期パスワードは `go_echo_bbs` です。

```shell
$psql --host=127.0.0.1 --port=5432 --username=go_echo_bbs --password go_echo_bbs_development < schema.sql
```

server.goを実行します。

```shell
$go run server.go
```

## .envの読み込み

`.env` は `GO_ENV` が `development` の時のみ読み込むようにしています。

## sqlc の生成

```shell
$docker run --rm -v $(pwd):/src -w /src kjconroy/sqlc generate
```
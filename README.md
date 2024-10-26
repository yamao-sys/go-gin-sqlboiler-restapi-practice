# go-gin-sqlboiler-restapi-practice
Go Gin × SQLBoilerによるRESTAPIの練習

- sqlboilerとgo-txdbの依存ライブラリの競合が...

## コマンド類
- Webサーバ起動
```
godotenv -f /app/.env go run main.go
```

- テスト用DBのマイグレーション
```
godotenv -f /app/.env.test.local sql-migrate up -env="mysql"
```

- テスト実行
```
godotenv -f /app/.env.test.local test -v ./...
```

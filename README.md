# Lesson Gin by Golang

```
mkdir sand-gin && cd $_
go mod init sand-gin
```

依存関係のインストール

```
go get github.com/gin-gonic/gin
go get gorm.io/gorm
go get gorm.io/driver/mysql
```

上から順に

- Gin
- Gorm（ORM）
- MySQL と接続するためのドライバー

実行

```
go run main.go
```

.env を読み込んでくれるやつ

```
go get github.com/joho/godotenv
```

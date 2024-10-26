package main

import (
	models "app/models/generated"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	dsn := os.Getenv("MYSQL_USER") +
		":" + os.Getenv("MYSQL_PASS") +
		"@tcp(" + os.Getenv("MYSQL_HOST") + ":" + os.Getenv("MYSQL_PORT") + ")/" +
		os.Getenv("MYSQL_DBNAME") +
		"?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true&loc=Local"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	ctx := context.Background()

	// NOTE: SQL文を標準出力に出す
	boil.DebugMode = true

	hash, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln(err)
	}
	user := &models.User{
		Name:     "test_name_2",
		Email:    "test_2@example.com",
		Password: string(hash),
	}
	createErr := user.Insert(ctx, db, boil.Infer())
	if createErr != nil {
		log.Fatalln(createErr)
	}
	fmt.Println(user.Reload(ctx, db))
}

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

// добавить в базу:  & "C:\Program Files\PostgreSQL\16\bin\psql.exe" -U postgres -h localhost -d go_dev -f db.sql --password
// посмотреть таблицу users: & "C:\Program Files\PostgreSQL\16\bin\psql.exe" -U postgres -h localhost -d go_dev -c "\d users"
// узнать текущего пользователя: & "C:\Program Files\PostgreSQL\16\bin\psql.exe" -U postgres -h localhost -c "SELECT current_user;"
// !!! предварительно необходимо добавить свой пароль к базе данных командой: $env:DB_PASSWORD="свой пароль"
// проверить пароль: echo $env:DB_PASSWORD

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	databaseUrlFormat := os.Getenv("DATABASE_URL_FORMAT")
	if databaseUrlFormat == "" {
		panic("DATABASE_URL_FORMAT is empty")
	}

	databaseUrl := fmt.Sprintf(databaseUrlFormat, os.Getenv("DB_PASSWORD"))

	fmt.Println("Database URL:", databaseUrl)

	conn, err := pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
}

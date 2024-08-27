package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
)

// & "C:\Program Files\PostgreSQL\16\bin\psql.exe" -U postgres -h localhost -d go_dev -c "DROP TABLE IF EXISTS users CASCADE;"

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	databaseUrlFormat := os.Getenv("DATABASE_URL_FORMAT")
	if databaseUrlFormat == "" {
		panic("DATABASE_URL_FORMAT is empty")
	}

	databaseUrl := fmt.Sprintf(databaseUrlFormat, os.Getenv("DB_PASSWORD"))

	fmt.Println("Database URL:", databaseUrl)

	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	// Вывод текущей рабочей директории
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	migrationsPath := filepath.Join(cwd, "../../migrations")

	if err := goose.Up(db, migrationsPath); err != nil {
		panic(err)
	}
}
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	databaseUrlFormat := os.Getenv("DATABASE_URL_FORMAT")
	if databaseUrlFormat == "" {
		panic("DATABASE_URL_FORMAT is empty")
	}

	databaseUrl := fmt.Sprintf(databaseUrlFormat+"?sslmode=disable", os.Getenv("DB_PASSWORD"))

	fmt.Println("Database URL:", databaseUrl)

	db, err := sql.Open("postgres", databaseUrl)

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"postgres", driver)
	if err != nil {
		panic(err)
	}

	err = m.Up()
	if err != nil {
		if err != migrate.ErrNoChange {
			panic(err)
		}
	}
	fmt.Println("ok")
}

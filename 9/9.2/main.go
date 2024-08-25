package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

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

	user, err := GetUser(conn, 2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("get user: %+v\n", user)

	users, err := GetUsers(conn)
	if err != nil {
		panic(err)
	}
	fmt.Println("get users: ")
	for _, user := range users {
		fmt.Printf("%+v\n", user)
	}

	userId, err := InsertUser(conn, User{Name: "Test", Email: "test@test"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("new user id: %d\n", userId)

	rowsAffected, err := DeleteUser(conn, 1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("delete users: %d\n", rowsAffected)
}

type User struct {
	Id     int
	Name   string
	Email  string
	Photos []Photo
}

type Photo struct {
	UserId    int
	Filename  string
	Width     int
	Height    int
	CreatedAt time.Time
}

func GetUser(conn *pgx.Conn, userId int) (User, error) {
	var user User
	row := conn.QueryRow(
		context.Background(),
		"select id, name, email from users where id=$1",
		userId,
	)
	err := row.Scan(
		&user.Id,
		&user.Name,
		&user.Email,
	)
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetUsers(conn *pgx.Conn) ([]User, error) {
	users := make([]User, 0)

	rows, err := conn.Query(
		context.Background(),
		"select id, name, email from users",
	)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Email,
		)
		if err != nil {
			return []User{}, err
		}
		users = append(users, user)
	}

	return users, nil
}

func InsertUser(conn *pgx.Conn, user User) (int, error) {
	var id int
	err := conn.QueryRow(
		context.Background(),
		"insert into users(name, email) values($1, $2) returning id",
		user.Name, user.Email,
	).Scan(&id)

	if err != nil {
		return 0, err
	}
	return id, nil
}

func DeleteUser(conn *pgx.Conn, userId int) (int, error) {
	tag, err := conn.Exec(
		context.Background(),
		"delete from users where id=$1",
		userId,
	)
	if err != nil {
		return 0, err
	}
	return int(tag.RowsAffected()), nil
}
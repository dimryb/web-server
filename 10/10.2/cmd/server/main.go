package main

import (
	"example/internal/app"
)

func main() {
	app := app.NewApp()
	err := app.Config.Load(".env")
	if err != nil {
		panic(err)
	}
}

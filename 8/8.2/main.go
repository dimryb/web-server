package main

import (
	"os"
	"text/template"
)

func main() {
	tmpl, err := template.ParseFiles(
		"templates/layout2.txt",
		"templates/page1.txt",
	)
	if err != nil {
		panic(err)
	}

	data := map[string]any{
		"text": "some text",
		"number": 42,
	}

	tmpl.ExecuteTemplate(os.Stdout, "layout2.txt", data)
}
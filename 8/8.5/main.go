package main

import (
	"os"
	"text/template"
)

func main() {
	tmpl, err := template.New("some.txt").Parse(`
		header
		{{.text}}
		footer
	`)
	if err != nil {
		panic(err)
	}

	tmpl, err = template.New("bbb.txt").Parse(`
		bbb content
	`)
	if err != nil {
		panic(err)
	}


	data := map[string]any{
		"text": "some text",
		"number": 42,
	}

	tmpl.ExecuteTemplate(os.Stdout, "bbb.txt", data)
}
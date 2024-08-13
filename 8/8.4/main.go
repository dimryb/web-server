package main

import (
	"io"
	"net/http"
	"html/template"
)

func main() {
	http.HandleFunc("/", HomeHandler)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}

	type User struct {
		Name string
	}
	user := User{Name: "<script>alert('hacker')</script>"}

	tmpl.ExecuteTemplate(w, "index.html", user)
}
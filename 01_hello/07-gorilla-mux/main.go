package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/hello/{username}", hello)
	r.HandleFunc(`/product/{id:\d+}`, product)
	r.HandleFunc(`/form`, form).Methods("POST", "PUT")
	r.NotFoundHandler = http.HandlerFunc(handler404)

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}


func home(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Home")
}

func hello(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userName := vars["username"]
	fmt.Fprintf(w, "Hello %s!", userName)
}

func product(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "Product ID %s", id)
}

func form(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "From")
}

func handler404(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 Page not Found"))
}

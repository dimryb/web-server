package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/hello/", hello)

	http.Handle("/foo", foo())

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}

func foo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Foo"))
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		handler404(w, r)
		return
	}
	w.Write([]byte("Home"))
}

func hello(w http.ResponseWriter, r *http.Request) {
	pathRegexp := regexp.MustCompile(`^/hello/\w+$`)
	if !pathRegexp.Match([]byte(r.URL.Path)) {
		handler404(w, r)
		return
	}
	name := strings.Split(r.URL.Path, "/")[2]
	w.Write([]byte(fmt.Sprintf("Hello, %s", name)))
}

func handler404(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 Page not Found"))
}

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/foo", FooHandler)

	middlewares := []func(http.Handler) http.Handler{
		LoggingMiddleware,
		SecondMiddleware,
	}

	handler := http.Handler(mux)
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}

	err := http.ListenAndServe(":3000", handler)
	if err != nil {
		panic(err)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("home")
	w.Write([]byte("Home"))
}

func FooHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("foo")	
	w.Write([]byte("Foo"))
}

type MyResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *MyResponseWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func LoggingMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w2 := &MyResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		handler.ServeHTTP(w2, r)
		log.Printf("%s : [%d]\n", r.RequestURI, w2.statusCode)
	})
}

func SecondMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("before 2")
		handler.ServeHTTP(w, r)
		fmt.Println("after 2")
	})
}

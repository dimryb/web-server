package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

func main() {
	http.HandleFunc("/form", FormHandler)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 * 1024 * 1024)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
		return
	}

	file, header, err := r.FormFile("myfile")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}
	defer file.Close()
	fmt.Println("filename", header.Filename)
	fmt.Println("MIME Type", header.Header)
	fmt.Println("Size", header.Size)

	ext := path.Ext(header.Filename)
	tempDir := os.TempDir()
	tempFile, err := ioutil.TempFile(tempDir, "*"+ext)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}
	defer tempFile.Close()

	fmt.Println("Save to: ", tempFile.Name())

	bytes, err := io.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}

	_, err = tempFile.Write(bytes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}

	io.WriteString(w, "File uploaded!")
}

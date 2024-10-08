package rest

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Ok     bool   `json:"ok"`
	Result any    `json:"result,omitempty"`
	Error  string `json:"json:error,omitempty"`
}

func WriteJSON(w http.ResponseWriter, status int, v any) {
	bytes, _ := json.Marshal(v)
	w.WriteHeader(status)
	w.Write(bytes)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, Response{
		Ok:    false,
		Error: err.Error(),
	})
}

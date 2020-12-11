package main

import (
	"net/http"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello from Golang Server"))
	}
}

func main() {
	http.HandleFunc("/", mainHandler)
	http.ListenAndServe(":4000", nil)
}

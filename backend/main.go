package main

import (
	"backend/apis"
	"net/http"
)

func main() {
	http.HandleFunc("/generate/password", apis.GeneratePassword)

	_ = http.ListenAndServe(":8080", nil)
}

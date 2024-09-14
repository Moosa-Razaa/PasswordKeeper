package main

import (
	"backend/apis"
	"backend/middlewares"
	"net/http"
)

func main() {
	http.Handle("/generate/password", middlewares.VerifyPost(http.HandlerFunc(apis.GeneratePassword)))
	_ = http.ListenAndServe(":8080", nil)
}

package main

import (
	"backend/middlewares"
	"backend/services"
	"net/http"
)

func main() {
	http.Handle("/generate/password", middlewares.VerifyPost(http.HandlerFunc(services.PasswordGeneratorService)))
	_ = http.ListenAndServe(":8080", nil)
}

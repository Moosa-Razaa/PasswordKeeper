package main

import (
	"backend/controllers"
	"backend/middlewares"
	"net/http"
)

func main() {
	http.Handle("/generate/password", middlewares.VerifyPost(http.HandlerFunc(controllers.PasswordGeneratorController)))
	_ = http.ListenAndServe(":8080", nil)
}

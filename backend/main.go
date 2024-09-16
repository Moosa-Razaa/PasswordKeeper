package main

import (
	"backend/controllers"
	"backend/middlewares"
	"net/http"
)

func main() {
	http.Handle("/generate/password", middlewares.VerifyPost(http.HandlerFunc(controllers.PasswordGeneratorService)))
	_ = http.ListenAndServe(":8080", nil)
}

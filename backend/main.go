package main

import (
	"backend/controllers"
	"backend/middlewares"
	"backend/repositories"
	"log"
	"net/http"
)

var fileHandler *repositories.FileHandler

func InitializeFileHandler() {
	var fileHandlerError error

	fileHandler, fileHandlerError = repositories.GetFileHandlerInstance()
	if fileHandlerError != nil {
		log.Fatalf("Error while initializing file handler: %v", fileHandlerError)
	}
}

func main() {
	InitializeFileHandler()

	http.Handle("/generate/password", middlewares.VerifyPost(http.HandlerFunc(controllers.PasswordGeneratorController)))
	_ = http.ListenAndServe(":8080", nil)
}

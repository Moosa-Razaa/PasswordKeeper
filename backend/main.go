package main

import (
	"backend/controllers"
	"backend/middlewares"
	"backend/repositories"
	"net/http"
)

func main() {
	initializationError := repositories.InitializeFileIfNotExist()

	if initializationError != nil {
		panic(initializationError)
	}

	http.Handle("/generate/password", middlewares.VerifyPost(http.HandlerFunc(controllers.PasswordGeneratorController)))
	http.Handle("/create", http.HandlerFunc(controllers.AddNewPassword))

	serverListeningError := http.ListenAndServe(":8080", nil)

	if serverListeningError != nil {
		panic(serverListeningError)
	}
}

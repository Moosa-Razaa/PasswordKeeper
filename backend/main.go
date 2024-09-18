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
	http.Handle("/create", middlewares.VerifyPost(http.HandlerFunc(controllers.AddNewPassword)))
	http.Handle("/delete", middlewares.VerifyDelete(http.HandlerFunc(controllers.DeletePassword)))

	serverListeningError := http.ListenAndServe(":8080", nil)

	if serverListeningError != nil {
		panic(serverListeningError)
	}
}

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
	http.Handle("/update", middlewares.VerifyPatch(http.HandlerFunc(controllers.UpdatePasswordController)))
	http.Handle("/get", middlewares.VerifyGet(http.HandlerFunc(controllers.GetAllPasswordsController)))
	http.Handle("/check/repository", middlewares.VerifyGet(http.HandlerFunc(controllers.CheckFile)))

	serverListeningError := http.ListenAndServe(":8080", nil)

	if serverListeningError != nil {
		panic(serverListeningError)
	}
}

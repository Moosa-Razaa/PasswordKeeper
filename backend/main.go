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

	mux := http.NewServeMux()

	mux.Handle("/generate/password", middlewares.VerifyPost(http.HandlerFunc(controllers.PasswordGeneratorController)))
	mux.Handle("/create", middlewares.VerifyPost(http.HandlerFunc(controllers.AddNewPassword)))
	mux.Handle("/delete", middlewares.VerifyDelete(http.HandlerFunc(controllers.DeletePassword)))
	mux.Handle("/update", middlewares.VerifyPatch(http.HandlerFunc(controllers.UpdatePasswordController)))
	mux.Handle("/get", middlewares.VerifyGet(http.HandlerFunc(controllers.GetAllPasswordsController)))
	mux.Handle("/check/repository", middlewares.VerifyGet(http.HandlerFunc(controllers.CheckFile)))

	corsMux := middlewares.CORS(mux)

	serverListeningError := http.ListenAndServe(":8080", corsMux)

	if serverListeningError != nil {
		panic(serverListeningError)
	}
}
